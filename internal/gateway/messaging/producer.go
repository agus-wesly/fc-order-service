package messaging

import (
	"encoding/json"
	"log"
	"order-service/internal/model"
	"order-service/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer[T model.Event] interface {
	Send(pattern string, event T) error
}

type ProducerStruct[T model.Event] struct {
	Producer *rabbitmq.RabbitMQProducer
}

type Body[T model.Event] struct {
	Pattern string `json:"pattern"`
	Data    T      `json:"data"`
}

func (p *ProducerStruct[T]) Send(pattern string, event T) error {
	value, err := json.Marshal(p.transformEventToBody(pattern, event))
	if err != nil {
		log.Println("failed to marshal event")
		return err
	}

	if err := p.Producer.Channel.Publish(
		rabbitmq.APP_EXCHANGE_NAME,
		rabbitmq.APP_QUEUE_NAME,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Type:        pattern,
			Body:        value,
		},
	); err != nil {
		return err
	}

	log.Printf("Message sent to topic %s", pattern)
	return nil
}

// This function is needed because nest.js expects incoming message bodies
// to have pattern field
// See : https://stackoverflow.com/questions/56097750/nestjs-there-is-no-matching-event-handler-defined-in-the-remote-service
func (p *ProducerStruct[T]) transformEventToBody(pattern string, event T) Body[T] {
	return Body[T]{
		Pattern: pattern,
		Data:    event,
	}
}
