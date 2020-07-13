package webserver

import (
	"encoding/json"
	"github.com/dhghf/mcauth/internal/bot"
	"github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	bot    *bot.Bot
	store  *db.Store
	config *common.WebServerConfig
}

func StartServer(config *common.WebServerConfig, bot *bot.Bot, store *db.Store) {
	router := mux.NewRouter()
	server := Server{bot: bot, store: store, config: config}

	router.HandleFunc("/isValidPlayer/{playerID}", server.isPlayerValid)

	log.Printf("web-server is listening on port %d\n", config.Port)
	err := http.ListenAndServe(
		":"+strconv.Itoa(config.Port),
		router,
	)

	if err != nil {
		log.Fatalln("Failed to start web-server, is the port available?\n" + err.Error())
	}
}

/* error handlers */
// Missing Player UUID Attribute.
func NoPlayerError(res http.ResponseWriter) {
	noPlayer := common.ErrorResponse{
		ErrorCode: "NO_PLAYER_ID",
		Message:   "There wasn't a player ID provided",
	}
	res.Header().Set("Content-Type", "application/json")
	serialized, err := json.Marshal(noPlayer)

	if err != nil {
		log.Printf("Failed to serialize\n%s\nbecause\n%s\n", noPlayer, err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	_, err = res.Write(serialized)

	if err != nil {
		log.Println("Failed to respond to a request\n" + err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Ship(res http.ResponseWriter, response interface{}) {
	res.Header().Set("Content-Type", "application/json")
	serialized, err := json.Marshal(response)

	if err != nil {
		log.Printf("Failed to serialize\n%s\nbecause\n%s\n", response, err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = res.Write(serialized)

	if err != nil {
		log.Println("Failed to respond to a request\n" + err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
