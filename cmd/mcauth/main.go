package main

import (
	"github.com/dhghf/mcauth/internal/bot"
	"github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"github.com/dhghf/mcauth/internal/webserver"
)

func main() {
	config := common.GetConfig("./config.yml")
	store := db.GetStore(config.DB.Location)

	dBot := bot.StartBot(&config.Discord, &store)
	webserver.StartServer(&config.WebServer, &dBot, &store)
}
