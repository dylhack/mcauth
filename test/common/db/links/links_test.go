package links

import (
	"github.com/dhghf/mcauth/internal/common/db"
	db2 "github.com/dhghf/mcauth/test/common/db"
	"testing"
)

var (
	playerUUID = "5bce3068e4f3489fb66b5723b2a7cdb1"
	discordID  = "164837347156951040"
	store      *db.LinksTable
)

func TestMain(m *testing.M) {
	if store == nil {
		storeDB := db.GetStore(db2.TestConfig)
		store = &storeDB.Links
	}
	m.Run()
}

func TestLinks(t *testing.T) {
	t.Run("NewLink", testNewLink)
	t.Run("SetLink", testSetLink)
	t.Run("GetDiscordID", testGetDiscordID)
	t.Run("GetPlayerUUID", testGetPlayerUUID)
	t.Run("GetAllLinks", testGetAllLinks)
	t.Run("Unlink", func(t *testing.T) {
		t.Run("UnLinkPlayerUUID", testUnLinkPlayerUUID)
		t.Run("Re-link", testNewLink)
		t.Run("UnLinkDiscordID", testUnLinkDiscordID)
	})
}

func testNewLink(t *testing.T) {
	err := store.NewLink(discordID, playerUUID)

	if err != nil {
		t.Error(err)
	}
}

func testSetLink(t *testing.T) {
	err := store.SetLink(discordID, playerUUID)

	if err != nil {
		t.Error(err)
	}
}

func testGetPlayerUUID(t *testing.T) {
	result, err := store.GetPlayerID(discordID)

	if err != nil {
		t.Error("GetPlayerID failed because, ", err)
	}

	if playerUUID != result {
		t.Errorf("GetPlayerID failed because \"%s\" != \"%s\"\n", playerUUID, result)
	}
}

func testGetDiscordID(t *testing.T) {
	result, err := store.GetDiscordID(playerUUID)

	if err != nil {
		t.Error("GetDiscordID failed because, ", err)
	}

	if discordID != result {
		t.Errorf("GetDiscordID failed because \"%s\" != \"%s\"\n", discordID, result)
	}
}

func testGetAllLinks(t *testing.T) {
	result, err := store.GetAllLinks()

	if err != nil {
		t.Error("GetAllLinks failed because, ", err)
	}

	if len(result) == 0 {
		t.Error("GetAllLinks returned nothing")
	}

	for _, accLink := range result {
		if accLink.PlayerID == playerUUID && accLink.DiscordID == discordID {
			return
		}
	}
	t.Errorf("Could find link for (%s/%s)\n", playerUUID, discordID)
}

func testUnLinkDiscordID(t *testing.T) {
	err := store.UnLink(discordID)

	if err != nil {
		t.Error(err)
	}
}

func testUnLinkPlayerUUID(t *testing.T) {
	err := store.UnLink(playerUUID)

	if err != nil {
		t.Error(err)
	}
}
