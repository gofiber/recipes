package api

import (
	"fiber-docker-nginx/database"
	"fiber-docker-nginx/middlewares"
	"fiber-docker-nginx/routes"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

/*Init : set the port,cors,api and then serve the api*/
func Init() {
	app := fiber.New()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
	}
	app.Use(cors.New(config))
	app.Use(middlewares.CheckDB)

	if database.Config.ENV == "dev" {
		app.Use(middlewares.RequestLogger)
	}

	/*User : Login and Register routes*/
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	/*Protected Routes*/
	app.Use(middlewares.CheckToken)

	/*Profile : Get and Edit*/
	app.Get("/profile", routes.GetProfile)
	app.Put("/profile", routes.UpdateProfile)

	err := app.Listen(5000)
	if err != nil {
		log.Fatal(err.Error())
	}
}
