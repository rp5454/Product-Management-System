package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
	"github.com/product-service/models"
)

var RabbitMQConn *amqp091.Connection
var RabbitMQChannel *amqp091.Channel

func InitRabbitMQ() {
	var err error
	RabbitMQConn, err = amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Failed to create RabbitMQ channel:", err)
	}

	_, err = RabbitMQChannel.QueueDeclare(
		"product_queue", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare RabbitMQ queue:", err)
	}
}

func PublishToRabbitMQ(product models.Product) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Println("Failed to marshal product:", err)
		return
	}

	err = RabbitMQChannel.Publish(
		"", "product_queue", false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        productJSON,
		},
	)
	if err != nil {
		log.Println("Failed to publish message to RabbitMQ:", err)
	}
}
