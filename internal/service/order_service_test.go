package service_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order-service/internal/gateway/messaging"
	"order-service/internal/mocks"
	"order-service/internal/model"
	"order-service/internal/service"
)

/*
mockgen -source=internal/gateway/messaging/producer.go -destination=internal/mocks/mock_producer.go -package=mocks
*/

func TestOrderService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	validate := validator.New()

	mock.ExpectBegin()
	mock.ExpectCommit()

	mockRepo := mocks.NewMockOrderRepository(ctrl)
	mockGateway := mocks.NewMockProductGateway(ctrl)
	mockProducer := mocks.NewMockProducer[*model.OrderEvent](ctrl)
	mockOrderProducer := &messaging.OrderProducer{
		Producer: mockProducer,
	}

	productId := "f3b2e38a-b709-4e03-bd4d-4d0aa54619a9"

	request := &model.CreateOrderRequest{
		ProductId: productId,
	}

	mockGateway.EXPECT().
		GetProductInfo(productId).
		Return(&model.ProductResponse{
			Id:    productId,
			Price: 5000,
			Qty:   10,
		}, nil)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	mockProducer.EXPECT().
		Send("order.created", gomock.Any()).
		Times(1)

	orderService := service.NewOrderService(gormDB, validate, mockRepo, mockOrderProducer, mockGateway)

	resp, err := orderService.Create(context.Background(), request)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}
}
