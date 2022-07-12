package main

import (
	"log"
	"robot-monitoreo/databases"
	"robot-monitoreo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	//Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	//app.Use(requestid.New())
	//Middleware

	//connect database
	databases.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
