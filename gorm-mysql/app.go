package main

import (
	"log"
	"net/http"

	"gorm-mysql/database"
	"gorm-mysql/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(c.SendStatus(http.StatusNotFound)) error {
		return c.SendStatus(http.StatusNotFound) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
