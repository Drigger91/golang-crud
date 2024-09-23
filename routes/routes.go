package routes

import (
	"github.com/Drigger91/golang-crud/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		// Health check
		api.GET("/health-check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "SUCCESS",
				"message": "health check successfully",
			})
		})
		// User routes
		api.GET("/users", controllers.GetUsers)
		api.POST("/users/create", controllers.CreateUser)
		api.GET("/users/:id", controllers.GetUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

		// Transaction routes
		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
		api.GET("/transactions/:id", controllers.GetTransaction)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("/transactions/:id", controllers.DeleteTransaction)
	}

	return router
}
