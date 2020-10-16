package db

import (
	"testing"

	"github.com/dylhack/mcauth/internal/common/db"
)

// TestGetStore tests connecting to the database.
func TestGetStore(t *testing.T) {
	db.GetStore(TestConfig)
}
