package config

import (
	"log"
	"os"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT", "5000")
	// DB returns the name of the sqlite database
	DB = getEnv("DB", "gotodo.db")
	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP = getEnv("TOKEN_EXP", "10h")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	log.Fatalf("Environment variable not found :: %s", name)
	return ""
}
