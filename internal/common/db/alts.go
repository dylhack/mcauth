package db

import (
	"github.com/jinzhu/gorm"
)

// AltsTable - The Minecraft server administrators can use alt account management to claim Minecraft
// players as their alts. This will allow those players to join without authentication. This is
// super useful for mostly alts, but also any player you want to join without question. It will
// still check to see if the owner of the alt is authenticated which can only be done by
// administrators of the server.
type AltsTable struct {
	gDB *gorm.DB
}

// AltAcc represents an alt account claimed by an admin.
type AltAcc struct {
	// The person who claimed the alt
	Owner string `json:"alt_owner" gorm:"column:owner;type:text;unique;not null"`
	// The ID of the alt account (this will be used for verifying)
	PlayerID string `json:"alt_id" gorm:"column:player_id;type:text;primary_key;not null"`
	// The name of the alt, this doesn't need to be consistent since it's mostly for listing without
	// reaching out to Mojang's API
	PlayerName string `json:"alt_name" gorm:"column:player_name;type:text;unique;not null"`
}

// TableName gives GORM the table name for account management.
func (AltAcc) TableName() string {
	return "alt_accounts"
}

// GetAltsTable will initialize the table if it doesn't exist. It will then return AltsTable
// where other functions can access this database table.
func GetAltsTable(gDB *gorm.DB) AltsTable {
	gDB.AutoMigrate(&AltAcc{})

	return AltsTable{
		gDB: gDB,
	}
}

// AddAlt adds a new alt account.
// owner must be the owner player name.
// playerID must be the player's UUID.
// playerName must be the player's in-game name.
func (at *AltsTable) AddAlt(owner string, playerID string, playerName string) error {
	altAcc := AltAcc{
		Owner:      owner,
		PlayerID:   playerID,
		PlayerName: playerName,
	}

	return at.gDB.
		Create(&altAcc).
		Error
}

// RemAlt removes an alt account from the table
// identifier can be player name or player ID.
func (at *AltsTable) RemAlt(identifier string) error {
	return at.gDB.
		Where("player_name = ? OR player_id = ?", identifier, identifier).
		Delete(AltAcc{}).
		Error
}

// GetAllAlts will get all the alt accounts in the database.
func (at *AltsTable) GetAllAlts() (result []AltAcc, err error) {
	err = at.gDB.
		Find(&result).
		Error
	return result, err
}

// GetAlt is used by bot/verify.go, it can get an alt account based on a playerID
// but if the alt doesn't exist all the attributes of AltAcc will be empty.
func (at *AltsTable) GetAlt(playerID string) (result AltAcc, err error) {
	err = at.gDB.
		First(&result, "player_id = ?", playerID).
		Error
	return result, err
}

// GetAltsOf will get all the alts associated with an owner (the person who claimed the alts).
func (at *AltsTable) GetAltsOf(owner string) (result []AltAcc, err error) {
	err = at.gDB.
		Where("owner = ?", owner).Find(&result).
		Error
	return result, err
}
