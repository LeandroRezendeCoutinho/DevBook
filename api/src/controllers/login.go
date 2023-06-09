package controllers

import (
	"api/src/auth"
	"api/src/dtos"
	"api/src/factories"
	"api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var userLogin dtos.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	userService, _ := factories.NewUserService()
	user, err := userService.FindOneBy("email", userLogin.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	err = utils.VerifyPassword(userLogin.Password, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, token)
}
