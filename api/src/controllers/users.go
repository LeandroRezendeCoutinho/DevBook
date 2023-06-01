package controllers

import (
	"api/src/entities"
	"api/src/factories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUsers(c echo.Context) error {
	var user entities.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	userService, err := factories.NewUserService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	createdUser, err := userService.Create(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func FindAllUsers(c echo.Context) error {
	return c.String(http.StatusOK, "List of users")
}

func FindUserByID(c echo.Context) error {
	return c.String(http.StatusOK, "User found")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "User updated")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "User deleted")
}
