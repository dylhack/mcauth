package routes

import (
	"net/http"

	c "github.com/dylhack/mcauth/internal/common"
	"github.com/dylhack/mcauth/internal/common/db"
	"github.com/gorilla/mux"
)

// Get all the alt accounts associated with an owner.
func (server *Server) getAltsOf(res http.ResponseWriter, req *http.Request) {
	store := server.Store.Alts
	args := mux.Vars(req)
	owner, isOK := args["owner"]

	if !isOK {
		NoOwnerError(res)
		return
	}

	ownerID := c.GetPlayerID(owner)

	if len(ownerID) == 0 {
		InvalidOwnerError(res)
		return
	}

	alts, err := store.GetAltsOf(ownerID)
	var response c.AltsOfResponse

	if err != nil {
		response = c.AltsOfResponse{
			Accounts: []db.AltAcc{},
			Owner:    owner,
		}
	} else {
		response = c.AltsOfResponse{
			Accounts: alts,
			Owner:    owner,
		}
	}
	Ship(res, response)
}

func (server *Server) getAllAlts(res http.ResponseWriter, _ *http.Request) {
	store := server.Store.Alts
	alts, _ := store.GetAllAlts()
	response := c.AllAltsResponse{Accounts: alts}

	Ship(res, response)
}

// NewAlt represents a new incoming alt account.
type NewAlt struct {
	PlayerName string `json:"player_name"`
}

// Create a new alt account.
func (server *Server) postAlt(res http.ResponseWriter, req *http.Request) {
	store := server.Store.Alts
	args := mux.Vars(req)
	ownerName, isOK := args["owner"]

	if !isOK {
		NoOwnerError(res)
		return
	}
	playerName, isOK := args["alt_name"]

	if !isOK {
		NoAltNameError(res)
		return
	}

	// check the owner and playerName already exist
	ownerID := c.GetPlayerID(ownerName)
	playerID := c.GetPlayerID(playerName)

	if len(ownerID) == 0 {
		InvalidOwnerError(res)
		return
	} else if len(playerID) == 0 {
		InvalidAltNameError(res)
		return
	}

	err := store.AddAlt(ownerID, playerID, playerName)

	if err != nil {
		AltAlreadyExistsError(res)
	} else {
		res.WriteHeader(http.StatusOK)
	}
}

// Remove an alt account
func (server *Server) deleteAlt(res http.ResponseWriter, req *http.Request) {
	store := server.Store.Alts
	args := mux.Vars(req)
	playerName, isOK := args["alt_name"]

	if !isOK {
		NoAltNameError(res)
		return
	}

	err := store.RemAlt(playerName)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusOK)
	}
}
