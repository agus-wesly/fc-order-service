package service_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order-service/internal/mocks"
	"order-service/internal/model"
	"order-service/internal/service"
)

/*
mockgen -source=internal/repository/order_repository.go -destination=internal/mocks/mock_order_repository.go -package=mocks
*/

func TestOrderService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqlDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	validate := validator.New()

	mockRepo := mocks.NewMockOrderRepository(ctrl)
	mockGateway := mocks.NewMockProductGatewayInterface(ctrl)
	mockProducer := mocks.NewMockOrderProducerInterface(ctrl)

	// Prepare input
	request := &model.CreateOrderRequest{
		ProductId: "p123",
	}

	// Mock expectations
	mockGateway.EXPECT().
		GetProductInfo("p123").
		Return(&model.ProductResponse{
			Id:    "p123",
			Price: 5000,
			Qty:   10,
		}, nil)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	mockProducer.EXPECT().
		Send("order.created", gomock.Any()).
		Times(1)

	orderService := service.NewOrderService(gormDB, validate, mockRepo, mockProducer, mockGateway)

	// Execute
	resp, err := orderService.Create(context.Background(), request)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}
}
