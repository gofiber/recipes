package config

import (
	"os"
)

// Config returns the value of an environment variable by key.
// Call godotenv.Load in main() before using this function.
func Config(key string) string {
	return os.Getenv(key)
}
