package service

import (
	"order-service/internal/entity"
	"order-service/internal/gateway/http"
	"order-service/internal/model"
	"order-service/internal/model/converter"
	"order-service/internal/repository"

	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderService struct {
	DB              *gorm.DB
	OrderRepository *repository.OrderRepository
	ProductGateway  *httpgateway.ProductGateway
	Validate        *validator.Validate
}

func NewOrderService(db *gorm.DB, validate *validator.Validate, orderRepository *repository.OrderRepository, productGateway *httpgateway.ProductGateway) *OrderService {
	return &OrderService{
		DB:              db,
		OrderRepository: orderRepository,
		ProductGateway:  productGateway,
		Validate:        validate,
	}
}

func (c *OrderService) Create(ctx context.Context, request *model.CreateOrderRequest) (*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		log.Println("error validating request body")
		return nil, fiber.ErrBadRequest
	}

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

	// TODO : publish

	if err := c.OrderRepository.Create(tx, order); err != nil {
		log.Println("error creating order")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("error creating order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}

func (c *OrderService) GetByProductId(ctx context.Context, request *model.GetOrderByProductIdRequest) (*model.GetOrdersByProductIdResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	orders := new([]entity.Order)
	if err := c.OrderRepository.FindByProductId(tx, ctx, orders, request.Id); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("error getting order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrdersToResponse(orders), nil
}
