package main

import (
	"app/template"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

const (
	appName    = "Sveltekit Embed App"
	apiVersion = "v1"
	port       = ":3000"
)

func main() {
	// Create new Fiber Instance
	app := fiber.New(fiber.Config{AppName: appName})
	defer app.Shutdown()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Serve static files
	app.All("/*", static.New("", static.Config{
		FS:           os.DirFS(template.Dist()),
		NotFoundFile: "index.html",
		IndexNames:   []string{"index.html"},
	}))

	// Start the server
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
