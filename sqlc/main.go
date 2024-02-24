package main

import (
	"fiber-sqlc/api/route"
	"fiber-sqlc/database"
	"log"

	"github.com/gofiber/fiber/v2"
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
