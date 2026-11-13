package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/gofiber/fiber/v3/middleware/static"
)

// browserDist is where `ng build` writes the client bundle (see web/angular.json).
const browserDist = "./web/dist/web/browser"

// App creates and configures the Fiber application for serving Angular SSR.
// In Angular SSR, static browser assets (JS, CSS, images, fonts) are served
// directly by Fiber from browserDir, while page routes are forwarded to the
// Node.js SSR engine with a bounded timeout. If the SSR engine is unavailable
// or disabled, Angular's client-side rendering shell (index.csr.html) is served.
func App(nodeSSREndpoint, browserDir string) *fiber.App {
	app := fiber.New()

	// 1. Native API routes handled directly by Fiber
	api := app.Group("/api")
	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "fiber-angular-ssr",
		})
	})

	// 2. Serve static browser assets (JS chunks, CSS, fonts, favicon, icons).
	// Only paths with a file extension are assets, so pages (including "/") reach SSR;
	// a missing file falls through to the next handler.
	app.Use(static.New(browserDir, static.Config{
		Next: func(c fiber.Ctx) bool { return !strings.Contains(c.Path(), ".") },
	}))

	spaHandler := static.New(filepath.Join(browserDir, "index.csr.html"))

	// 3. Fallback / SSR Proxy: Forward HTML page requests to Angular SSR backend with a timeout
	if nodeSSREndpoint != "" {
		app.Use(func(c fiber.Ctx) error {
			// Bounded proxy timeout of 5 seconds to prevent hung requests
			err := proxy.DoTimeout(c, nodeSSREndpoint+c.OriginalURL(), 5*time.Second)
			if err != nil || c.Response().StatusCode() >= fiber.StatusInternalServerError {
				// Gracefully fallback to client-side rendering when SSR backend is unreachable or errors
				c.Response().Reset()
				return spaHandler(c)
			}
			return nil
		})
	} else {
		// Pure client-side rendering mode
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
	app := App(ssrEndpoint, browserDist)

	log.Printf("Starting Fiber Angular SSR server on port %s (SSR backend: %s)", port, ssrEndpoint)
	log.Fatal(app.Listen(":" + port))
}
