// authcodes.go is responsible for managing authentication codes
package db

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strings"
)

type AuthTable struct {
	db   *sql.DB
	fast map[string]string
}

func GetAuthTable(db *sql.DB) AuthTable {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth_codes (auth_code TEXT UNIQUE NOT NULL, player_id TEXT PRIMARY KEY NOT NULL)")

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}
	return AuthTable{
		db:   db,
		fast: map[string]string{},
	}
}

// create a new authorization code. If the player already has an authorization code it will be
// returned rather than making a new one.
func (at *AuthTable) NewAuthCode(playerID string) (authCode string) {
	oldAuthCode := at.GetAuthCode(playerID)

	if len(oldAuthCode) > 0 {
		return oldAuthCode
	}

	newUUID := uuid.New()
	authCode = strings.Split(newUUID.String(), "-")[0]

	prep, _ := at.db.Prepare("INSERT INTO auth_codes (auth_code, player_id) VALUES (?,?)")

	_, err := prep.Exec(authCode, playerID)

	defer prep.Close()

	if err != nil {
		log.Printf("Failed to store (%s/%s), because\n%s",
			playerID, authCode, err.Error())
	} else {
		at.fastStore(playerID, authCode)
	}

	return authCode
}

func (at *AuthTable) GetAuthCode(playerID string) (authCode string) {
	authCode, isOK := at.fastLoad(playerID)
	if isOK {
		return authCode
	}

	prep, _ := at.db.Prepare("SELECT auth_code FROM auth_codes WHERE player_id=?")
	rows, err := prep.Query(playerID)

	if err != nil {
		return ""
	}

	defer prep.Close()
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&authCode)

		if err != nil {
			return ""
		} else {
			at.fastStore(playerID, authCode)
			return authCode
		}
	}
	return ""
}

// Authorize a given authentication code. It will return the player ID associated with the given
// auth code.
func (at *AuthTable) Authorize(authCode string) (string, bool) {
	playerID := at.GetPlayerID(authCode)

	// see if they have an authentication code
	if len(playerID) > 0 {
		// remove them from the database
		go at.RemoveCode(authCode)
		go at.fastRemove(authCode)

		return playerID, true
	} else {
		return "", false
	}
}

// Get the player ID associated with the given authentication code.
func (at *AuthTable) GetPlayerID(authCode string) (playerID string) {
	playerID, isOK := at.fastLoad(authCode)

	if isOK {
		return playerID
	}

	prep, _ := at.db.Prepare("SELECT player_id FROM auth_codes WHERE auth_code=?")
	rows, _ := prep.Query(authCode)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&playerID)

		if err != nil {
			return ""
		} else {
			at.fastStore(playerID, authCode)
			return playerID
		}
	}
	return ""
}

func (at *AuthTable) RemoveCode(authCode string) {
	prep, err := at.db.Prepare("DELETE FROM auth_codes WHERE auth_code=?")
	if err != nil {
		panic(err)
	}

	defer prep.Close()
	_, err = prep.Exec(authCode)

	if err != nil {
		log.Printf(
			"Failed to remove \"%s\", because\n%s",
			authCode,
			err.Error(),
		)
	}
}

func (at *AuthTable) fastStore(authCode string, playerID string) {
	at.fast[authCode] = playerID
	at.fast[playerID] = authCode
}

func (at *AuthTable) fastRemove(identifier string) {
	authCode, isOK := at.fast[identifier]
	if isOK {
		delete(at.fast, authCode)
		delete(at.fast, identifier)
	} else {
		playerID, isOK := at.fast[identifier]
		if isOK {
			delete(at.fast, identifier)
			delete(at.fast, playerID)
		}
	}
}

func (at *AuthTable) fastLoad(identifier string) (string, bool) {
	result, isOK := at.fast[identifier]
	return result, isOK
}
