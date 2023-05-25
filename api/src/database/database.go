package database

import (
	"api/src/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var err error
	var database *gorm.DB

	database, err = gorm.Open(sqlite.Open(config.DBConnection), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil, err
	}

	return database, nil
}
