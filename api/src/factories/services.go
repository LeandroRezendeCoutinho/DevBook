package factories

import (
	"api/src/repositories"
	"api/src/services"
)

func NewUserService() (*services.UserService, error) {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil, err
	}
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	return userService, nil
}
