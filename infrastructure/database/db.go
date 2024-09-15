package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// getDatabase returns the instance of the connected database
// createDatabase checks if the database exists and creates it if it doesn't
func createDatabase(dbname, user, password, host string, port int) error {
	// Connect to the default `postgres` database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable", host, user, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to default postgres database: %v", err)
	}

	// Check if the target database exists, and create it if not
	query := fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := db.Exec("SELECT 1 FROM pg_database WHERE datname = ?", dbname).Error; err != nil {
		// If the query fails, create the database
		if err := db.Exec(query).Error; err != nil {
			return fmt.Errorf("failed to create database %s: %v", dbname, err)
		}
		fmt.Printf("Database %s created successfully.\n", dbname)
	} else {
		fmt.Printf("Database %s already exists.\n", dbname)
	}

	return nil
}

// ConnectToDatabase initializes the connection to the target database
func ConnectToDatabase() {
	// Load environment variables from .env

	// Get the database connection details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	timezone := os.Getenv("DB_TIMEZONE")

	// Try to create the database if it doesn't exist
	if err := createDatabase(dbname, user, password, host, port); err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

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
