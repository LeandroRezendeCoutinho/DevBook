package controllers

import (
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(200, "home.html", map[string]interface{}{
		"name": "John Doe",
	})
}
