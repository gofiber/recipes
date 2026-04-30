package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"clean-architecture/api/routes"
	"clean-architecture/pkg/book"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatalf("Database Connection Error: %s", err)
	}
	fmt.Println("Database connection success!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	defer cancel()
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}
