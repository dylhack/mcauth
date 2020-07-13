package bot

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"github.com/dhghf/mcauth/internal/common"
	"log"
)

/* Regular Commands */

func (bot *Bot) authCMD(msg *dg.MessageCreate, args []string) {
	// args = [<prefix>, "auth", <auth code>]

	if len(args) < 3 {
		util.Reply(bot.client, msg.Message,
			fmt.Sprintf("%s auth <authentication code>", bot.config.Prefix),
		)
		return
	}

	// check if they're not already linked with an account
	if account := bot.store.Links.GetPlayerID(msg.Author.ID); len(account) > 0 {
		util.Reply(bot.client, msg.Message, "You're already linked with an account.")
		return
	}

	authCode := args[2]
	if playerID, isOK := bot.store.Auth.Authorize(authCode); isOK {
		err := bot.store.Links.SetLink(msg.Author.ID, playerID)
		if err == nil {
			util.Reply(bot.client, msg.Message, "Linked.")
		} else {
			log.Printf("Something went wrong while linking \"%s\" because \n%s\n",
				msg.Author.ID, err.Error())
		}
	} else {
		util.Reply(bot.client, msg.Message, "Invalid authentication code.")
	}
}

func (bot *Bot) whoAmI(msg *dg.MessageCreate) {
	playerID := bot.store.Links.GetPlayerID(msg.Author.ID)

	if len(playerID) == 0 {
		util.Reply(bot.client, msg.Message, "You aren't linked with any Minecraft accounts.")
		return
	}

	playerName := common.GetPlayerName(playerID)

	if len(playerName) > 0 {
		util.Reply(bot.client, msg.Message, "You are: "+playerName)
	} else {
		util.Reply(bot.client, msg.Message, "I failed to find your associated Minecraft player name")
	}
}

/* Administrator Commands */
