package bot

import (
	c "github.com/dhghf/mcauth/internal/common"
	"log"
)

// if reason is empty and isValid is true then ignore reason
func (bot *Bot) VerifyPlayer(playerID string) (bool, string) {
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
	isWhitelisted, hasAdmin := bot.CheckRoles(member.Roles)

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
