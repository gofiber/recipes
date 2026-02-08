package main

import (
	"log"

	"fiber-colly-gorm/internals/consts"
	"fiber-colly-gorm/internals/services/database"
	"fiber-colly-gorm/internals/services/scrapers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	config, err := consts.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables!\n", err.Error())
	}
	database.ConnectDb(&config)

	app := fiber.New()
	micro := fiber.New()
	scrape := fiber.New()

	app.Use("/api", micro)
	app.Use("/scrape", scrape)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: true,
	}))

	micro.Get("/healthchecker", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and Colly",
		})
	})

	scrape.Get("quotes", func(c fiber.Ctx) error {
		go scrapers.Quotes()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Start scraping quotes.toscrape.com ...",
		})
	})

	scrape.Get("coursera", func(c fiber.Ctx) error {
		go scrapers.CourseraCourses()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Start scraping courses details from coursera.org...",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
