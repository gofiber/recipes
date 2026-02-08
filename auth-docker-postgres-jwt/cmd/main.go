package main

import (
	"log"

	"app/database"
	"app/router"

	"github.com/gofiber/fiber/v3"
	// "github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	// app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000", fiber.ListenConfig{EnablePrefork: true}))
}
