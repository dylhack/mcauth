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

	client.AddHandler(bot.OnReady)
	client.AddHandler(bot.OnMessage)

	log.Println("Connecting to Discord...")

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
