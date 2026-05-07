package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm-postgres/database"
	"gorm-postgres/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Get("/book/:id", routes.Book)
	app.Post("/book", routes.AddBook)
	app.Put("/book/:id", routes.Update)
	app.Delete("/book/:id", routes.Delete)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())

	setUpRoutes(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
