package bot

import (
	dg "github.com/bwmarrin/discordgo"
	c "github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"log"
	"strings"
)

type Bot struct {
	client *dg.Session
	store  *db.Store
	config *c.DiscordConfig
	sync   SyncHandler
	// maintenance mode
	locked bool
}

func StartBot(config *c.DiscordConfig, store *db.Store) *Bot {
	client, _ := dg.New("Bot " + config.Token)

	client.Identify.Intents = dg.MakeIntent(
		dg.IntentsGuildMessages + dg.IntentsGuildMembers,
	)

	bot := &Bot{
		client: client,
		store:  store,
		config: config,
		locked: false,
		sync:   GetSyncHandler(),
	}

	client.AddHandler(bot.OnMessage)
	client.AddHandler(bot.OnGuildMemberUpdate)
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
func (bot *Bot) CheckRoles(roles []string) (isWhitelisted, isAdmin bool) {
	isWhitelisted = false
	isAdmin = false
	for _, role := range roles {
		for _, adminRole := range bot.config.AdminRoles {
			if adminRole == role {
				isWhitelisted = true
				isAdmin = true
				return isWhitelisted, isAdmin
			}
		}
		for _, whitelisted := range bot.config.Whitelist {
			if whitelisted == role {
				isWhitelisted = true
				isAdmin = false
			}
		}
	}
	return isWhitelisted, isAdmin
}

func (bot *Bot) Sync(memberID string) {
	log.Printf("Syncing roles for \"%s\"\n", memberID)
	member, err := bot.client.GuildMember(bot.config.Guild, memberID)

	if err != nil {
		if strings.Contains(err.Error(), "Unknown Member") {
			bot.sync.SyncRoles(memberID, &[]string{})
		} else {
			log.Fatalf(
				"Failed to fetch member \"%s\", please investigate \n%s",
				memberID, err,
			)
		}
	} else {
		bot.sync.SyncRoles(memberID, &member.Roles)
	}
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
