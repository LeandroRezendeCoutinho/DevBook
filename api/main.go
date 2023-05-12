package main

import (
	"api/src/router"
	"api/src/router/routes"
	"fmt"
)

func main() {
	fmt.Println("Running API")
	router := router.Generate()

	routes.DrawUsers(router)

	router.Start(":3000")
}
