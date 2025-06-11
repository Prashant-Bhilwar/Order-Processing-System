package mq

import (
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

var Channel *amqp091.Channel

func InitRabbitMQ() {
	conn, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	_, err = ch.QueueDeclare(
		"order_events",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	Channel = ch
	log.Println("RabbitMQ connected and queue declared")
}
