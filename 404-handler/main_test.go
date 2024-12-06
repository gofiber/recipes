package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Setup function to initialize the Fiber app
func setupApp() *fiber.App {
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return app
}

func TestHelloRoute(t *testing.T) {
	// Initialize the app
	app := setupApp()

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	resp, _ := app.Test(req, -1) // -1 disables timeout

	// Check the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	defer resp.Body.Close()

	// Assert the response body
	assert.Equal(t, "I made a ☕ for you!", string(body))
}

func TestNotFoundRoute(t *testing.T) {
	// Initialize the app
	app := setupApp()

	// Create a test request for an unknown route
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	resp, _ := app.Test(req, -1) // -1 disables timeout

	// Check the response
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
