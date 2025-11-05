package main

import (
    "fmt"
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Welcome to seenode 👋")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
