package service

import (
	"order-service/internal/entity"
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
}

func NewOrderService(db *gorm.DB, orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{
		DB:              db,
		OrderRepository: orderRepository,
	}
}

func (c *OrderService) Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error) {
	order := &entity.Order{
		Id:         uuid.New().String(),
		ProductId:  request.ProductId,
		TotalPrice: request.TotalPrice,
		Status:     request.Status,
	}

	if err := c.OrderRepository.Create(c.DB, order); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}
