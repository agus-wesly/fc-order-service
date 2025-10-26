package entity

type Order struct {
	Id         string `gorm:"column:id;primaryKey"`
	ProductId  string `gorm:"column:product_id"`
	TotalPrice int64  `gorm:"column:total_price"`
	Status     string `gorm:"column:status"`
	CreatedAt  string `gorm:"column:created_at;autoCreateTime:milli"`
}

func (c *Order) TableName() string {
	return "orders"
}
