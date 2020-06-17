package main

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(3000)

	defer database.DB.Close()
}
