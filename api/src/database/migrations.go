package database

import (
	"api/src/models"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	return database.AutoMigrate(&models.User{})
}
