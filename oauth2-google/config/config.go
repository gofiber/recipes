package config

import (
	"os"
)

// Config exported via godotenv
func Config(key string) string {
	return os.Getenv(key)
}
