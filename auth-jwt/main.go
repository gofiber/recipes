package main

import (
	"fmt"
	"log"

	"auth-jwt-gorm/database"
	"auth-jwt-gorm/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env once at startup; ignore error in production where env vars are set externally
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: no .env file found, using environment variables")
	}

	app := fiber.New()
	// restrict in production
	app.Use(cors.New(cors.Config{AllowOrigins: []string{"http://localhost:3000"}}))

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
