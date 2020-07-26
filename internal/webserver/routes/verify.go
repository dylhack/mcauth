package routes

import (
	"github.com/dhghf/mcauth/internal/common"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (server *Server) verifyPlayer(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)

	playerID, isOK := args["playerID"]

	if !isOK {
		NoPlayerError(res)
		return
	}

	// reason possibilities: NotWhitelisted,
	isValid, reason := server.Bot.VerifyPlayer(playerID)
	log.Printf(
		`Validating player "%s"
 - Verified: %t
 - Reason: %s`, playerID, isValid, reason,
	)

	if isValid {
		response := common.ValidPlayer{
			Verified: true,
		}
		Ship(res, response)
		return
	}

	// if the reason was "no link" then check if they have a pending authentication code.
	// If they don't have one then make a new one
	if reason == common.NoLink {
		server.newAuthCode(res, playerID)
		return
	}

	// for other reasons just tell the client what went wrong
	response := common.InvalidPlayer{
		Verified: false,
		Reason:   reason,
	}
	Ship(res, response)
}

func (server *Server) newAuthCode(res http.ResponseWriter, playerID string) {
	store := server.Store.Auth
	authCode, _ := store.NewAuthCode(playerID)

	response := common.InvalidPlayerAuth{
		Reason:   common.AuthCode,
		Verified: false,
		AuthCode: authCode,
	}
	Ship(res, response)
}
