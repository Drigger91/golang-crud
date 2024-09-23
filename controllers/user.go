package controllers

import (
	"net/http"

	"github.com/Drigger91/golang-crud/infrastructure/database"
	"github.com/Drigger91/golang-crud/models"
	"github.com/gin-gonic/gin"
)

// GetUsers fetches all users
func GetUsers(c *gin.Context) {
	var users []models.User
	db := database.GetDatabase()
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.GetDatabase()
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// GetUser fetches a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := database.GetDatabase()
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user's information
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := database.GetDatabase()
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := database.GetDatabase()
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
