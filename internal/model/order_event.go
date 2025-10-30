package model

type OrderEvent struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
}

func (c *OrderEvent) GetId() string {
	return c.Id
}

type OrderEventPayload struct {
	Pattern string     `json:"pattern"`
	Data    OrderEvent `json:"data"`
}
