package http

import (
	"github.com/gofiber/fiber/v2"
	"order-service/internal/model"
	"order-service/internal/service"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateOrderRequest)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	response, err := c.OrderService.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *OrderController) GetByProductId(ctx *fiber.Ctx) error {
	request := &model.GetOrderByProductIdRequest{
		Id: ctx.Params("productId"),
	}

	response, err := c.OrderService.GetByProductId(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
