package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const APP_QUEUE_NAME = "app_queue"
const APP_EXCHANGE_NAME = "order_service"

type RabbitMQProducer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}
