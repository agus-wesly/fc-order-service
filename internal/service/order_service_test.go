package service_test

import (
	"context"
	"order-service/internal/entity"
	"order-service/internal/model"
	"order-service/internal/service"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) WithContext(ctx context.Context) *MockDB {
	return m
}
func (m *MockDB) Begin() *MockDB {
	return m
}
func (m *MockDB) Commit() *MockDB {
	m.Called()
	return m
}
func (m *MockDB) Rollback() {}

type MockProductGateway struct{ mock.Mock }
func (m *MockProductGateway) GetProductInfo(id string) (*model.ProductResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*model.ProductResponse), args.Error(1)
}

type MockOrderRepository struct{ mock.Mock }
func (m *MockOrderRepository) Create(tx any, order *entity.Order) error {
	args := m.Called(tx, order)
	return args.Error(0)
}

type MockOrderProducer struct{ mock.Mock }
func (m *MockOrderProducer) Send(event string, payload *model.OrderEvent) {
	m.Called(event, payload)
}

// --- Fake validator ---
type FakeValidator struct{}
func (v *FakeValidator) Struct(i interface{}) error { return nil }


func TestOrderService_Create_Success(t *testing.T) {
	mockDB := new(MockDB)
	mockRepo := new(MockOrderRepository)
	mockGateway := new(MockProductGateway)
	mockProducer := new(MockOrderProducer)

	mockDB.On("Commit").Return(mockDB)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
	mockGateway.On("GetProductInfo", "prod-1").Return(&model.ProductResponse{
		Id:    "prod-1",
		Price: 1000,
		Qty:   10,
	}, nil)
	mockProducer.On("Send", "order.created", mock.Anything).Return()

	orderService := &service.OrderService{
		DB:              mockDB,
		Validate:        &FakeValidator{},
		OrderRepository: mockRepo,
		ProductGateway:  mockGateway,
		OrderProducer:   mockProducer,
	}

	req := &model.CreateOrderRequest{ProductId: "prod-1"}
	resp, err := orderService.Create(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "prod-1", resp.ProductId)
	mockRepo.AssertExpectations(t)
	mockGateway.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}
