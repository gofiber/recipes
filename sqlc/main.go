package main

import (
	"log"

	"fiber-sqlc/api/route"
	"fiber-sqlc/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	route.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
