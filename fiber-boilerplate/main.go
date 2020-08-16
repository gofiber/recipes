package main

import (
	"flag"
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/libraries"
	"github.com/itsursujit/fiber-boilerplate/middlewares"
	. "github.com/itsursujit/fiber-boilerplate/migrations"
	"github.com/itsursujit/fiber-boilerplate/routes"
)

func main() {
	Log = libraries.SetupZeroLog()
	migrate := flag.Bool("migrate", false, "Migrate the pending migration files")
	flag.Parse()
	if *migrate {
		InitMigrate()
		return
	}
	Serve()
}

func Serve() {
	Boot()
	App.Use(middlewares.LogMiddleware)
	routes.Load()
	App.Use(func(c *fiber.Ctx) {
		var err fiber.Error
		err.Code = fiber.StatusNotFound
		config.CustomErrorHandler(c, &err)
	})
	// go libraries.Consume("webhook-callback")               //nolint:wsl
	err := App.Listen(config.AppConfig.App_Port) //nolint:wsl
	if err != nil {
		panic("App not starting: " + err.Error() + "on Port: " + config.AppConfig.App_Port)
	}

	defer DB.Close()
}

func Boot() {
	config.LoadEnv()
	config.BootApp()
	LoadComponents()
}

func LoadComponents() {
	config.LoadQueueConfig()
	config.LoadPaypalConfig()
	Queue = libraries.SetupQueue() //nolint:wsl
}
