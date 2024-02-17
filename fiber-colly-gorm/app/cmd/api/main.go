package main

import (
	"log"

	"fiber-colly-gorm/internals/consts"
	"fiber-colly-gorm/internals/services/database"
	"fiber-colly-gorm/internals/services/scrapers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := consts.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables!\n", err.Error())
	}
	database.ConnectDb(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()
	scrape := fiber.New()

	app.Mount("/api", micro)
	app.Mount("/job", scrape)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		go scrapers.StartScraper()
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and Colly",
		})
	})

	scrape.Get("scrape", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Start Scraping...",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
