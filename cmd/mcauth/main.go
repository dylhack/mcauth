package main

import (
	"github.com/dhghf/mcauth/internal/bot"
	"github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"github.com/dhghf/mcauth/internal/webserver"
	"log"
)

const version = "2.0.0"

func main() {
	log.Printf("Starting MCAuth [%s]\n", version)
	config := common.GetConfig("./config.yml")
	store := db.GetStore(config.DB)

	discordBot := bot.StartBot(&config.Discord, &store)
	webserver.StartServer(&config.WebServer, discordBot, &store)
}
