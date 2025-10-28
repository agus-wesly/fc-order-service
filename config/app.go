package config

import (
	"order-service/internal/delivery/http"
	"order-service/internal/delivery/http/route"
	"order-service/internal/gateway/http"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/pkg/dotenv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const APP_PORT = 5959

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
	Cache    *redis.Client
}

func Bootstrap(config *BootstrapConfig) {
	productGateway := httpgateway.NewProductGateway(dotenv.Getenv("PRODUCT_SERVICE_URL"))
	orderRepository := repository.NewOrderRepository(config.Cache)
	orderService := service.NewOrderService(config.DB, config.Validate, orderRepository, productGateway)
	orderController := http.NewOrderController(orderService)

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
