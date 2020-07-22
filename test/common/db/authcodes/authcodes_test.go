package authcodes

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"os"
	"testing"
)

var playerUUID = "5bce3068e4f3489fb66b5723b2a7cdb1"
var authCode string
var store *db.AuthTable

func TestMain(m *testing.M) {
	if store == nil {
		storeDB := db.GetStore("./mcauth-test.db")
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
	result := store.NewAuthCode(playerUUID)

	if len(result) == 0 {
		t.Error("NewAuthCode returned nothing")
	}
	authCode = result
}

func testGetAuthCode(t *testing.T) {
	result := store.GetAuthCode(playerUUID)

	if result != authCode {
		t.Errorf("Failed GetAuthCode because \"%s\" != \"%s\"\n", authCode, result)
	}
}

func testGetPlayerID(t *testing.T) {
	result := store.GetPlayerID(authCode)

	if playerUUID != result {
		t.Errorf("Failed GetPlayerID because \"%s\" != \"%s\"\n", playerUUID, result)
	}
}

func testAuthorize(t *testing.T) {
	result, _ := store.Authorize(authCode)

	if playerUUID != result {
		t.Errorf("Failed Authorize because \"%s\" != \"%s\"\n", playerUUID, result)
	}
}
