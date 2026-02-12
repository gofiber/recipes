// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
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
	fireApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("failed to initialize firebase app: %v", err)
	}

	authClient, err := fireApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize firebase auth client: %v", err)
	}

	// Unauthenticated routes
	app.Get("/salaanthe", salanthe)

	ignoreRoutes := map[string]struct{}{
		"GET::/salut":     {},
		"POST::/ciao":     {},
		"GET::/salaanthe": {},
	}
	app.Use(func(c fiber.Ctx) error {
		if _, ok := ignoreRoutes[c.Method()+"::"+c.Path()]; ok {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
		}

		idToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if idToken == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing bearer token")
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		c.Locals("user", token.Claims)
		return c.Next()
	})

	// Authenticaed Routes.
	app.Get("/hello", hello)
	app.Get("/salut", salut) // Ignore the auth by IgnoreUrls config
	app.Post("/ciao", ciao)  // Ignore the auth by IgnoreUrls config
	app.Get("/ayubowan", ayubowan)

	// Start server.
	log.Fatal(app.Listen(":3001"))
}

/**
*
* Controllers
*
 */

func hello(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func salut(c fiber.Ctx) error {
	return c.SendString("Salut, World 👋!")
}

func ciao(c fiber.Ctx) error {
	return c.SendString("Ciao, World 👋! ")
}

func ayubowan(c fiber.Ctx) error {
	// Get authenticated user from context.
	claims := c.Locals("user")
	fmt.Println(claims)
	return c.SendString("Ayubowan👋! ")
}

func salanthe(c fiber.Ctx) error {
	return c.SendString("Salanthe👋! ")
}
