package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	// Read RabbitMQ URL from environment with fallback.
	rabbitmqURL := getEnv("RABBITMQ_URL", "amqp://user:password@localhost:5672/")

	// Create a new RabbitMQ connection.
	connRabbitMQ, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer connRabbitMQ.Close()

	// Create a new Fiber instance.
	app := fiber.New()

	// Add middleware.
	app.Use(
		logger.New(), // add simple logger
	)

	// Add route.
	app.Get("/send", func(c fiber.Ctx) error {
		// Checking, if query is empty.
		if c.Query("msg") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "msg parameter required",
			})
		}

		// Let's start by opening a channel to our RabbitMQ instance
		// over the connection we have already established
		ch, err := connRabbitMQ.Channel()
		if err != nil {
			return err
		}
		defer ch.Close()

		// With this channel open, we can then start to interact.
		// With the instance and declare Queues that we can publish and subscribe to.
		_, err = ch.QueueDeclare(
			"TestQueue",
			true,
			false,
			false,
			false,
			nil,
		)
		// Handle any errors if we were unable to create the queue.
		if err != nil {
			return err
		}

		// Attempt to publish a message to the queue.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err = ch.PublishWithContext(
			ctx,
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(c.Query("msg")),
			},
		)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"status": "message sent"})
	})

	// Graceful shutdown: listen for OS signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	// Start Fiber API server.
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
