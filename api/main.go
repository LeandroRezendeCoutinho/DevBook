package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"api/src/router/routes"
	"fmt"
)

func main() {
	config.Load()
	initializeDatabase()

	router := router.Generate()
	routes.DrawUsers(router)
	router.Start(fmt.Sprintf(":%d", config.Port))
}

func initializeDatabase() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	database.Migrate(db)
}
