package routes

import (
	"encoding/json"
	"github.com/dhghf/mcauth/internal/common"
	"log"
	"net/http"
)

/* error handlers */
// Missing Player UUID Attribute.
func NoPlayerError(res http.ResponseWriter) {
	noPlayer := common.ErrorResponse{
		ErrorCode: "NO_PLAYER_ID",
		Message:   "There wasn't a player ID provided",
	}

	ShipError(res, noPlayer)
}

/* error handlers - alt account management */

// Missing Player Name Attribute.
func NoAltNameError(res http.ResponseWriter) {
	noAltName := common.ErrorResponse{
		ErrorCode: "MISSING_ALT_NAME",
		Message:   "An alt player name wasn't provided",
	}

	ShipError(res, noAltName)
}

// Invalid Alt Name.
func InvalidAltNameError(res http.ResponseWriter) {
	invalidAltName := common.ErrorResponse{
		ErrorCode: "INVALID_ALT_NAME",
		Message:   "The alt account name provided is not a valid player name",
	}

	ShipError(res, invalidAltName)
}

// Missing Owner Attribute.
func NoOwnerError(res http.ResponseWriter) {
	noOwner := common.ErrorResponse{
		ErrorCode: "NO_OWNER",
		Message:   "An owner attribute was not provided",
	}

	ShipError(res, noOwner)
}

// Invalid Owner.
func InvalidOwnerError(res http.ResponseWriter) {
	invalidOwner := common.ErrorResponse{
		ErrorCode: "INVALID_OWNER",
		Message:   "The owner provided is not a valid player name",
	}

	ShipError(res, invalidOwner)
}

// Alt Already Added.
func AltAlreadyExistsError(res http.ResponseWriter) {
	altAlreadyExists := common.ErrorResponse{
		ErrorCode: "ALT_ALREADY_ADDED",
		Message:   "The alt provided is already stored in the database",
	}

	ShipError(res, altAlreadyExists)
}

func ShipError(res http.ResponseWriter, response interface{}) {
	res.Header().Set("Content-Type", "application/json")
	serialized, err := json.Marshal(response)

	if err != nil {
		log.Printf("Failed to serialize\n%s\nbecause\n%s\n", response, err.Error())
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
