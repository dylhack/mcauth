package db

import (
	"github.com/dhghf/mcauth/internal/common/db"
	"time"
)

var TestConfig = db.Config{
	Host:               "localhost",
	Port:               5432,
	User:               "",
	Password:           "",
	Database:           "mcauth",
	MaxConnections:     50,
	MaxIdleConnections: 50,
	ConnLifespan:       1 * time.Hour,
}
