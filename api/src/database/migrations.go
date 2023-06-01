package database

import (
	"api/src/entities"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	return database.AutoMigrate(&entities.User{})
}
