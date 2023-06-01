package factories

import (
	"api/src/database"

	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}
