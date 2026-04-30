package main

import (
	"catalog/api"
	"catalog/config"
	"catalog/repository"
	"catalog/service"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	conf, err := config.NewConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewProductService(repo)
	handler := api.NewHandler(service)

	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	r.Get("/products/{code}", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/{code}", handler.Delete)
	r.Get("/products", handler.GetAll)
	r.Put("/products", handler.Put)
	r.Listen(":8080")
}
