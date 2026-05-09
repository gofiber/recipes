// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	ports := []string{":3000", ":3001"}

	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			if err := app.Listen(p); err != nil {
				log.Printf("Error starting server on port %s: %v", p, err)
			}
		}(port)
	}

	wg.Wait()
}
