package common

import (
	"testing"

	"github.com/dylhack/mcauth/internal/common"
)

const (
	playerName = "dhmci"
	playerUUID = "5bce3068e4f3489fb66b5723b2a7cdb1"
)

// TestGetUUID checks if the player name can resolve the linked UUID.
func TestGetUUID(t *testing.T) {
	result := common.GetPlayerID(playerName)

	if playerUUID != result {
		t.Errorf(
			"The UUID doesn't match the result. \"%s\" != \"%s\"\n",
			playerUUID,
			result,
		)
	}
}

// TestGetName checks if the player UUID can resolve the linked name.
func TestGetName(t *testing.T) {
	result := common.GetPlayerName(playerUUID)

	if playerName != result {
		t.Errorf(
			"The name doesn't match the result. \"%s\" != \"%s\"\n",
			playerName,
			result,
		)
	}
}
