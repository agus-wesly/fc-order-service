package messaging

import (
	"order-service/internal/model"
	"order-service/pkg/rabbitmq"
)

type OrderProducer struct {
	Producer[*model.OrderEvent] 
}

func NewOrderProducer(producer *rabbitmq.RabbitMQProducer) *OrderProducer {
	return &OrderProducer{
		Producer: &ProducerStruct[*model.OrderEvent]{
			Producer: producer,
		},
	}
}
