// For more information about each table visit their file. This is where
// they're all grouped together in the Store struct.
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// Postgres configuration
type Config struct {
	Host               string        `yaml:"host"`
	Port               int           `yaml:"port"`
	User               string        `yaml:"username"`
	Password           string        `yaml:"password"`
	Database           string        `yaml:"database_name"`
	MaxConnections     int           `yaml:"max_connections"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	ConnLifespan       time.Duration `yaml:"conn_lifespan"`
}

type Store struct {
	db    *sql.DB
	Alts  AltsTable
	Auth  AuthTable
	Links LinksTable
}

func GetStore(config Config) (c Store) {
	connConfig := fmt.Sprintf(
		"user=%s password=%s host=%s database=%s port=%d sslmode=disable",
		config.User, config.Password, config.Host, config.Database, config.Port,
	)
	db, err := sql.Open("postgres", connConfig)

	if err != nil {
		log.Fatalln("Failed to connect to the postgres database\n", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Failed to ping the postgres database\n", err.Error())
	}

	c.db = db

	c.db.SetMaxOpenConns(config.MaxConnections)
	c.db.SetMaxIdleConns(config.MaxIdleConnections)
	c.db.SetConnMaxLifetime(config.ConnLifespan)

	// Alt account management table
	c.Alts = GetAltsTable(db)
	// Authentication code table
	c.Auth = GetAuthTable(db)
	// Linked accounts table
	c.Links = GetLinksTable(db)

	return c
}
