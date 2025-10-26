package service

import (
	"order-service/internal/entity"
	"order-service/internal/model"
	"order-service/internal/repository"

	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		// Repository
	}
}

type Order struct {
}

func (c *OrderService) Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error) {
	order := &entity.Order{
		Id:         uuid.New().String(),
		ProductId:  request.ProductId,
		TotalPrice: request.TotalPrice,
		Status:     request.Status,
	}

	if err := c.orderRepository.Create(order); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}
