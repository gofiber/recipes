// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Custom Timer middleware
	app.Use(Timer())

	// Routes
	app.Get("/", func(c *fiber.Ctx) {
		time.Sleep(2 * time.Second) // Sleep 2 seconds
		c.Send("That took a while 😞")
	})

	// Start server
	log.Fatal(app.Listen(3000))
}

// Timer will measure how long it takes before a response is returned
func Timer() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		// start timer
		start := time.Now()
		// next routes
		c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
	}
}
