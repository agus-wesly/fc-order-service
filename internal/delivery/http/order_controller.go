package http

import (
	"github.com/gofiber/fiber/v2"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

type Order struct {
	ProductId  string `json:"product_id"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	order := new(Order)
	if err := ctx.BodyParser(order); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(order)
}
