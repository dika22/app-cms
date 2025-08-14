package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func Connection() *amqp091.Channel {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

	return ch
}

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}