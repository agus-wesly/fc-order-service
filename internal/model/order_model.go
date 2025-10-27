package model

type CreateOrderRequest struct {
	ProductId  string `json:"productId" validate:"required,uuid"`
}

type OrderResponse struct {
	Id         string `json:"id"`
	ProductId  string `json:"productId"`
	TotalPrice int64    `json:"totalPrice"`
	Status     string `json:"status"`
}
