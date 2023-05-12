package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUsers(c echo.Context) error {
	return c.String(http.StatusCreated, "User created")
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
