package main

import (
	"log"
	"os"

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
	// Name of Docker container with RabbitMQ: dev-rabbitmq (see docker-compose.yml)
	rabbitmqURL := getEnv("RABBITMQ_URL", "amqp://user:password@dev-rabbitmq:5672/")

	// Create a new RabbitMQ connection.
	connRabbitMQ, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer connRabbitMQ.Close()

	// Open a new channel.
	channel, err := connRabbitMQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	// Start delivering queued messages.
	messages, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Welcome message.
	log.Println("Successfully connected to RabbitMQ instance")
	log.Println("[*] - Waiting for messages")
	log.Println("[*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>")

	// Open a channel to receive messages.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, just show received message in console.
			log.Printf("Received message: %s\n", message.Body)
		}
	}()

	// Block until channel is closed.
	<-forever
}
