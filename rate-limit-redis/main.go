package main

import (
	"ratelimitfiber/config"
	"ratelimitfiber/middleware"
	"time"

	"ratelimitfiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	config.RedisInit()
	app.Use(middleware.RateLimiterMiddleware(config.RedisClient, time.Minute, 10))

	app.Get("/home", handlers.Home)

	app.Listen(":8080")
}
