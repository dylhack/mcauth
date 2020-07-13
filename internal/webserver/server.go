package webserver

import (
	"encoding/json"
	"github.com/dhghf/mcauth/internal/bot"
	"github.com/dhghf/mcauth/internal/common"
	"github.com/dhghf/mcauth/internal/common/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	bot   *bot.Bot
	store *db.Store
}

func StartServer(bot *bot.Bot, store *db.Store) {
	router := mux.NewRouter()
	server := Server{bot: bot, store: store}

	router.HandleFunc("/isValidPlayer/:playerID", server.isPlayerValid)
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

	_, err = res.Write(serialized)

	if err != nil {
		log.Println("Failed to respond to a request\n" + err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
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

	res.WriteHeader(http.StatusOK)
}
