package main

import (
	"log"

	"github.com/dylhack/mcauth/internal/bot"
	"github.com/dylhack/mcauth/internal/common"
	"github.com/dylhack/mcauth/internal/common/db"
	"github.com/dylhack/mcauth/internal/webserver"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	log.Printf("Starting MCAuth [%s]\n", common.Version)
	config := common.GetConfig("./config.yml")
	store := db.GetStore(config.DB)

	discordBot := bot.StartBot(&config.Discord, &store)
	webserver.StartServer(&config.WebServer, discordBot, &store)
}
