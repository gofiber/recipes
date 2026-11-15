package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"

	"app/template"
)

const (
	appName    = "Sveltekit Embed App"
	apiVersion = "v1"
	port       = ":3000"
)

func main() {
	// Create new Fiber Instance
	app := fiber.New(fiber.Config{AppName: appName})

	// Middleware
	app.Use(cors.New(cors.Config{
		// Development only - restrict in production
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Serve static files
	dist := template.Dist()
	app.All("/*", static.New("", static.Config{
		FS: dist,
		// A path that is not a file in the build, such as a client-side route,
		// gets index.html so the SvelteKit router can handle it. The static
		// middleware has set the status to 404 by then, so reset it to 200.
		NotFoundHandler: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusOK).SendFile("index.html", fiber.SendFile{FS: dist})
		},
		IndexNames: []string{"index.html"},
	}))

	// Listen for OS signals to trigger graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		if err := app.Shutdown(); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}
	}()

	// Start the server
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
