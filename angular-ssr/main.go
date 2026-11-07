package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/gofiber/fiber/v3/middleware/static"
)

// App creates and configures the Fiber application for serving Angular SSR.
// In Angular SSR (Universal), dynamic routes are rendered server-side by the Node/Express SSR engine,
// while static browser assets (JS chunks, CSS, images) are served directly by Fiber.
func App(nodeSSREndpoint string) *fiber.App {
	app := fiber.New()

	// 1. API routes handled by Fiber backend
	api := app.Group("/api")
	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "fiber-angular-ssr",
		})
	})

	// 2. Serve static browser assets built by Angular CLI (`dist/<project>/browser`)
	app.Use("/assets", static.New("./dist/browser/assets"))
	app.Use("/favicon.ico", static.New("./dist/browser/favicon.ico"))
	app.Use("/*.js", static.New("./dist/browser"))
	app.Use("/*.css", static.New("./dist/browser"))

	// 3. Fallback / SSR Proxy: Forward HTML page requests to the Angular SSR Node.js server
	if nodeSSREndpoint != "" {
		app.Use(func(c fiber.Ctx) error {
			return proxy.Do(c, nodeSSREndpoint+c.OriginalURL())
		})
	} else {
		// If Node SSR server is not running or disabled, fallback to serving the client-side SPA index.html
		app.Use(static.New("./dist/browser/index.html"))
	}

	return app
}

func main() {
	ssrEndpoint := os.Getenv("SSR_ENDPOINT")
	if ssrEndpoint == "" {
		ssrEndpoint = "http://127.0.0.1:4000"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := App(ssrEndpoint)
	log.Printf("Starting Fiber Angular SSR server on port %s (SSR backend: %s)", port, ssrEndpoint)
	log.Fatal(app.Listen(":" + port))
}
