package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(filePath string) (*Database, error) {
	var database = &Database{}

	// check if we need to initialize a new db
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		defer database.Initialize()
	} else {
		return nil, err
	}

	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}

	// Pragma Settings
	db.Exec(`PRAGMA foreign_keys = ON;`)
	db.Exec(`PRAGMA journal_mode = WAL;`)
	db.Exec(`PRAGMA synchronous = NORMAL;`)

	database.db = db

	return database, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
