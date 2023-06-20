package main

import (
	"fiber-oauth-google/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	router.Routes(app)
	app.Listen(":3300")

}
