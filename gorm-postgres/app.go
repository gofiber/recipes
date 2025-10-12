package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zeimedee/go-postgres/database"
	"github.com/zeimedee/go-postgres/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Post("/addbook", routes.AddBook)
	app.Post("/book", routes.Book)
	app.Put("/update", routes.Update)
	app.Delete("/delete", routes.Delete)
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
