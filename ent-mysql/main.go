package main

import (
	"context"
	"log"
	"os"

	"ent-mysql/database"
	"ent-mysql/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Use(cors.New())
	setSchema()
	setRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func setSchema() {
	if err := database.DBConn.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func setRoutes(app *fiber.App) {
	app.Post("/create", routes.CreateBook)
	app.Get("/book/:id", routes.GetBook)
	app.Get("/book", routes.GetAllBook)
	app.Put("/update/:id", routes.UpdateBook)
	app.Delete("/delete/:id", routes.DeleteBook)
}
