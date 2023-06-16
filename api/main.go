package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"api/src/router/routes"
	"fmt"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Load()
	initializeDatabase()

	router := router.New()
	routes.DrawUsers(router)

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
