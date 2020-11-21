package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dylhack/mcauth/internal/common"
)

// NoPlayerError means that the request is missing the player UUID attribute.
func NoPlayerError(res http.ResponseWriter) {
	noPlayer := common.ErrorResponse{
		ErrorCode: "NO_PLAYER_ID",
		Message:   "There wasn't a player ID provided",
	}

	ShipError(res, noPlayer)
}

// NoAltNameError means that the request is missing the player name attribute.
func NoAltNameError(res http.ResponseWriter) {
	noAltName := common.ErrorResponse{
		ErrorCode: "MISSING_ALT_NAME",
		Message:   "An alt player name wasn't provided",
	}

	ShipError(res, noAltName)
}

// InvalidAltNameError means the request is has an invalid alt player name.
func InvalidAltNameError(res http.ResponseWriter) {
	invalidAltName := common.ErrorResponse{
		ErrorCode: "INVALID_ALT_NAME",
		Message:   "The alt account name provided is not a valid player name",
	}

	ShipError(res, invalidAltName)
}

// NoOwnerError means the request is missing the owner attribute.
func NoOwnerError(res http.ResponseWriter) {
	noOwner := common.ErrorResponse{
		ErrorCode: "NO_OWNER",
		Message:   "An owner attribute was not provided",
	}

	ShipError(res, noOwner)
}

// InvalidOwnerError means the request gave an invalid owner Minecraft player name.
func InvalidOwnerError(res http.ResponseWriter) {
	invalidOwner := common.ErrorResponse{
		ErrorCode: "INVALID_OWNER",
		Message:   "The owner provided is not a valid player name",
	}

	ShipError(res, invalidOwner)
}

// AltAlreadyExistsError means the alt already exists in the database.
func AltAlreadyExistsError(res http.ResponseWriter) {
	altAlreadyExists := common.ErrorResponse{
		ErrorCode: "ALT_ALREADY_ADDED",
		Message:   "The alt provided is already stored in the database",
	}

	ShipError(res, altAlreadyExists)
}

// MissingIDError occurs when an ID isn't provided in the /api/resolve endpoint
func MissingIDError(res http.ResponseWriter) {
	missingID := common.ErrorResponse{
		ErrorCode: "MISSING_ID",
		Message:   "A resolvable ID is missing in the URL path.",
	}

	ShipError(res, missingID)
}

// InvalidIDError occurs when someone tries to resolve an unresolvable ID
func InvalidIDError(res http.ResponseWriter) {
	invalidID := common.ErrorResponse{
		ErrorCode: "INVALID_ID",
		Message:   "The ID provided isn't in the database or is invalid.",
	}

	ShipError(res, invalidID)
}

// ShipError prepares and sends an error response given.
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
