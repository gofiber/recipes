package config

import (
	"os"

	"github.com/joho/godotenv"
)

func DatabaseEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}