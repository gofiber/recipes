package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var envOnce sync.Once

// GetEnv retrieves environment variables
func GetEnv(key string, defaults ...string) string {
	envOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found or error loading: %v", err)
		}
	})
	if len(defaults) > 0 {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaults[0]
	}
	return os.Getenv(key)
}
