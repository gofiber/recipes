package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"sveltekit-embed/handler"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed public/*.*
//go:embed public/*/*.*
//go:embed public/*/*/*.*
//go:embed public/*/*/*/*.*
var svelteFS embed.FS

func main() {

	// Sub
	publicFS, err := fs.Sub(svelteFS, "public")
	if err != nil {
		log.Fatal(err)
	}

	// Create new fiber instance
	app := fiber.New()

	// Main GEO handler that is cached for 10 minutes
	// original code at https://github.com/gofiber/recipes/blob/master/geoip
	app.Get("/:ip.json", handler.Cache(10*time.Minute), handler.GEO())

	// Serve Single Page application
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(publicFS),
		NotFoundFile: "index.html",
	}))

	// Listen on port :8080
	log.Fatal(app.Listen(":8080"))

}
