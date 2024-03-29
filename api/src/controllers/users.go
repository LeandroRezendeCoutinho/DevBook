package controllers

import (
	"api/src/auth"
	"api/src/entities"
	"api/src/factories"
	"api/src/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateUsers(c echo.Context) error {
	var user entities.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	password, _ := utils.HashPassword(user.Password)
	user.Password = string(password)

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

	tokenUserId, err := auth.ExtractUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	if uint64(tokenUserId) != id {
		return c.JSON(http.StatusForbidden, map[string]string{
			"error": "You can't update a user that is not yours",
		})
	}

	fmt.Println(tokenUserId, id)

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

	tokenUserId, err := auth.ExtractUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	if uint64(tokenUserId) != id {
		return c.JSON(http.StatusForbidden, map[string]string{
			"error": "You can't delete a user that is not yours",
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

func FindAllUsersByName(c echo.Context) error {
	userService, err := factories.NewUserService()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	name := c.QueryParam("name")
	users, err := userService.FindAllBy("name", name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

func FindFollowers(c echo.Context) error {
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

	followers, err := userService.FindFollowers(uint(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, followers)
}

func CreateFollower(c echo.Context) error {
	userService, err := factories.NewUserService()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	followerId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	tokenUserId, err := auth.ExtractUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	if uint64(tokenUserId) == followerId {
		return c.JSON(http.StatusForbidden, map[string]string{
			"error": "You can't follow yourself",
		})
	}

	var follower entities.Follower
	follower.UserID = uint(tokenUserId)
	follower.FollowerID = uint(followerId)
	_, err = userService.CreateFollower(&follower)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}
