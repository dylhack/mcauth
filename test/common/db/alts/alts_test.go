package alts

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"os"
	"testing"
	"time"
)

var (
	store      *db.AltsTable
	owner      = "c3b9feea3d7b4ae48d049ea190761877"
	playerName = "LacedLiquid"
	playerID   = "a1ddced8bb20466db456184d9de50346"
)

func TestMain(m *testing.M) {
	dbConfig := db.Config{
		Host:               "",
		Port:               5432,
		User:               "",
		Password:           "",
		Database:           "",
		MaxConnections:     50,
		MaxIdleConnections: 50,
		ConnLifespan:       1 * time.Hour,
	}
	if store == nil {
		storeDB := db.GetStore(dbConfig)
		store = &storeDB.Alts
	}
	m.Run()
}

func TestAlts(t *testing.T) {
	t.Run("AddAlt", testAddAlt)
	t.Run("GetAlt", testGetAlt)
	t.Run("GetAllAlts", testGetAllAlts)
	t.Run("RemAlt", func(t *testing.T) {
		t.Run("RemAlt by Player UUID", testRemAltPlayerID)
		t.Run("Add alt again", testAddAlt)
		t.Run("RemAlt by Player Name", testRemAltPlayerName)
	})
	t.Cleanup(func() { os.Remove("./mcauth-test.db") })
}

func testAddAlt(t *testing.T) {
	err := store.AddAlt(owner, playerID, playerName)

	if err != nil {
		t.Error("AddAlt didn't add the alt account", err)
	}
}

func testGetAlt(t *testing.T) {
	alt, err := store.GetAlt(playerID)

	if err != nil {
		t.Error("GetAlt failed to get the alt because, ", err)
	}

	if alt.PlayerID != playerID {
		t.Errorf("GetAlt failed because of player ID mismatch \"%s\" != \"%s\"\n", playerID, alt.PlayerID)
	}

	if alt.PlayerName != playerName {
		t.Errorf("GetAlt failed because of player name mismatch \"%s\" != \"%s\"\n", playerName, alt.PlayerName)
	}

	if alt.PlayerName != playerName {
		t.Errorf("GetAlt failed because of owner mismatch \"%s\" != \"%s\"\n", owner, alt.Owner)
	}
}

func testGetAllAlts(t *testing.T) {
	alts, err := store.GetAllAlts()

	if err != nil {
		t.Error("GetAllAlts failed because, ", err)
	}

	if len(alts) == 0 {
		t.Error("GetAllAlts returned nothing")
	}

	for _, alt := range alts {
		if alt.PlayerName == playerName {
			if alt.PlayerID == playerID {
				if alt.Owner == owner {
					return
				}
			}
		}
	}
	t.Errorf("GetAllAlts failed because we couldn't find (%s, %s, %s)\n", owner, playerID, playerName)
}

func testRemAltPlayerName(t *testing.T) {
	if err := store.RemAlt(playerName); err != nil {
		t.Error("RemAlt by player name failed because, ", err)
	}
}

func testRemAltPlayerID(t *testing.T) {
	if err := store.RemAlt(playerID); err != nil {
		t.Error("RemAlt by player UUID failed because, ", err)
	}
}
