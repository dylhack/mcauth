// authcodes.go is responsible for managing authentication codes
package db

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strings"
)

type AuthTable struct {
	db *sql.DB
}

func GetAuthTable(db *sql.DB) AuthTable {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth_codes (auth_code TEXT UNIQUE NOT NULL, player_id TEXT PRIMARY KEY NOT NULL)")

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}
	return AuthTable{db: db}
}

// create a new authorization code. If the player already has an authorization code it will be
// returned rather than making a new one.
func (at *AuthTable) NewAuthCode(playerID string) (authCode string, err error) {
	oldAuthCode := at.GetAuthCode(playerID)

	if len(oldAuthCode) > 0 {
		return oldAuthCode, nil
	}
	newUUID := uuid.New()
	authCode = strings.Split(newUUID.String(), "-")[0]
	prep, _ := at.db.Prepare("INSERT INTO auth_codes (auth_code, player_id) VALUES (?,?)")

	_, err = prep.Exec(authCode, playerID)

	if err != nil {
		return "", err
	} else {
		return authCode, nil
	}
}

func (at *AuthTable) GetAuthCode(playerID string) (authCode string) {
	prep, _ := at.db.Prepare("SELECT auth_code WHERE player_id=?")
	rows, err := prep.Query(playerID)

	if err != nil {
		return ""
	}

	for rows.Next() {
		err = rows.Scan(&authCode)

		if err != nil {
			return ""
		} else {
			return authCode
		}
	}
	return ""
}

// Authorize a given authentication code. It will return the player ID associated with the given
// auth code.
func (at *AuthTable) Authorize(authCode string) (playerID string) {
	playerID = at.GetPlayerID(authCode)

	// see if they have an authentication code
	if len(playerID) > 0 {
		// remove them from the database
		go at.RemoveCode(authCode)

		return playerID
	} else {
		return ""
	}
}

// Get the player ID associated with the given authentication code.
func (at *AuthTable) GetPlayerID(authCode string) (playerID string) {
	prep, _ := at.db.Prepare("SELECT player_id WHERE auth_code=?")
	rows, _ := prep.Query(authCode)

	for rows.Next() {
		err := rows.Scan(&playerID)

		if err != nil {
			return ""
		} else {
			return playerID
		}
	}
	return ""
}

func (at *AuthTable) RemoveCode(authCode string) {
	prep, _ := at.db.Prepare("DELETE FROM auth_codes WHERE auth_code=?")

	_, err := prep.Exec()

	if err != nil {
		log.Printf(
			"Failed to remove \"%s\", because\n%s",
			authCode,
			err.Error(),
		)
	}
}

