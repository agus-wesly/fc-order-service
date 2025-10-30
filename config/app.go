package config

import (
	"order-service/internal/delivery/http"
	"order-service/internal/delivery/http/route"
	"order-service/internal/gateway/http"
	"order-service/internal/gateway/messaging"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/pkg/dotenv"
	"order-service/pkg/rabbitmq"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
	Cache    *redis.Client
	Producer *rabbitmq.RabbitMQProducer
}

func Bootstrap(config *BootstrapConfig) {
	orderRepository := repository.NewOrderRepository(config.Cache)

	orderProducer := messaging.NewOrderProducer(config.Producer)

	productGateway := httpgateway.NewProductGateway(dotenv.Getenv("PRODUCT_SERVICE_URL"))

	orderService := service.NewOrderService(config.DB, config.Validate, orderRepository, orderProducer, productGateway)

	orderController := http.NewOrderController(orderService)

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
