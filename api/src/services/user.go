package services

import (
	"api/src/models"
	"api/src/repositories"
)

type usersService struct {
	repository *repositories.UsersRepository
}

func NewUserService(repository *repositories.UsersRepository) *usersService {
	return &usersService{repository}
}

func (service usersService) Create(user models.User) (models.User, error) {
	return service.repository.Create(user)
}
