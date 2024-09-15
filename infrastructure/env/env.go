package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Loading env file variables")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
	}

}
