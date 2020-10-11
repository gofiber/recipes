package main

import (
	"log"

	"github.com/dinopuguh/gofiber-recipes/swagger/database"
	_ "github.com/dinopuguh/gofiber-recipes/swagger/docs"
	"github.com/dinopuguh/gofiber-recipes/swagger/models"
	"github.com/dinopuguh/gofiber-recipes/swagger/routes"
)

// @title Book App
// @version 1.0
// @description This is an API for Book Application

// @contact.name Dino Puguh
// @contact.email dinopuguh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Book{})

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
