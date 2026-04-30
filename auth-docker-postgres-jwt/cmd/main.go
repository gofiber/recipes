package main

import (
	"fmt"
	"log"

	"app/database"
	"app/router"

	"github.com/gofiber/fiber/v3"
	// "github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env once at startup; ignore error in production where env vars are set externally
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: no .env file found, using environment variables")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	// app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000", fiber.ListenConfig{EnablePrefork: true}))
}
