// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
    "crypto/tls"
    "github.com/gofiber/fiber"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/protocol", func(c *fiber.Ctx) {
		c.Send(c.Protocol()) // => https
	})

	cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	if err != nil {
    	log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	// Start server with https/ssl enabled on http://localhost:443
	app.Listen(443, config)
}
