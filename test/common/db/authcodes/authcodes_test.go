package authcodes

import (
	"github.com/dhghf/mcauth/internal/common/db"
	db2 "github.com/dhghf/mcauth/test/common/db"
	"testing"
)

var playerUUID = "5bce3068e4f3489fb66b5723b2a7cdb1"
var authCode string
var store *db.AuthTable

// TestMain is for getting the database connection before testing.
func TestMain(m *testing.M) {
	if store == nil {
		storeDB := db.GetStore(db2.TestConfig)
		store = &storeDB.Auth
	}
	m.Run()
}

// TestAuthCodes tests the whole auth_codes database table.
func TestAuthCodes(t *testing.T) {
	t.Run("NewAuthCode", testNewAuthCode)
	t.Run("GetAuthCode", testGetAuthCode)
	t.Run("GetPlayerID", testGetPlayerID)
	t.Run("Authorize", testAuthorize)
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
