package bot

import "fmt"

func (bot *Bot) getPlayersOnline() int {
	return 0
}

func (bot *Bot) getLinkedAccounts() int {
	linked := bot.store.Links.GetAllLinks()

	return len(linked)
}

// output:
// - role name 1
// - role name 2
// - role name 3
func (bot *Bot) getWhitelistedRoles() string {
	var list string
	for _, roleID := range bot.config.Whitelist {
		role, err := bot.client.State.Role(bot.config.Guild, roleID)

		if err == nil {
			list += fmt.Sprintf(" * %s\n", role.Name)
		}
	}
	return list
}

// output:
// - role name 1
// - role name 2
// - role name 3
func (bot *Bot) getAdminRoles() string {
	var list string
	for _, roleID := range bot.config.AdminRoles {
		role, err := bot.client.State.Role(bot.config.Guild, roleID)

		if err == nil {
			list += fmt.Sprintf(" * %s\n", role.Name)
		}
	}
	return list
}
