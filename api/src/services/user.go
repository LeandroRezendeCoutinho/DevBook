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

func (service UserService) FindOneBy(field string, email string) (*entities.User, error) {
	return service.repository.GenericRepository.FindOneBy(field, email)
}

func (service UserService) FindAllBy(field string, name string) (*[]entities.User, error) {
	return service.repository.GenericRepository.FindAllBy(field, name)
}

func (service UserService) FindFollowers(id uint) (*[]entities.User, error) {
	return service.repository.FindFollowers(id)
}

func (service UserService) CreateFollower(follower *entities.Follower) (*entities.Follower, error) {
	return service.repository.CreateFollower(follower)
}
