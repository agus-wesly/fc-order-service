package converter

import (
	"order-service/internal/entity"
	"order-service/internal/model"
)

func OrderToResponse(order *entity.Order) *model.OrderResponse {
	return &model.OrderResponse{
		Id:         order.Id,
		ProductId:  order.ProductId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
	}
}

func OrdersToResponse(ordersPtr *[]entity.Order) *model.OrderListResponse {
	orders := *ordersPtr
	orderResponses := make([]*model.OrderResponse, len(orders))
	for i := range len(orders) {
		orderResponses[i] = OrderToResponse(&orders[i])
	}

	return &model.OrderListResponse{
		Orders: orderResponses,
	}
}
