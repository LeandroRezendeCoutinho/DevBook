package routes

import (
	"api/src/controllers"

	"github.com/labstack/echo/v4"
)

func DrawUsers(r *echo.Echo) {
	r.GET("/users", controllers.FindAllUsers)
	r.POST("/users", controllers.CreateUsers)
	r.GET("/users/:id", controllers.FindUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}
