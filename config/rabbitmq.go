package config

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"order-service/pkg/dotenv"
	"order-service/pkg/rabbitmq"
)

func NewRabbitMqProducer() *rabbitmq.RabbitMQProducer {
	connection, err := amqp.Dial(
		// urls: [`amqp://${process.env.RABBITMQ_HOST}:${process.env.RABBITMQ_PORT}`],
		fmt.Sprintf("amqp://%s:%s", dotenv.Getenv("RABBITMQ_HOST"), dotenv.Getenv("RABBITMQ_PORT")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = channel.QueueDeclare(
		rabbitmq.APP_QUEUE_NAME,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return &rabbitmq.RabbitMQProducer{
		Connection: connection,
		Channel:    channel,
	}
}
