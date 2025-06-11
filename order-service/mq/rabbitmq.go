package mq

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/prashant-bhilwar/order-processing-system/order-service/model"
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

func StartOrderConsumer(ctx context.Context) {
	msgs, err := Channel.Consume(
		"order_events",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}

	go func() {
		for {
			select {
			case d := <-msgs:
				var order model.Order
				if err := json.Unmarshal(d.Body, &order); err != nil {
					log.Println("Error decoding order event:", err)
					continue
				}
				log.Printf("Order Received from Queue: %+v", order)

			case <-ctx.Done():
				log.Println("Stopping RabbitMQ consumer")
				return
			}
		}
	}()
}
