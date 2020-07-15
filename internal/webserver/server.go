package webserver

import (
	"github.com/dhghf/mcauth/internal/bot"
	"github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"github.com/dhghf/mcauth/internal/webserver/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartServer(config *common.WebServerConfig, bot *bot.Bot, store *db.Store) {
	router := mux.NewRouter()

	// Start routes
	routes.StartAllRoutes(bot, store, config, router)

	log.Printf("web-server is listening on port %d\n", config.Port)
	err := http.ListenAndServe(
		":"+strconv.Itoa(config.Port),
		router,
	)

	if err != nil {
		log.Fatalln("Failed to start web-server, is the port available?\n" + err.Error())
	}
}
