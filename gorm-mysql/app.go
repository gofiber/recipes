package main

import (
	"log"

	"gorm-mysql/database"
	"gorm-mysql/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Get("/book/:id", routes.GetBook)
	app.Post("/book", routes.AddBook)
	app.Put("/book/:id", routes.Update)
	app.Delete("/book/:id", routes.Delete)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
