package main

import (
	"log"

	"openapi/database"
	"openapi/models"
	"openapi/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Book{})

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
