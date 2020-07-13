package bot

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	c "github.com/dhghf/mcauth/internal/common"
	"log"
)

// if reason is empty and isValid is true then ignore reason
func (bot *Bot) ValidatePlayer(playerID string) (bool, string) {
	userID := bot.store.Links.GetDiscordID(playerID)

	if len(userID) == 0 {
		return false, c.NoLink
	}

	member, err := bot.client.GuildMember(
		bot.config.Guild,
		userID,
	)

	if err != nil {
		log.Printf("Failed to fetch \"%s\", because\n%s\n", userID, err.Error())
		return false, c.NoLink
	}

	// check whether they have an administrator roles. they pass any exceptions.
	hasAdmin, _ := util.HasRole(member.Roles, bot.config.AdminRoles)

	if hasAdmin {
		return true, ""
	}

	// if they're a regular user ...
	isWhitelisted, _ := util.HasRole(member.Roles, bot.config.Whitelist)

	if !isWhitelisted {
		return false, c.NotWhitelisted
	}

	// check if maintenance mode is on
	if bot.locked {
		return false, c.Maintenance
	}

	// finally return true
	return true, ""
}
