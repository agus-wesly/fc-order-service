package http

import (
	"order-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct{
	orderService *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{}
}


func (c *OrderController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateOrderRequest)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}
}
