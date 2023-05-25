package repositories

import (
	"api/src/models"

	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (repository UsersRepository) Create(user models.User) (models.User, error) {
	result := repository.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
