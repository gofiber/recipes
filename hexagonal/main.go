package main

import (
	"log"

	"catalog/api"
	"catalog/config"
	"catalog/repository"
	"catalog/service"

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

	// Fiber route parameters are written ":name". A "{code}" segment would only ever match
	// the literal text "{code}", so these two routes could never be reached.
	r.Get("/products/:code", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/:code", handler.Delete)
	r.Get("/products", handler.GetAll)
	r.Put("/products", handler.Put)

	addr := conf.Server.Port
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(r.Listen(addr))
}
