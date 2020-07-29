package common

import "github.com/dhghf/mcauth/internal/common/db"

// ErrorResponse represents a generic error
type ErrorResponse struct {
	ErrorCode string `json:"errcode"`
	// a human readable reason why an issue occurred
	Message string `json:"message"`
}

/* isPlayerValid possible responses */

// InvalidPlayer describes why a player isn't valid
type InvalidPlayer struct {
	Verified bool   `json:"valid"`
	Reason   string `json:"reason"`
}

// InvalidPlayerAuth gives the minecraft server the provided player's
// pending authentication code.
type InvalidPlayerAuth struct {
	Reason   string `json:"reason"`
	Verified bool   `json:"valid"`
	AuthCode string `json:"auth_code"`
}

const (
	// NoLink means the player isn't linked with an account.
	NoLink = "no_link"
	// NotWhitelisted means the player's Discord account has missing required roles.
	NotWhitelisted = "no_role"
	// Maintenance means maintenance mode is on.
	Maintenance = "maintenance"
	// AuthCode means the player has a pending authentication code.
	AuthCode = "auth_code"
	//PlayerIsBanned = "banned"
)

// ValidPlayer says that the provided player is valid and ready to join the Minecraft server.
type ValidPlayer struct {
	Verified bool `json:"valid"`
}

// AltsOfResponse has the owner of the alts and an slice of alt accounts.
type AltsOfResponse struct {
	Accounts []db.AltAcc `json:"alt_accs"`
	// The owner's Minecraft player name.
	Owner string `json:"owner"`
}

// AllAltsResponse unlike AltsOfResponse this will show all the alts in the database.
type AllAltsResponse struct {
	Accounts []db.AltAcc `json:"alt_accs"`
}

// NewAltResponse is what the client gives us when a new alt is being claimed.
type NewAltResponse struct {
	// The name of the alt account being claimed
	PlayerName string `json:"player_name"`
	// The owner of the alt account
	Owner string `json:"owner"`
}

// DelAltResponse represents whether or not an alt given was removed.
type DelAltResponse struct {
	// whether or not it was successfully removed
	IsRemoved bool `json:"is_deleted"`
}
