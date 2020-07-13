package bot

import (
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

}
