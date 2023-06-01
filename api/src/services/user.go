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

func (service UserService) Update(user *entities.User) (*entities.User, error) {
	return service.repository.Update(user)
}

func (service UserService) Delete(user *entities.User) error {
	return service.repository.Delete(user)
}

func (service UserService) FindAll() (*[]entities.User, error) {
	return service.repository.FindAll()
}

func (service UserService) FindByID(id uint64) (*entities.User, error) {
	return service.repository.FindByID(id)
}
