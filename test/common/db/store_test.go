package db

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"testing"
)

func TestGetStore(t *testing.T) {
	db.GetStore(TestConfig)
}
