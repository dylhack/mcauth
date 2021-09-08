package bot

import (
	c "github.com/dylhack/mcauth/internal/common"
)

// PlayerDetails will resolve the Minecraft
// player's MCAuth state and Discord roles.
func (bot *Bot) PlayerDetails(playerID string) c.PlayerDetailsRes {
	userID, _ := bot.store.Links.GetDiscordID(playerID)
	res := c.PlayerDetailsRes{
		Id:       "",
		State:    "",
		Roles:    make([]string, 0),
		AuthCode: "",
	}

	// if userID doesn't resolve then
	// they're not linked to MCAuth
	if len(userID) == 0 {
		res.State = c.NoLink
		return res
	} else {
		res.Id = userID
	}

	roles, state := bot.getUserState(playerID, userID)

	res.State = state
	res.Roles = *roles

	return res
}

func (bot *Bot) getUserState(
	playerID string,
	userID string,
) (*[]string, string) {
	roles, isOK := bot.sync.GetRoles(userID)

	if !isOK {
		bot.syncMember(userID)
		roles, _ = bot.sync.GetRoles(userID)
	}

	// check if they're an alt
	alt, _ := bot.store.Alts.GetAlt(playerID)

	if len(alt.Owner) > 0 {
		return roles, c.IsAlt
	}

	// check whether they have an administrator roles. they pass any exceptions.
	isWhitelisted, hasAdmin := bot.CheckRoles(*roles)

	// if they're an admin then they pass all exceptions
	if hasAdmin {
		return roles, c.IsAdmin
	}

	// if they're a regular user then check if they have the right roles
	if !isWhitelisted {
		return roles, c.NotWhitelisted
	}

	// finally return true
	return roles, c.Whitelisted
}
