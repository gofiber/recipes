package main

import (
	"log"
	"os"

	app "example.com/GofiberFirebaseBoilerplate"
)

func main() {
	port := "3001"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := app.Start(port); err != nil {
		log.Fatalf("app.Start: %v\n", err)
	}
}
