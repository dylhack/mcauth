package bot

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"github.com/dhghf/mcauth/internal/common"
	"log"
)

const commands = `These are all the commands:
 - {prefix} auth <authentication code>
 - {prefix} help
 - {prefix} whoami
 - {prefix} whois <player name or @ Discord user>
`

/* Regular Commands */
func (bot *Bot) cmdAuth(msg *dg.MessageCreate, args []string) {
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

func (bot *Bot) cmdWhoAmI(msg *dg.MessageCreate) {
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

func (bot *Bot) cmdWhoIs(msg *dg.MessageCreate, args []string) {
	var playerID, playerName string
	// first let's see if they mentioned a user
	if len(msg.Mentions) > 0 {
		user := msg.Mentions[0]
		playerID = bot.store.Links.GetPlayerID(user.ID)

		if len(playerID) == 0 {
			util.Reply(bot.client, msg.Message, "That user isn't linked with anything")
			return
		}
		playerName = common.GetPlayerName(playerID)

		if len(playerName) == 0 {
			util.Reply(
				bot.client,
				msg.Message,
				"I failed to get the player name but this is the ID"+
					" they're linked with "+playerID,
			)
			return
		}
		util.Reply(
			bot.client,
			msg.Message,
			fmt.Sprintf("%s is %s", user.Mention(), playerName),
		)
		return
	}

	// if they didn't mention a user then check if they're talking a minecraft
	// args = [<prefix>, "whois", <minecraft player name>]
	if len(args) < 3 {
		util.Reply(
			bot.client, msg.Message,
			fmt.Sprintf("%s whois <Minecraft player name>", bot.config.Prefix),
		)
		return
	}

	playerName = args[2]
	playerID = common.GetPlayerID(playerName)

	if len(playerID) == 0 {
		util.Reply(bot.client, msg.Message, "That isn't a Minecraft player")
		return
	}

	userID := bot.store.Links.GetDiscordID(playerID)
	if len(userID) == 0 {
		util.Reply(bot.client, msg.Message, "That user isn't linked with anything")
		return
	}

	util.Reply(
		bot.client,
		msg.Message,
		fmt.Sprintf("%s is <@%s> (%s)", playerName, userID, userID),
	)
}

// there are two ways of unlinking.
// 1. Just by saying "unlink" which will unlink the account associated with your account
// 2. An admin can unlink someone's account
// 2.1 Based on Discord user
// 2.2 Based on Minecraft player name
func (bot *Bot) cmdUnlink(msg *dg.MessageCreate, args []string) {
	var err error
	// 1. Just by saying "unlink" which will unlink the account associated with your account
	// then -> args should be [<prefix>, unlink]
	if len(args) < 3 {
		if err = bot.store.Links.UnLink(msg.Author.ID); err != nil {
			util.Reply(bot.client, msg.Message, "You aren't linked with an account.")
		} else {
			util.Reply(bot.client, msg.Message, "Unlinked.")
		}
		return
	}

	/* 2. An admin can unlink someone's account */
	// then -> args is [<prefix>, unlink, <@Discord User> OR <Minecraft player name>]

	// 2.1 Based on Discord user
	// then -> msg.Mentions should be greater than 0
	if len(msg.Mentions) > 0 {
		user := msg.Mentions[0]
		if err = bot.store.Links.UnLink(user.ID); err != nil {
			util.Reply(bot.client, msg.Message, "That user wasn't linked with any account.")
		} else {
			util.Reply(bot.client, msg.Message, "Unlinked "+user.Mention()+".")
		}
		return
	}

	// 2.2 Based on Minecraft player name
	// then -> args = [<prefix>, unlink, <player name>]
	playerName := args[2]
	playerID := common.GetPlayerID(playerName)

	if len(playerID) == 0 {
		util.Reply(bot.client, msg.Message, playerName+" isn't a Minecraft account.")
		return
	}

	if err = bot.store.Links.UnLink(playerID); err != nil {
		util.Reply(bot.client, msg.Message, "You aren't linked with an account.")
	} else {
		util.Reply(bot.client, msg.Message, "Unlinked "+playerName+".")
	}
}

/* Administrator Commands */
