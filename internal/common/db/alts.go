package db

import (
	"database/sql"
	"log"
)

type AltsTable struct {
	db *sql.DB
}

type AltAcc struct {
	Owner string `json:"alt_owner"`
	PlayerID string `json:"alt_id"`
	PlayerName string `json:"alt_name"`
}

func GetAltsTable(db *sql.DB) AltsTable {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS alts (player_id text UNIQUE NOT NULL, player_name text UNIQUE NOT NULL, owner text NOT NULL)",
	)

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}
	return AltsTable{db: db}
}

func (at *AltsTable) AddAlt(owner string, playerID string, playerName string) error {
	prep, _ := at.db.Prepare(
		"INSERT INTO alts (owner, player_id, player_name) VALUES (?,?,?)",
	)

	_, err := prep.Exec(owner, playerID, playerName)

	if err != nil {
		log.Printf(
			"Failed to insert (%s/%s/%s), because\n%s\n",
			owner,playerID,playerName,err.Error(),
		)
	}

	return err
}

// Identifier can be player name or player ID.
func (at *AltsTable) RemAlt(identifier string) error {
	prep, _ := at.db.Prepare("DELETE FROM alts WHERE player_id=? OR player_name=?")
	_, err := prep.Exec(identifier, identifier)

	if err != nil {
		log.Printf(
			"Failed to delete (%s), because\n%s\n",
			identifier,err.Error(),
		)
	}
	return err
}

// get all the current stored alt accounts.
func (at *AltsTable) GetAlt(identifier string) (result []AltAcc, err error) {
	prep, _ := at.db.Prepare(
		"SELECT * FROM alts (player_name, player_id, owner) WHERE player_id=? OR player_name=?",
	)
	rows, err := prep.Query(identifier, identifier)
	row := AltAcc{}

	if err != nil {
		log.Printf("Failed to get all for \"%s\", because\n%s",
			identifier, err.Error())
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&row.PlayerName, &row.PlayerID, &row.Owner)

		if err != nil {
			log.Printf("Failed to scan \"%s\", because\n%s",
				identifier, err.Error())
			continue
		}
		result = append(result, row)
	}

	return result, nil
}
