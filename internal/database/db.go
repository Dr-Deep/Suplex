package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func Open(databaseFilePath string) (*Database, error) {
	db, err := sql.Open("sqlite3", "FILE")
	if err != nil {
		return nil, err
	}

	var database = &Database{
		db: db,
	}

	return database, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
