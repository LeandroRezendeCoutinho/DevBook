package repositories

import (
	"api/src/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	GenericRepository *GenericRepository[entities.User]
	db                *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	genericRepository := NewGenericRepository[entities.User](db)
	return &UserRepository{genericRepository, db}
}

func (repository UserRepository) FindFollowers(id uint) (*[]entities.User, error) {
	var followers []entities.User

	err := repository.db.Table("users").
		Select("users.id, users.name, users.email").
		Joins("JOIN followers ON users.id = followers.user_id").
		Where("followers.follower_id = ?", id).
		Scan(&followers).
		Error

	if err != nil {
		return nil, err
	}

	return &followers, nil
}

func (repository UserRepository) CreateFollower(follower *entities.Follower) (*entities.Follower, error) {
	err := repository.db.Create(follower).Error

	if err != nil {
		return nil, err
	}

	return follower, nil
}
