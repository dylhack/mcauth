package common

import "github.com/dhghf/mcauth/internal/common/db"

// generic error
type ErrorResponse struct {
	ErrorCode string `json:"errcode"`
	// a human readable reason why an issue occurred
	Message string `json:"message"`
}

/* isPlayerValid possible responses */

type PlayerValidation struct {
	Valid bool `json:"valid"`
}

// this describes why a player isn't valid
type InvalidPlayer struct {
	PlayerValidation
	Reason string `json:"reason"`
}

// this gives the minecraft server the provided player's pending authentication code.
type InvalidPlayerAuth struct {
	InvalidPlayer
	AuthCode string `json:"auth_code"`
}

// this says that the provided player is valid and ready to join the Minecraft server.
type ValidPlayer struct {
	PlayerValidation
}

/* getAltsOf possible responses */
type AltsOfResponse struct {
	Accounts []db.AltAcc `json:"alt_accs"`
}

/* newAlt possible responses */
type NewAltResponse struct {
	// The name of the alt account being claimed
	PlayerName string `json:"player_name"`
	// The owner of the alt account
	Owner string `json:"owner"`
}

/* delAlt possible responses */
type DelAltResponse struct {
	// whether or not it was successfully removed
	IsRemoved bool `json:"is_deleted"`
}
