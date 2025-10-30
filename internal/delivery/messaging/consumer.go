package messaging

import (
	"context"
	"encoding/json"
	"log"
	"order-service/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerHandler func(data any) error

func ConsumeTopic(ctx context.Context, messages <-chan amqp.Delivery, topic string, handler ConsumerHandler) {
	log.Println("Waiting for messages")

	go func() {
		for message := range messages {
			log.Printf(" > Received message: %s\n", message.Body)
			if message.Type == topic {
				var payload = new(model.OrderEventPayload)
				if err := json.Unmarshal(message.Body, payload); err != nil {
					log.Println(err)
				}

				if err := handler(payload.Data); err != nil {
					log.Println(err)
				}

			}
		}
	}()

	<-ctx.Done()
}
