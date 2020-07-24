// The Minecraft server administrators can use alt account management to claim Minecraft players
// as their alts. This will allow those players to join without authentication. This is super
// useful for mostly alts, but also any player you want to join without question. It will still
// check to see if the owner of the alt is authenticated which can only be done by administrators
// of the server.
package db

import (
	"database/sql"
	"log"
)

// This struct has methods for managing the "alt" table.
type AltsTable struct {
	db   *sql.DB
	fast map[string]*AltAcc // For quickly storing things in memory
}

type AltAcc struct {
	Owner    string `json:"alt_owner"` // The person who claimed the alt
	PlayerID string `json:"alt_id"`    // The ID of the alt account (this will be used for verifying)
	// The name of the alt, this doesn't need to be consistent since it's mostly for listing without
	// reaching out to Mojang's API
	PlayerName string `json:"alt_name"`
}

// This will initialize the table if it doesn't exist. It will then return AltsTable where other
// functions can access this database table.
func GetAltsTable(db *sql.DB) AltsTable {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS alts (
	player_id TEXT UNIQUE NOT NULL, 
	player_name TEXT UNIQUE NOT NULL, 
	owner TEXT NOT NULL)`,
	)

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}

	return AltsTable{
		db:   db,
		fast: make(map[string]*AltAcc),
	}
}

// Add a new alt account.
func (at *AltsTable) AddAlt(owner string, playerID string, playerName string) error {
	prep, err := at.db.Prepare(`
INSERT INTO "alts" (owner, player_id, player_name) 
VALUES ($1, $2, $3)`,
	)

	if err != nil {
		panic(err)
	}

	defer prep.Close()
	_, err = prep.Exec(owner, playerID, playerName)

	if err != nil {
		log.Printf(
			"Failed to insert (%s/%s/%s), because\n%s\n",
			owner, playerID, playerName, err.Error(),
		)
		go at.fastStore(
			playerID,
			&AltAcc{
				Owner:      owner,
				PlayerID:   playerID,
				PlayerName: playerName,
			},
		)
	}

	return err
}

// RemAlt removes an alt account from the table
// identifier can be player name or player ID.
func (at *AltsTable) RemAlt(identifier string) error {
	prep, err := at.db.Prepare(
		`DELETE FROM "alts" 
WHERE player_id   = $1 
   OR player_name = $2`,
	)

	if err != nil {
		panic(err)
	}

	_, err = prep.Exec(identifier, identifier)
	defer prep.Close()

	if err != nil {
		log.Printf(
			"Failed to delete (%s), because\n%s\n",
			identifier, err.Error(),
		)
	} else {
		go at.fastRemove(identifier)
	}
	return err
}

// This will get all the alt accounts in the database.
func (at *AltsTable) GetAllAlts() (result []AltAcc) {
	rows, err := at.db.Query(`SELECT * FROM "alts"`)

	if err != nil {
		panic(err)
	}

	row := AltAcc{}
	result = []AltAcc{}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&row.PlayerID, &row.PlayerName, &row.Owner)

		if err != nil {
			log.Printf("Failed scan an alt of all alts, because\n%s", err.Error())
			continue
		}
		result = append(result, row)
	}

	return result
}

// GetAlt is used by bot/verify.go, it can get an alt account based on a playerID
// but if the alt doesn't exist all the attributes of AltAcc will be empty.
func (at *AltsTable) GetAlt(playerID string) (result AltAcc) {
	altAcc := at.fastLoad(playerID)

	if altAcc != nil {
		return *altAcc
	}

	prep, err := at.db.Prepare(
		`SELECT * FROM "alts" WHERE player_id = $1`,
	)

	if err != nil {
		panic(err)
	}

	defer prep.Close()
	rows, err := prep.Query(playerID)

	if err != nil {
		log.Printf("Failed to get alt \"%s\", because\n%s\n",
			playerID, err.Error())
		return AltAcc{}
	}

	for rows.Next() {
		err = rows.Scan(&result.PlayerID, &result.PlayerName, &result.Owner)

		if err != nil {
			log.Printf("Failed to scan alt \"%s\", because\n%s",
				playerID, err.Error())
			return result
		} else {
			at.fastStore(playerID, &result)
		}
	}
	return result
}

// GetAltsOf will get all the alts associated with an owner (the person who claimed the alts).
func (at *AltsTable) GetAltsOf(owner string) (result []AltAcc, err error) {
	prep, err := at.db.Prepare(
		`SELECT * FROM "alts" WHERE owner = $1`,
	)

	if err != nil {
		panic(err)
	}

	defer prep.Close()

	rows, err := prep.Query(owner)
	row := AltAcc{}
	result = []AltAcc{}

	if err != nil {
		log.Printf("Failed to get all for \"%s\", because\n%s",
			owner, err.Error())
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&row.PlayerID, &row.PlayerName, &row.Owner)

		if err != nil {
			log.Printf("Failed to scan \"%s\", because\n%s",
				owner, err.Error())
			continue
		}
		result = append(result, row)
	}

	return result, nil
}

func (at *AltsTable) fastStore(playerID string, acc *AltAcc) {
	at.fast[playerID] = acc
}

func (at *AltsTable) fastLoad(playerID string) *AltAcc {
	return at.fast[playerID]
}

func (at *AltsTable) fastRemove(playerID string) {
	delete(at.fast, playerID)
}
