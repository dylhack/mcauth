package bot

import (
	dg "github.com/bwmarrin/discordgo"
	c "github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"log"
)

type Bot struct {
	client *dg.Session
	store  *db.Store
	config *c.DiscordConfig
	// maintenance mode
	locked bool
}

func StartBot(config *c.DiscordConfig, store *db.Store) Bot {
	client, _ := dg.New("Bot " + config.Token)

	client.Identify.Intents = dg.MakeIntent(
		dg.IntentsGuildMessages + dg.IntentsGuildMembers,
	)

	bot := Bot{
		client: client,
		store:  store,
		config: config,
		locked: true,
	}

	client.AddHandler(bot.OnMessage)
	client.AddHandlerOnce(bot.OnReady)

	log.Println("Starting to Discord bot...")

	if err := client.Open(); err != nil {
		log.Fatalln("Failed to connect to Discord\n" + err.Error())
	}
	return bot
}

func (bot *Bot) IsAdmin(member *dg.Member) bool {
	for _, roleID := range member.Roles {
		for _, adminID := range bot.config.AdminRoles {
			if adminID == roleID {
				return true
			}
		}
	}
	return false
}

// The first boolean returned is whether they're whitelisted or not.
// the second boolean returned is whether or not they're an administrator
func (bot *Bot) CheckRoles(role []string) (bool, bool) {
	for _, role := range role {
		for _, adminRole := range bot.config.AdminRoles {
			if adminRole == role {
				return true, true
			}
		}
		for _, whitelisted := range bot.config.Whitelist {
			if whitelisted == role {
				return true, false
			}
		}
	}
	return false, false
}

func (bot *Bot) AdminCheck(role []string) bool {
	for _, role := range role {
		for _, adminRole := range bot.config.AdminRoles {
			if adminRole == role {
				return true
			}
		}
	}
	return false
}
