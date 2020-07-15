package db

import (
	"database/sql"
	"log"
)

type AltsTable struct {
	db   *sql.DB
	fast map[string]*AltAcc
}

type AltAcc struct {
	Owner      string `json:"alt_owner"`
	PlayerID   string `json:"alt_id"`
	PlayerName string `json:"alt_name"`
}

func GetAltsTable(db *sql.DB) AltsTable {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS alts (player_id text UNIQUE NOT NULL, player_name text UNIQUE NOT NULL, owner text NOT NULL)",
	)

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}
	return AltsTable{
		db:   db,
		fast: make(map[string]*AltAcc),
	}
}

func (at *AltsTable) AddAlt(owner string, playerID string, playerName string) error {
	prep, err := at.db.Prepare(
		"INSERT INTO alts (owner, player_id, player_name) VALUES (?,?,?)",
	)

	if err != nil {
		panic(err)
	}

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

// Identifier can be player name or player ID.
func (at *AltsTable) RemAlt(identifier string) error {
	prep, _ := at.db.Prepare("DELETE FROM alts WHERE player_id=? OR player_name=?")
	_, err := prep.Exec(identifier, identifier)

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

func (at *AltsTable) GetAllAlts() (result []AltAcc) {
	rows, err := at.db.Query("SELECT * FROM alts")

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

func (at *AltsTable) GetAlt(playerID string) (result AltAcc) {
	altAcc := at.fastLoad(playerID)

	if altAcc != nil {
		return *altAcc
	}

	prep, err := at.db.Prepare(
		"SELECT * FROM alts WHERE player_id=?",
	)

	if err != nil {
		panic(err)
	}

	rows, err := prep.Query(playerID)

	if err != nil {
		log.Printf("Failed to get alt \"%s\", because\n%s",
			playerID, err.Error())
		return AltAcc{}
	}

	for rows.Next() {
		err = rows.Scan(&result.PlayerID, &result.PlayerName, &result.Owner)

		if err != nil {
			log.Printf("Failed to scan alt \"%s\", because\n%s",
				playerID, err.Error())
			continue
		}
		go at.fastStore(playerID, &result)
		return result
	}
	return AltAcc{}
}

// get all the current stored alt accounts of an owner.
func (at *AltsTable) GetAltsOf(owner string) (result []AltAcc, err error) {
	prep, err := at.db.Prepare(
		"SELECT * FROM alts WHERE owner=?",
	)

	if err != nil {
		panic(err)
	}

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
