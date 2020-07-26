package bot

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (bot *Bot) OnReady(_ *dg.Session, ready *dg.Ready) {
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
	CheckRoles(roles, bot.config.Whitelist)

	// * Check all the roles they gave us through the config exists (administrator roles)
	CheckRoles(roles, bot.config.AdminRoles)

	log.Printf("Ready as %s, serving %s\n", ready.User.Username, guild.Name)
}

func CheckRoles(roles []*dg.Role, given []string) {
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

func (bot *Bot) OnMessage(_ *dg.Session, msg *dg.MessageCreate) {
	// ignore bots and messages that don't start with the prefix
	if msg.Author.Bot || !strings.HasPrefix(msg.Content, bot.config.Prefix) {
		return
	}

	// args = [<prefix>, <sub-command>]
	args := strings.Fields(msg.Content)
	isAdmin := bot.IsAdmin(msg.Member)
	// whether they attempted to run an administrator related command
	adminAttempt := false

	if len(args) < 2 {
		return
	}

	switch args[1] {
	/* User Commands */
	case "auth":
		bot.cmdAuth(msg, args)
		break
	case "whoami":
		bot.cmdWhoAmI(msg)
		break
	case "whois":
		bot.cmdWhoIs(msg, args)
		break
	case "unlink":
		bot.cmdUnlink(msg, args)
		break
	case "commands":
		util.Reply(
			bot.client, msg.Message,
			strings.Replace(commands, "{prefix}", bot.config.Prefix, -1),
		)
		break
	/* Administrator Commands */
	case "status":
		if isAdmin {
			bot.cmdStatus(msg.Message)
		} else {
			adminAttempt = true
		}
		break
	case "lock":
		if isAdmin {
			bot.locked = true
			util.Reply(bot.client, msg.Message, "Maintenance mode is now on.")
		} else {
			adminAttempt = true
		}
		break
	case "unlock":
		if isAdmin {
			bot.locked = false
			util.Reply(bot.client, msg.Message, "Maintenance mode is now off.")
		} else {
			adminAttempt = true
		}
		break
	default:
		util.Reply(bot.client, msg.Message, bot.config.Help)
		break
	}
	if adminAttempt {
		util.Reply(bot.client, msg.Message,
			"You must be an administrator to run this command.",
		)
	}
}

func (bot *Bot) OnGuildMemberUpdate(_ *dg.Session, event *dg.GuildMemberUpdate) {
	bot.Sync(event.User.ID)
}

func (bot *Bot) onGuildMemberAdd(_ *dg.Session, event *dg.GuildMemberAdd) {
	bot.Sync(event.User.ID)
}

func (bot *Bot) onGuildMemberRemove(_ *dg.Session, event *dg.GuildMemberRemove) {
	bot.Sync(event.User.ID)
}
