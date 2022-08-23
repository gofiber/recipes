package config

import (
	"os"
)

// loading of config by dotenv
func Config(key string) string {
	return os.Getenv(key)
}
