// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Custom Timer middleware
	app.Use(Timer())

	// Routes
	app.Get("/", func(c fiber.Ctx) error {
		time.Sleep(2 * time.Second) // Sleep 2 seconds
		return c.SendString("That took a while 😞")
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// Do something with response
		// dur value must be in milliseconds per W3C spec
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%.2f", float64(time.Since(start).Microseconds())/1000.0))
		// return stack error if exist
		return err
	}
}
