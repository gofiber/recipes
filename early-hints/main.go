// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := setupApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on http://127.0.0.1:%s", port)
	log.Fatal(app.Listen(":" + port))
}

func setupApp() *fiber.App {
	app := fiber.New()

	// Route that sends HTTP 103 Early Hints before long-running computation
	app.Get("/", func(c fiber.Ctx) error {
		// Define critical assets for the client browser to preload
		hints := []string{
			"</css/main.css>; rel=preload; as=style",
			"</js/bundle.js>; rel=preload; as=script",
		}

		// Send 103 Early Hints to client
		// If the client/protocol doesn't support Early Hints, Fiber continues gracefully
		if err := c.SendEarlyHints(hints); err != nil {
			log.Printf("SendEarlyHints warning: %v", err)
		}

		// Simulate server-side processing (e.g. database query, API aggregation)
		time.Sleep(50 * time.Millisecond)

		// Return the final HTML response
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Fiber Early Hints Example</title>
    <link rel="stylesheet" href="/css/main.css">
    <script src="/js/bundle.js" defer></script>
</head>
<body>
    <h1>Fiber Early Hints (HTTP 103)</h1>
    <p>Early Hints were dispatched while the server was preparing this page.</p>
</body>
</html>`)
	})

	// Static asset endpoints for demonstration
	app.Get("/css/main.css", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/css")
		return c.SendString("body { font-family: sans-serif; background: #f9f9f9; padding: 2rem; }")
	})

	app.Get("/js/bundle.js", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "application/javascript")
		return c.SendString("console.log('Preloaded bundle loaded successfully');")
	})

	// Healthcheck endpoint
	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}
