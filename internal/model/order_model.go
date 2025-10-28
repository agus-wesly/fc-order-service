package model

type CreateOrderRequest struct {
	ProductId string `json:"productId" validate:"required,uuid"`
}

type GetOrderByProductIdRequest struct {
	Id string `json:"-" validate:"required"`
}

type GetOrderByIdRequest struct {
	Id string `json:"-" validate:"required"`
}

type OrderResponse struct {
	Id         string `json:"id"`
	ProductId  string `json:"productId"`
	TotalPrice int64  `json:"totalPrice"`
	Status     string `json:"status"`
	CreatedAt  int64 `json:"createdAt"`
}

type GetOrdersByProductIdResponse struct {
	Orders []*OrderResponse `json:"orders"`
}
