package webserver

import (
	"github.com/dhghf/mcauth/internal/common"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *Server) isPlayerValid(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)

	playerID, isOK := args["playerID"]

	if !isOK {
		NoPlayerError(res)
		return
	}

	// reason possibilities: NotWhitelisted,
	isValid, reason := s.bot.ValidatePlayer(playerID)
	log.Printf(
		`Validating player "%s"
 - Valid: %t
 - Reason: %s`, playerID, isValid, reason,
	)

	if isValid {
		response := common.ValidPlayer{
			Valid: true,
		}
		Ship(res, response)
		return
	}

	// if the reason was "no link" then check if they have a pending authentication code.
	// If they don't have one then make a new one
	if reason == common.NoLink {
		s.newAuthCode(res, playerID)
		return
	}

	// for other reasons just tell the client what went wrong
	response := common.InvalidPlayer{
		Valid:  false,
		Reason: reason,
	}
	Ship(res, response)
}

func (s *Server) newAuthCode(res http.ResponseWriter, playerID string) {
	authCode, err := s.store.Auth.NewAuthCode(playerID)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := common.InvalidPlayerAuth{
		Reason:   common.AuthCode,
		Valid:    false,
		AuthCode: authCode,
	}
	Ship(res, response)
}
