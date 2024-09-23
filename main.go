package main

import (
	"fmt"

	"github.com/Drigger91/golang-crud/infrastructure/database"
	"github.com/Drigger91/golang-crud/infrastructure/env"
	"github.com/Drigger91/golang-crud/routes"
)

// Runs before main function automatically, can load all the required stuff here
func init() {
	env.LoadEnv()
	database.ConnectToDatabase()
	database.AutoMigrate()
}

func main() {
	fmt.Println("CRUD project -- Golang")

	router := routes.InitializeRoutes()
	for _, route := range router.Routes() {
		fmt.Printf("Registered route: %s - %s\n", route.Method, route.Path)
	}
	router.Run()

}
