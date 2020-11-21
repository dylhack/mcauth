package routes

import (
	"net/http"
	"strings"

	"github.com/dylhack/mcauth/internal/common"
	"github.com/gorilla/mux"
)

func (server *Server) resolveID(res http.ResponseWriter, req *http.Request) {
	store := &server.Store.Links
	args := mux.Vars(req)
	id, isOK := args["id"]

	if !isOK {
		MissingIDError(res)
		return
	}

	id = strings.ReplaceAll(id, "-", "")
	var resolved string
	var err error

	// UUIDv4's will always be 32 characters
	if len(id) == 32 {
		resolved, err = store.GetDiscordID(id)
	} else { // else it's a Discord Snowflake ID
		resolved, err = store.GetPlayerID(id)
	}

	if err != nil {
		InvalidIDError(res)
		return
	}

	Ship(
		res,
		common.ResolveIDResponse{
			Resolved: resolved,
		},
	)
}
