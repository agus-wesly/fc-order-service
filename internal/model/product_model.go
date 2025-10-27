package model

type ProductResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Qty       int64  `json:"qty"`
	CreatedAt string `josn:"createdAt"`
}
