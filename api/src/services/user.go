package services

import (
	"api/src/entities"
	"api/src/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository}
}

func (service UserService) Create(user *entities.User) (*entities.User, error) {
	return service.repository.Create(user)
}
