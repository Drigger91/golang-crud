package main

import (
	"fmt"

	"github.com/Drigger91/golang-crud/infrastructure/database"
	"github.com/Drigger91/golang-crud/infrastructure/env"
	"github.com/gin-gonic/gin"
)

func init() {

	env.LoadEnv()
	database.ConnectToDatabase()
}

func main() {
	fmt.Println("CRUD project -- Golang")

	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "SUCCESS",
			"message": "health check successfully",
		})
	})
	r.Run()
}
