package main

// thanks to https://github.com/Learn-by-doing/csrf-examples
import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"main/routes"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html").Layout("views/layouts/main")

	go func() {
		// ### EVIL SERVER ###
		// Fiber instance
		app := fiber.New(fiber.Config{Views: engine})
		routes.RegisterEvilRoutes(app)
		fmt.Println("\"Evil\" server started and listening at localhost:3001")
		// Start server
		log.Fatal(app.Listen(":3001"))
	}()

	// ### NORMAL SERVER ###
	// Fiber instance
	app := fiber.New(fiber.Config{Views: engine})
	routes.RegisterRoutes(app)
	fmt.Printf("Server started and listening at localhost:3000 - csrfActive: %v\n", len(os.Args) > 1 && os.Args[1] == "withoutCsrf")
	// Start server
	log.Fatal(app.Listen(":3000"))
}
