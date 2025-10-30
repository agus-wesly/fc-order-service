package messaging

import (
	"errors"
	"log"
	"order-service/internal/model"
)

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (c *OrderHandler) OrderCreated(data any) error {
	orderEvent, ok := data.(model.OrderEvent)
	if !ok {
		return errors.New("Unexpected")
	}

	log.Println("Received order.created event with data ", orderEvent)

	return nil
}
