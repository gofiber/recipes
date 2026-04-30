// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"os"

	"github.com/amalshaji/stoyle/database"
	"github.com/amalshaji/stoyle/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("warn: .env file not found, relying on environment variables: %v", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
