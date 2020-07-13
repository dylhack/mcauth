package db

import (
	"database/sql"
	util "github.com/Floor-Gang/utilpkg/db"
	"log"
)

type Store struct {
	db *sql.DB
	Alts AltsTable
	Auth AuthTable
	Links LinksTable
}

func GetStore(dbPath string) (c Store) {
	db, err := util.GetSQLiteDB(dbPath)

	if err != nil {
		log.Fatalln("Failed to SQLite database\n" + err.Error())
	}

	c.db = db
	c.Alts = GetAltsTable(db)
	c.Auth = GetAuthTable(db)
	c.Links = GetLinksTable(db)

	return c
}
