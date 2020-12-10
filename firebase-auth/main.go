// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth"
	"google.golang.org/api/option"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Fiber instance
	app := fiber.New()

	// get firauth admin sdk json from .env
	file, fileExi := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")

	if !fileExi {
		log.Println("fireauth config not found")
	}
	// 2) create firebase app
	opt := option.WithCredentialsFile(file)
	fireApp, _ := firebase.NewApp(context.Background(), nil, opt)

	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		// Firebase Authentication App Object
		// Mandatory
		FirebaseApp: fireApp,
		// Ignore urls array
		// Optional
		IgnoreUrls: []string{"GET::/salut", "POST::/testauth "},
	}))

	// Routes
	app.Get("/hello", hello)
	app.Get("/salut", salut)
	app.Post("/ciao", ciao)
	app.Post("/ciao", createCiao)

	// Start server
	log.Fatal(app.Listen(":3001"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func salut(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func ciao(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func createCiao(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
