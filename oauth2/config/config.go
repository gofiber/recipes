package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config function to get value from env file
// perhaps not the best implementation
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
