package libraries

import (
	"encoding/json"
	"fmt"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/services"
	"github.com/streadway/amqp"
)

func Consume(queue string) {
	conn := Queue

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var webhook services.Webhook
			if err := json.Unmarshal(d.Body, &webhook); err != nil {
				panic(err)
			}
			go webhook.Dispatch()
			Log.Printf("Received a message: %s", d.Body)
		}
	}()

	Log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func Publish(queue string, payload interface{}) {
	conn := Queue

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload.(string)),
		})
	Log.Printf(" [x] Sent %s", payload)
	failOnError(err, "Failed to publish a message")
}

func SetupQueue() *amqp.Connection {
	AMQPConnectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s", config.QueueConfig.Queue_User, config.QueueConfig.Queue_Pass, config.QueueConfig.Queue_Host, config.QueueConfig.Queue_Port)
	conn, err := amqp.Dial(AMQPConnectionURL)
	failOnError(err, "Can't connect to AMQP")
	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(err)
	}
}
