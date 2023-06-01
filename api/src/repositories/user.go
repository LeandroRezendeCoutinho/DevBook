package repositories

import (
	"api/src/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository UserRepository) Create(user *entities.User) (*entities.User, error) {
	result := repository.db.Create(&user)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}

	return user, nil
}

func (repository UserRepository) Update(user *entities.User) (*entities.User, error) {
	result := repository.db.Save(&user)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}

	return user, nil
}

func (repository UserRepository) Delete(user *entities.User) error {
	result := repository.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository UserRepository) FindAll() (*[]entities.User, error) {
	var users []entities.User

	result := repository.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (repository UserRepository) FindByID(id uint64) (*entities.User, error) {
	var user entities.User

	result := repository.db.First(&user, id)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}

	return &user, nil
}
