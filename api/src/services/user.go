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
	return service.repository.GenericRepository.Create(*user)
}

func (service UserService) Update(user *entities.User) (*entities.User, error) {
	return service.repository.GenericRepository.Update(*user)
}

func (service UserService) Delete(user *entities.User) error {
	return service.repository.GenericRepository.Delete(*user)
}

func (service UserService) FindAll() (*[]entities.User, error) {
	return service.repository.GenericRepository.FindAll()
}

func (service UserService) FindByID(id uint) (*entities.User, error) {
	return service.repository.GenericRepository.FindByID(id)
}
