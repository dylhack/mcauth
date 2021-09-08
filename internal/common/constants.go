package common

// Version is the MCAuth version.

const (
	// Maintenance means maintenance mode is on.
	Maintenance = "maintenance"

	// MCAuth Version
	Version = "3.0.0"
)

// States of a player
const (
	// The player is an admin.
	IsAdmin = "admin"
	// AuthCode means the player has a pending authentication code.
	AuthCode = "auth_code"
	// NoLink means the player isn't linked with an account.
	NoLink = "no_link"
	// NotWhitelisted means the player's Discord account
	// has missing required roles.
	NotWhitelisted = "no_role"
	// The player is an alt of an admin account.
	IsAlt = "alt_acc"
	// The player is verified.
	Whitelisted = "whitelisted"
	//PlayerIsBanned = "banned"
)
