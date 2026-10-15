package main

import (
	"fmt"
	"log"

	"numtostr/gotodo/app/dal"
	"numtostr/gotodo/app/routes"
	"numtostr/gotodo/config"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	database.Connect()
	if err := database.Migrate(&dal.User{}, &dal.Todo{}); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{ErrorHandler: utils.ErrorHandler})

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.PORT)))
}
