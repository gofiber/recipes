package config

import (
	"fmt"
	"os"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT", "8000")
	// DB returns the connection string of the database.
	DB = getEnv("DB", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "<token-key>")
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

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
