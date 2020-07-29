package db

import (
	"github.com/jinzhu/gorm"
)

// LinksTable - The account links feature mcauth is where the magic happens. Here Minecraft
// player UUID's (without hyphens) are associated with Discord user ID's (Twitter snowflakes).
// Once a user has been authenticated using their authentication code (see authcodes.go) they will
// be stored here with their Minecraft player UUID and Discord user ID on the same row entry
type LinksTable struct {
	gDB *gorm.DB
}

// LinkedAcc represents a linked account for a Discord or Minecraft user.
type LinkedAcc struct {
	// Their Discord user ID (Twitter snowflake)
	DiscordID string `gorm:"column:discord_id;type:text;unique;not null"`
	// Their Minecraft UUID (without hyphens)
	PlayerID string `gorm:"column:player_id;type:text;unique;not null"`
}

// TableName gives GORM the table that linked accounts are stored.
func (LinkedAcc) TableName() string {
	return "account_links"
}

// GetLinksTable will create the account_links table if it doesn't exist.
// it will return LinksTable which can be used to interface with
// the table.
func GetLinksTable(gDB *gorm.DB) LinksTable {
	gDB.AutoMigrate(&LinkedAcc{})

	return LinksTable{
		gDB: gDB,
	}
}

// GetAllLinks will get all the linked accounts in the table.
func (lt *LinksTable) GetAllLinks() (linkedList []LinkedAcc, err error) {
	err = lt.gDB.
		Find(&linkedList).
		Error

	return linkedList, err
}

// SetLink will set a link whether it exists or not.
func (lt *LinksTable) SetLink(discordID, playerID string) error {
	linked := LinkedAcc{
		DiscordID: discordID,
		PlayerID:  playerID,
	}

	return lt.gDB.
		Model(&linked).
		Where("discord_id = ? OR player_id = ?", discordID, playerID).
		Update(discordID, playerID).
		Error
}

// NewLink establishes a new account link.
func (lt *LinksTable) NewLink(discordID string, playerID string) error {
	linked := LinkedAcc{
		DiscordID: discordID,
		PlayerID:  playerID,
	}

	return lt.gDB.
		Create(&linked).
		Error
}

// UnLink will remove a link based on the given identifier.
// The identifier can be either their Discord user ID or
// Minecraft player UUID (without hyphens).
func (lt *LinksTable) UnLink(identifier string) error {
	return lt.gDB.
		Where("player_id = $1 OR discord_id = $1", identifier).
		Delete(&LinkedAcc{}).
		Error
}

// GetPlayerID will get a Discord user's associated Minecraft player UUID (without hyphens).
func (lt *LinksTable) GetPlayerID(discordID string) (playerID string, err error) {
	linked := LinkedAcc{
		PlayerID: "",
	}
	err = lt.gDB.
		First(&linked, "discord_id = ?", discordID).
		Error

	return linked.PlayerID, err
}

// GetDiscordID will get a Minecraft player's associated Discord user ID.
func (lt *LinksTable) GetDiscordID(playerID string) (discordID string, err error) {
	linked := LinkedAcc{
		DiscordID: "",
	}

	err = lt.gDB.
		First(&linked, "player_id = ?", playerID).
		Error

	return linked.DiscordID, err
}
