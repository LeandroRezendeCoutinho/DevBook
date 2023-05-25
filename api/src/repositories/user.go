package repositories

import (
	"api/src/models"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{db}
}

func (repository usersRepository) Create(user models.User) (models.User, error) {
	result := repository.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
