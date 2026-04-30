package main

import (
	"log"

	"swagger/database"
	_ "swagger/docs"
	"swagger/models"
	"swagger/routes"
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
		log.Fatal("Can't connect database:", err.Error())
	}

	if err := database.DBConn.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal(err)
	}

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
