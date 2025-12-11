package main

import (
	"log"

	"fiber-sqlc/api/route"
	"fiber-sqlc/database"

	"github.com/gofiber/fiber/v3"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	route.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
