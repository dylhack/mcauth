package bot

import "fmt"

func (bot *Bot) countAltAccounts() int {
	alts := bot.store.Alts.GetAllAlts()

	return len(alts)
}

func (bot *Bot) countPlayersOnline() int {
	return 0
}

func (bot *Bot) countLinkedAccounts() int {
	linked := bot.store.Links.GetAllLinks()

	return len(linked)
}

func (bot *Bot) countPendingAuthCodes() int {
	authCodes := bot.store.Auth.GetAllAuthCodes()

	return len(authCodes)
}

// output:
// - role name 1
// - role name 2
// - role name 3
func (bot *Bot) getWhitelistedRoles() string {
	var list string
	roles, err := bot.client.GuildRoles(bot.config.Guild)

	if err != nil {
		return "Failed to get roles"
	}

	for _, role := range roles {
		for _, roleID := range bot.config.Whitelist {
			if roleID == role.ID {
				list += fmt.Sprintf("%s\n", role.Name)
				break
			}
		}
	}

	if len(list) == 0 {
		list = "No roles"
	}

	return list
}

// output:
// - role name 1
// - role name 2
// - role name 3
func (bot *Bot) getAdminRoles() string {
	var list string
	roles, err := bot.client.GuildRoles(bot.config.Guild)

	if err != nil {
		return "Failed to get roles"
	}

	for _, role := range roles {
		for _, roleID := range bot.config.AdminRoles {
			if roleID == role.ID {
				list += fmt.Sprintf("%s\n", role.Name)
				break
			}
		}
	}

	if len(list) == 0 {
		list = "No roles"
	}

	return list
}
