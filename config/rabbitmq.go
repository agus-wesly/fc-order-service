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
		fmt.Sprintf("amqp://%s:%s", dotenv.Getenv("RABBITMQ_HOST"), dotenv.Getenv("RABBITMQ_PORT")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	err = channel.ExchangeDeclare(
		rabbitmq.APP_EXCHANGE_NAME, // exchange name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // args
	)
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

func NewRabbitMqConsumer() <-chan amqp.Delivery {
	connection, err := amqp.Dial(
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
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = channel.QueueBind(
		rabbitmq.APP_QUEUE_NAME,
		"",                     
		rabbitmq.APP_EXCHANGE_NAME,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	messages, err := channel.Consume(
		rabbitmq.APP_QUEUE_NAME,
		"",                      
		true,                    
		false,                   
		false,                   
		false,                   
		nil,                     
	)
	if err != nil {
		log.Fatalln(err)
	}

	return messages
}
