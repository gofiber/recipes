package main

import "os"

// Configuration is used to store values from environment variables
type Configuration struct {
	Port        string
	DatabaseURL string
}

// NewConfiguration reads environment variables and returns a new Configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		Port:        getEnvOrDefault("PORT", "3000"),
		DatabaseURL: getEnvOrDefault("DATABASE_URL", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
