package db

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"testing"
)

// TestGetStore tests connecting to the database.
func TestGetStore(t *testing.T) {
	db.GetStore(TestConfig)
}
