package model

type CreateOrderRequest struct {
	ProductId  string `json:"product_id" validate:"required,uuid"`
}

type OrderResponse struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	TotalPrice int64    `json:"total_price"`
	Status     string `json:"status"`
}
