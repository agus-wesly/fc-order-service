package converter

import (
	"order-service/internal/model"
	"order-service/internal/entity"
)

func OrderToResponse(order *entity.Order) *model.OrderResponse {
	return &model.OrderResponse {
		Id: order.Id,
		ProductId: order.ProductId,
		TotalPrice: order.TotalPrice,
		Status: order.Status,
	}
}
