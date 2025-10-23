package main

import (
    "fmt"
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
        port = "80"
    }
    
    app.Listen(fmt.Sprintf(":%s", port))
}
