package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creates a new Fiber instance.
	app := fiber.New()

	// Group router to subroutes.
	api := app.Group("/api/v1")

	// Create a small route to check for the health.
	api.Get("", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Hello World!",
		})
	})

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))
}
