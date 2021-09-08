package routes

import (
	"net/http"

	"github.com/dylhack/mcauth/internal/common"
	"github.com/gorilla/mux"
)

func (server *Server) details(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)

	playerID, isOK := args["playerID"]

	if !isOK {
		NoPlayerError(res)
		return
	}

	// reason possibilities: NotWhitelisted,
	resp := server.Bot.PlayerDetails(playerID)

	// if the reason was "no link" then check if they have a pending authentication code.
	// If they don't have one then make a new one
	if resp.State == common.NoLink {
		store := server.Store.Auth
		authCode, _ := store.NewAuthCode(playerID)
		resp.AuthCode = authCode
		resp.State = common.AuthCode
	}

	Ship(res, resp)
}
