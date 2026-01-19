package main

import (
	"log"

	"fiber-oauth-google/config"
	"fiber-oauth-google/router"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := config.Config("APP_PORT")
	if port == "" {
		port = "3300"
	}

	app := fiber.New()
	router.Routes(app)
	log.Fatal(app.Listen(":" + port))
}
