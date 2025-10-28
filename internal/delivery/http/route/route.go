package route

import (
	"github.com/gofiber/fiber/v2"
	"order-service/internal/delivery/http"
)

type RouteConfig struct {
	App             *fiber.App
	OrderController *http.OrderController
}

func (c *RouteConfig) Setup() {
	c.App.Post("/orders", c.OrderController.Create)
	c.App.Get("/orders/product/:productId", c.OrderController.GetByProductId)
	c.App.Get("/orders/:id", c.OrderController.GetById)
}
