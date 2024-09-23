package controllers

import (
	"net/http"

	"github.com/Drigger91/golang-crud/infrastructure/database"
	"github.com/Drigger91/golang-crud/models"
	"github.com/gin-gonic/gin"
)

// GetTransactions fetches all transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	db := database.GetDatabase()
	db.Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

// CreateTransaction creates a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.GetDatabase()
	db.Create(&transaction)
	c.JSON(http.StatusOK, transaction)
}

// GetTransaction fetches a single transaction by ID
func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	db := database.GetDatabase()
	if err := db.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction updates a transaction
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	db := database.GetDatabase()
	if err := db.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&transaction)
	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction deletes a transaction
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	db := database.GetDatabase()
	if err := db.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	db.Delete(&transaction)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
