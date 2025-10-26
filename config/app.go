package config

import (
	"order-service/internal/delivery/http"
	"order-service/internal/delivery/http/route"

	"github.com/gofiber/fiber/v2"
)

const APP_PORT = 5959

type BootstrapConfig struct {
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	orderController := http.NewOrderController()

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
