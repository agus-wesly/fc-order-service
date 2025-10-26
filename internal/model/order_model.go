package model

type CreateOrderRequest struct {
	ProductId  string `json:"product_id" validate:"required,uuid"`
	TotalPrice int64  `json:"total_price" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type OrderResponse struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}
