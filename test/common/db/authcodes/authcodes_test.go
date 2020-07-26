package authcodes

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"os"
	"testing"
	"time"
)

var playerUUID = "5bce3068e4f3489fb66b5723b2a7cdb1"
var authCode string
var store *db.AuthTable

func TestMain(m *testing.M) {
	dbConfig := db.Config{
		Host:               "localhost",
		Port:               5432,
		User:               "",
		Password:           "",
		Database:           "mcauth",
		MaxConnections:     50,
		MaxIdleConnections: 50,
		ConnLifespan:       1 * time.Hour,
	}
	if store == nil {
		storeDB := db.GetStore(dbConfig)
		store = &storeDB.Auth
	}
	m.Run()
}

func TestAuthCodes(t *testing.T) {
	t.Run("NewAuthCode", testNewAuthCode)
	t.Run("GetAuthCode", testGetAuthCode)
	t.Run("GetPlayerID", testGetPlayerID)
	t.Run("Authorize", testAuthorize)
	t.Cleanup(func() { os.Remove("./mcauth-test.db") })
}

func testNewAuthCode(t *testing.T) {
	result, err := store.NewAuthCode(playerUUID)

	if err != nil {
		t.Error("NewAuthCode returned nothing, because ", err)
	}
	authCode = result
}

func testGetAuthCode(t *testing.T) {
	result, err := store.GetAuthCode(playerUUID)

	if err != nil {
		t.Error("GetAuthCode failed because, ", err)
	}

	if result != authCode {
		t.Errorf("GetAuthCode failed because \"%s\" != \"%s\"\n", authCode, result)
	}
}

func testGetPlayerID(t *testing.T) {
	result, err := store.GetPlayerID(authCode)

	if err != nil {
		t.Error("GetPlayerID failed because, ", err)
	}

	if playerUUID != result {
		t.Errorf("GetPlayerID failed because \"%s\" != \"%s\"\n", playerUUID, result)
	}
}

func testAuthorize(t *testing.T) {
	result, _ := store.Authorize(authCode)

	if playerUUID != result {
		t.Errorf("Authorize failed because \"%s\" != \"%s\"\n", playerUUID, result)
	}
}
