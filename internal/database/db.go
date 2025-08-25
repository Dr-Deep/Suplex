package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
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

	// Open DB
	db, err := gorm.Open(
		sqlite.Open(filePath),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	database.db = db

	return database, nil
}

func (d *Database) Close() error {
	if d.db == nil {
		return nil
	}

	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
