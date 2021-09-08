package bot

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (bot *Bot) onReady(_ *dg.Session, ready *dg.Ready) {
	// let's do a couple things:
	// * Check that we're on the Discord server that we're serving.
	// * Check all the roles they gave us through the config exists
	roles, err := bot.client.GuildRoles(bot.config.Guild)
	guild, _ := bot.client.Guild(bot.config.Guild)

	// * Check that we're on the Discord server that we're serving.
	if err != nil {
		log.Fatalf("Failed to get the roles for \"%s\", because\n"+err.Error(), bot.config.Guild)
	}

	// * Check all the roles they gave us through the config exists (whitelist first)
	verifyRoles(roles, bot.config.Whitelist)

	// * Check all the roles they gave us through the config exists (administrator roles)
	verifyRoles(roles, bot.config.AdminRoles)

	log.Printf("Ready as %s, serving %s\n", ready.User.Username, guild.Name)
}

func verifyRoles(roles []*dg.Role, given []string) {
	var found bool
	for _, givenID := range given {
		found = false
		for _, role := range roles {
			if role.ID == givenID {
				found = true
				break
			}
		}
		if !found {
			log.Println("Failed to find role: " + givenID)
		}
	}
}

func (bot *Bot) onMessage(_ *dg.Session, msg *dg.MessageCreate) {
	// ignore bots and messages that don't start with the prefix
	if msg.Author.Bot || !strings.HasPrefix(msg.Content, bot.config.Prefix) {
		return
	}

	// args = [<prefix>, <sub-command>]
	args := strings.Fields(msg.Content)

	if len(args) < 2 {
		return
	}

	switch args[1] {
	/* User Commands */
	case "auth":
		bot.cmdAuth(msg, args)
	case "whoami":
		bot.cmdWhoAmI(msg)
	case "whois":
		bot.cmdWhoIs(msg, args)
	case "commands":
		_, _ = util.Reply(
			bot.client, msg.Message,
			strings.Replace(commands, "{prefix}", bot.config.Prefix, -1),
		)
	default:
		bot.onAdminCommand(msg.Message, args)
	}
}

func (bot *Bot) onAdminCommand(msg *dg.Message, args []string) {
	isAdmin := bot.isAdmin(msg.Member)

	adminAttempt := true

	switch args[1] {
	/* Administrator Commands */
	case "unlink":
		if isAdmin {
			bot.cmdUnlink(msg, args)
		}
	case "status":
		if isAdmin {
			bot.cmdStatus(msg)
		}
	case "lock":
		if isAdmin {
			bot.cmdSetLock(msg, true)
		}
	case "unlock":
		if isAdmin {
			bot.cmdSetLock(msg, false)
		}
	case "public":
		if isAdmin {
			bot.cmdSetMode(msg, false)
		}
	case "private":
		if isAdmin {
			bot.cmdSetMode(msg, true)
		}
	default:
		_, _ = util.Reply(bot.client, msg, bot.config.Help)
		adminAttempt = false
	}

	if adminAttempt && !isAdmin {
		_, _ = util.Reply(bot.client, msg,
			"You must be an administrator to run this command.",
		)
	}
}

func (bot *Bot) onGuildMemberUpdate(_ *dg.Session, event *dg.GuildMemberUpdate) {
	bot.syncMember(event.User.ID)
}

func (bot *Bot) onGuildMemberAdd(_ *dg.Session, event *dg.GuildMemberAdd) {
	bot.syncMember(event.User.ID)
}

func (bot *Bot) onGuildMemberRemove(_ *dg.Session, event *dg.GuildMemberRemove) {
	bot.syncMember(event.User.ID)
}
