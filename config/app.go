package config

import (
	"order-service/internal/delivery/http"
	"order-service/internal/delivery/http/route"
	"order-service/internal/repository"
	"order-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const APP_PORT = 5959

type BootstrapConfig struct {
	DB  *gorm.DB
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(config.DB, orderRepository)
	orderController := http.NewOrderController(orderService)

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
