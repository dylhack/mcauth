package bot

import (
	c "github.com/dhghf/mcauth/internal/common"
	"log"
)

// if reason is empty and isValid is true then ignore reason
func (bot *Bot) VerifyPlayer(playerID string) (bool, string) {
	alt, _ := bot.store.Alts.GetAlt(playerID)

	// check if they're an alt account
	if len(alt.Owner) > 0 {
		log.Printf(
			"%s (%s) is an alt account of %s\n",
			alt.PlayerName, alt.PlayerID,
			alt.Owner,
		)
		userID, _ := bot.store.Links.GetDiscordID(alt.Owner)

		if len(userID) == 0 {
			return false, c.NoLink
		}

		return bot.VerifyDiscordUser(userID)
	}

	userID, _ := bot.store.Links.GetDiscordID(playerID)

	if len(userID) == 0 {
		return false, c.NoLink
	}

	return bot.VerifyDiscordUser(userID)
}

func (bot *Bot) VerifyDiscordUser(userID string) (bool, string) {
	roles, isOK := bot.sync.GetRoles(userID)

	if !isOK {
		bot.Sync(userID)
		roles, _ = bot.sync.GetRoles(userID)
	}

	// check whether they have an administrator roles. they pass any exceptions.
	isWhitelisted, hasAdmin := bot.CheckRoles(*roles)

	// if they're an admin then they pass all exceptions
	if hasAdmin {
		return true, "Administrator"
	}

	// if they're a regular user then check if they have the right roles
	if !isWhitelisted {
		return false, c.NotWhitelisted
	}

	// check if maintenance mode is on, regular users aren't allowed to join during maintenance
	if bot.locked {
		return false, c.Maintenance
	}

	// finally return true
	return true, "Whitelisted"
}
