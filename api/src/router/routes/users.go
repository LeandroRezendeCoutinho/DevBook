package routes

import (
	"api/src/controllers"
	"api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func DrawUsers(r *echo.Echo) {
	r.GET("/users", controllers.FindAllUsers, middlewares.Authenticate)
	r.POST("/users", controllers.CreateUsers)
	r.GET("/users/:id", controllers.FindUserByID, middlewares.Authenticate)
	r.PUT("/users/:id", controllers.UpdateUser, middlewares.Authenticate)
	r.DELETE("/users/:id", controllers.DeleteUser, middlewares.Authenticate)

	r.GET("/users/name", controllers.FindAllUsersByName, middlewares.Authenticate)
	r.GET("/users/:id/followers", controllers.FindFollowers, middlewares.Authenticate)
	r.POST("/users/:id/followers", controllers.CreateFollower, middlewares.Authenticate)
}
