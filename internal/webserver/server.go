package webserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dylhack/mcauth/internal/bot"
	"github.com/dylhack/mcauth/internal/common"
	"github.com/dylhack/mcauth/internal/common/db"
	"github.com/dylhack/mcauth/internal/webserver/routes"
	"github.com/gorilla/mux"
)

// StartServer is for starting the HTTP webserver which communicates with the plugin
// MCAuth Client.
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
