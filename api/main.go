package main

import (
	"api/src/config"
	"api/src/controllers"
	"api/src/database"
	"api/src/middlewares"
	"api/src/router"
	"api/src/router/routes"
	"fmt"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Load()
	initializeDatabase()

	router := router.New()

	router.Renderer = middlewares.NewRenderer()
	router.GET("/home", controllers.HomeHandler)

	routes.DrawUsers(router)
	routes.DrawLogin(router)

	router.Use(middleware.Logger())

	router.Start(fmt.Sprintf(":%d", config.Port))
}

func initializeDatabase() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	database.Migrate(db)
}
