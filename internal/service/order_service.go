package service

import (
	"order-service/internal/entity"
	"order-service/internal/gateway/http"
	"order-service/internal/model"
	"order-service/internal/model/converter"
	"order-service/internal/repository"

	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderService struct {
	DB              *gorm.DB
	OrderRepository *repository.OrderRepository
	ProductGateway  *httpgateway.ProductGateway
}

func NewOrderService(db *gorm.DB, orderRepository *repository.OrderRepository, productGateway *httpgateway.ProductGateway) *OrderService {
	return &OrderService{
		DB:              db,
		OrderRepository: orderRepository,
		ProductGateway:  productGateway,
	}
}

func (c *OrderService) Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error) {
	orderResponse, err := c.ProductGateway.GetProductInfo(request.ProductId)
	if err != nil {
		return nil, err
	}

	order := &entity.Order{
		Id:         uuid.New().String(),
		ProductId:  orderResponse.Id,
		TotalPrice: orderResponse.Price,
		Status:     "CREATED",
	}

	if err := c.OrderRepository.Create(c.DB, order); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}
