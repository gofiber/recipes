// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth"
	"google.golang.org/api/option"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Fiber instance
	app := fiber.New()

	// Get google service account credentials
	serviceAccount, fileExi := os.LookupEnv("GOOGLE_SERVICE_ACCOUNT")

	if !fileExi {
		log.Fatalf("Please provide valid firebbase auth credential json!")
	}

	// Initialize the firebase app.
	opt := option.WithCredentialsFile(serviceAccount)
	fireApp, _ := firebase.NewApp(context.Background(), nil, opt)

	// Unauthenticated routes
	app.Get("/salaanthe", salanthe)

	// Initialize the middleware with config. See https://github.com/sacsand/gofiber-firebaseauth for more configuration options.
	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		// Firebase Authentication App Object
		// Mandatory
		FirebaseApp: fireApp,
		// Ignore urls array.
		// Optional. These url will ignore by middleware
		IgnoreUrls: []string{"GET::/salut", "POST::/ciao"},
	}))

	// Authenticaed Routes.
	app.Get("/hello", hello)
	app.Get("/salut", salut)      // Ignore the auth by IgnoreUrls config
	app.Post("/ciao", createCiao) // Ignore the auth by IgnoreUrls config
	app.Get("/ayubowan", ayubowan)

	// Start server.
	log.Fatal(app.Listen(":3001"))
}

/**
*
* Controllers
*
 */

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func salut(c *fiber.Ctx) error {
	return c.SendString("Salut, World 👋!")
}

func ciao(c *fiber.Ctx) error {
	return c.SendString("Ciao, World 👋! ")
}

func createCiao(c *fiber.Ctx) error {
	return c.SendString("Create Cuiao, World 👋! i am a post")
}

func ayubowan(c *fiber.Ctx) error {
	// Get authenticated user from context.
	claims := c.Locals("user")
	fmt.Println(claims)
	return c.SendString("Ayubowan👋! ")
}

func salanthe(c *fiber.Ctx) error {
	return c.SendString("Salanthe👋! ")
}
