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
