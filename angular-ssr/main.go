package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/gofiber/fiber/v3/middleware/static"
)

// App creates and configures the Fiber application for serving Angular SSR.
// In Angular SSR (Universal), static browser assets (JS, CSS, images, fonts)
// are served directly by Fiber from ./dist/browser, while dynamic page routes
// are forwarded to the Node.js SSR engine with a bounded timeout.
// If the SSR engine is unavailable or disabled, it falls back to serving the client-side SPA index.html.
func App(nodeSSREndpoint string) *fiber.App {
	app := fiber.New()

	// 1. Native API routes handled directly by Fiber
	api := app.Group("/api")
	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "fiber-angular-ssr",
		})
	})

	// 2. Serve static browser assets (JS chunks, CSS, fonts, favicon, icons)
	// We check if the requested file exists in dist/browser; if not, pass to the next route/SSR handler.
	spaHandler := static.New("./dist/browser/index.html")
	staticAssets := static.New("./dist/browser", static.Config{
		IndexNames: []string{}, // Do not automatically serve index.html for root/dirs so SSR can render pages
	})

	app.Use(func(c fiber.Ctx) error {
		path := c.Path()
		// If requesting a specific file with an extension, attempt static delivery
		if strings.Contains(path, ".") {
			if err := staticAssets(c); err == nil && c.Response().StatusCode() != fiber.StatusNotFound {
				return nil
			}
		}
		return c.Next()
	})

	// 3. Fallback / SSR Proxy: Forward HTML page requests to Angular SSR backend with a timeout
	if nodeSSREndpoint != "" {
		app.Use(func(c fiber.Ctx) error {
			// Bounded proxy timeout of 5 seconds to prevent hung requests
			err := proxy.DoTimeout(c, nodeSSREndpoint+c.OriginalURL(), 5*time.Second)
			if err != nil || c.Response().StatusCode() >= fiber.StatusInternalServerError {
				// Gracefully fallback to client-side SPA index when SSR backend is unreachable or errors
				c.Response().Reset()
				return spaHandler(c)
			}
			return nil
		})
	} else {
		// Pure client-side SPA mode
		app.Use(spaHandler)
	}

	return app
}

// resolveConfig resolves the SSR endpoint and server port from environment variables.
func resolveConfig() (ssrEndpoint string, port string) {
	// Preserve explicitly empty SSR_ENDPOINT to allow running in pure SPA mode
	if val, ok := os.LookupEnv("SSR_ENDPOINT"); ok {
		ssrEndpoint = val
	} else {
		ssrEndpoint = "http://127.0.0.1:4000"
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ssrEndpoint, port
}

// main initializes and starts the Fiber application.
func main() {
	// SSR_ENDPOINT is set by the operator, not the client
	proxy.WithSecurityPolicy(proxy.SecurityPolicy{AllowPrivateIPs: true})

	ssrEndpoint, port := resolveConfig()
	app := App(ssrEndpoint)

	log.Printf("Starting Fiber Angular SSR server on port %s (SSR backend: %s)", port, ssrEndpoint)
	log.Fatal(app.Listen(":" + port))
}
