package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a new RabbitMQ connection.
	connRabbitMQ, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// Create a new Fiber instance.
	app := fiber.New()

	// Add middleware.
	app.Use(
		logger.New(), // add simple logger
	)

	// Add route.
	app.Get("/send", func(c *fiber.Ctx) error {
		// Checking, if query is empty.
		if c.Query("msg") == "" {
			log.Println("Missing 'msg' query parameter")
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

		return nil
	})

	// Start Fiber API server.
	log.Fatal(app.Listen(":3000"))
}
