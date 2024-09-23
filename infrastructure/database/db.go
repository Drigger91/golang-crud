package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Drigger91/golang-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectToDatabase initializes the connection to the target database
func ConnectToDatabase() {
	// Get the database connection details from environment variables
	if db != nil {
		fmt.Println("Database already connected!")
		return
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	timezone := os.Getenv("DB_TIMEZONE")

	fmt.Println("DEBUG DB CONFIG: " + password + " " + host + " " + user)

	// Now connect to the target database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, user, password, dbname, port, timezone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to the database.")
}

// getDatabase returns the instance of the connected database
func GetDatabase() *gorm.DB {
	if db == nil {
		log.Fatal("Database is not connected")
	}
	return db
}
func AutoMigrate() {
	err := db.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
}
