// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Use(Timer())

	app.Get("/", func(c *fiber.Ctx) {
		time.Sleep(2 * time.Second)
		c.Send("That took a while 😞")
	})

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
		// follows server-timing spec
		// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Server-Timing
		c.Append("Server-Timing",
			fmt.Sprintf("app;dur=%v", stop.Sub(start).String()),
		)
	}
}
