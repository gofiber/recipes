// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"

	"main/handler"
	"main/middleware"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Fiber instance
	app := fiber.New()

	// Get google service account credentials
	serviceAccount, exists := os.LookupEnv("GOOGLE_SERVICE_ACCOUNT")
	if !exists {
		log.Fatalf("Please provide valid firebase auth credential json!")
	}

	// Initialize the firebase app.
	opt := option.WithCredentialsFile(serviceAccount)
	fireApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("failed to initialize firebase app: %v", err)
	}

	authClient, err := fireApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize firebase auth client: %v", err)
	}

	// Public routes — no authentication required.
	app.Get("/salut", handler.Salut)
	app.Post("/ciao", handler.Ciao)
	app.Get("/salanthe", handler.Salanthe)

	// Protected routes — Firebase auth middleware applied to the whole group.
	api := app.Group("/api", middleware.FirebaseAuth(authClient))
	api.Get("/hello", handler.Hello)
	api.Get("/ayubowan", handler.Ayubowan)

	// Start server.
	log.Fatal(app.Listen(":3001"))
}
