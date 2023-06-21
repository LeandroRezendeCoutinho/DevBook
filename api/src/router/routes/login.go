package routes

import (
	"api/src/controllers"

	"github.com/labstack/echo/v4"
)

func DrawLogin(r *echo.Echo) {
	r.POST("/login", controllers.Login)
}
