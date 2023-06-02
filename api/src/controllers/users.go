package controllers

import (
	"api/src/entities"
	"api/src/factories"
	"net/http"
	"strconv"

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
	userService, err := factories.NewUserService()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	users, err := userService.FindAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

func FindUserByID(c echo.Context) error {
	userService, err := factories.NewUserService()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user, err := userService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	var user entities.User

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	userService, err := factories.NewUserService()
	userFound, findErr := userService.FindByID(uint(id))

	if findErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": findErr.Error(),
		})
	}

	if user.Name != "" {
		userFound.Name = user.Name
	}

	if user.Email != "" {
		userFound.Email = user.Email
	}

	if user.Password != "" {
		userFound.Password = user.Password
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	updatedUser, err := userService.Update(userFound)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
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

	userFound, findErr := userService.FindByID(uint(id))
	if findErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": findErr.Error(),
		})
	}

	err = userService.Delete(userFound)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, id)
}
