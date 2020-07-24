// The account links feature mcauth is where the magic happens. Here Minecraft player UUID's
//(without hyphens) are associated with Discord user ID's (Twitter snowflakes). Once a user has
// been authenticated using their authentication code (see authcodes.go) they will be stored here
// with their Minecraft player UUID and Discord user ID on the same row entry
package db

import (
	"database/sql"
	"log"
)

type LinksTable struct {
	db   *sql.DB
	fast map[string]string // For quickly storing and getting accounts from memory
}

// This represents a linked account for a user
type LinkedAcc struct {
	DiscordID string // Their Discord user ID (Twitter snowflake)
	PlayerID  string // Their Minecraft UUID (without hyphens)
}

// This will create teh account_links table if it doesn't exist.
// it will return LinksTable which can be used to interface with
// the table.
func GetLinksTable(db *sql.DB) LinksTable {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS account_links (
    discord_id TEXT UNIQUE NOT NULL, 
    player_id TEXT UNIQUE NOT NULL)`,
	)

	if err != nil {
		log.Fatalln("Failed to init authentication table\n" + err.Error())
	}
	return LinksTable{
		db:   db,
		fast: map[string]string{},
	}
}

// This will get all the linked accounts in the table.
func (lt *LinksTable) GetAllLinks() (linkedList []LinkedAcc) {
	prep, err := lt.db.Prepare(`SELECT * FROM "account_links"`)
	if err != nil {
		panic(err)
	}

	rows, err := prep.Query()

	if err != nil {
		log.Println("Failed to get all account links because ", err.Error())
		return linkedList
	}
	prep.Close()
	defer rows.Close()

	for rows.Next() {
		var linked LinkedAcc

		err = rows.Scan(&linked.DiscordID, &linked.PlayerID)

		if err != nil {
			log.Printf("Failed to scan an account because %s\n", err.Error())
		} else {
			linkedList = append(linkedList, linked)
		}
	}

	return linkedList
}

// This will set a link whether it exists or not.
func (lt *LinksTable) SetLink(discordID, playerID string) error {
	// check if they already have a link
	oldID := lt.GetPlayerID(discordID)

	if len(oldID) > 0 {
		prep, err := lt.db.Prepare(`
UPDATE account_links 
SET discord_id = $1, 
    player_id  = $2 
WHERE discord_id = $1 
   OR player_id  = $2`,
		)

		if err != nil {
			panic(err)
		}
		defer prep.Close()

		_, err = prep.Exec(discordID, playerID)

		if err != nil {
			log.Printf(
				"Failed to set (%s/%s), because\n%s\n",
				discordID, playerID, err.Error(),
			)
		} else {
			go lt.fastStore(playerID, discordID)
		}

		return err
	} else {
		return lt.NewLink(discordID, playerID)
	}
}

// This establishes a new account link.
func (lt *LinksTable) NewLink(discordID string, playerID string) error {
	prep, err := lt.db.Prepare(`
INSERT INTO account_links (discord_id, player_id) 
VALUES ($1,$2)`,
	)

	if err != nil {
		panic(err)
	}
	defer prep.Close()

	_, err = prep.Exec(discordID, playerID)

	if err != nil {
		log.Printf(
			"Failed to insert (%s/%s), because\n%s\n",
			discordID, playerID, err.Error(),
		)
	} else {
		go lt.fastStore(playerID, discordID)
	}

	return err
}

// This will remove a link based on the given identifier.
// The identifier can be either their Discord user ID or
// Minecraft player UUID (without hyphens).
func (lt *LinksTable) UnLink(identifier string) error {
	prep, err := lt.db.Prepare(`
DELETE FROM account_links 
WHERE discord_id = $1 
   OR player_id  = $1`)

	if err != nil {
		panic(err)
	}

	defer prep.Close()

	_, err = prep.Exec(identifier)

	if err != nil {
		log.Printf(
			"Failed to remove (%s), because\n%s\n",
			identifier, err.Error(),
		)
	} else {
		prep.Close()
		lt.fastRemove(identifier)
	}

	return err
}

// This will get a Discord user's associated Minecraft player UUID (without hyphens).
func (lt *LinksTable) GetPlayerID(discordID string) (playerID string) {
	playerID, isOK := lt.fastLoad(discordID)

	if isOK {
		return playerID
	}

	prep, err := lt.db.Prepare(
		`SELECT player_id FROM account_links 
WHERE discord_id = $1`,
	)

	if err != nil {
		panic(err)
	}
	defer prep.Close()

	rows, err := prep.Query(discordID)

	if err != nil {
		log.Printf("Failed to get \"%s\"'s player ID, because\n%s\n", discordID, err.Error())
		return ""
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&playerID)

		if err != nil {
			log.Printf(
				"Failed to get \"%s\"'s player ID, because\n%s\n",
				discordID,
				err.Error(),
			)
			return ""
		} else {
			go lt.fastStore(playerID, discordID)
			return playerID
		}
	}
	return ""
}

// This will get a Minecraft player's associated Discord user ID.
func (lt *LinksTable) GetDiscordID(playerID string) (discordID string) {
	discordID, isOK := lt.fastLoad(playerID)

	if isOK {
		return discordID
	}

	prep, err := lt.db.Prepare(
		`SELECT discord_id FROM account_links WHERE player_id = $1`,
	)
	if err != nil {
		panic(err)
	}

	rows, err := prep.Query(playerID)

	defer prep.Close()

	if err != nil {
		log.Printf("Failed to get \"%s\"'s player ID, because\n%s\n", playerID, err.Error())
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&discordID)

		if err != nil {
			log.Printf(
				"Failed to get \"%s\"'s player ID, because\n%s\n",
				playerID,
				err.Error(),
			)
			return ""
		} else {
			go lt.fastStore(discordID, playerID)
			return discordID
		}
	}
	return ""
}

func (lt *LinksTable) fastStore(playerID string, discordID string) {
	lt.fast[playerID] = discordID
	lt.fast[discordID] = playerID
}

func (lt *LinksTable) fastRemove(identifier string) {
	discordID, isOK := lt.fast[identifier]
	if isOK {
		delete(lt.fast, discordID)
		delete(lt.fast, identifier)
	} else {
		playerID, isOK := lt.fast[identifier]
		if isOK {
			delete(lt.fast, identifier)
			delete(lt.fast, playerID)
		}
	}
}

func (lt *LinksTable) fastLoad(identifier string) (string, bool) {
	result, isOK := lt.fast[identifier]
	return result, isOK
}
