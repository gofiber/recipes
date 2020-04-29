package main

import (
	"log"
	"os"

	app "github.com/gofiber/gcloud"
)

func main() {

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := app.Start(port); err != nil {
		log.Fatalf("app.Start: %v\n", err)
	}
}
