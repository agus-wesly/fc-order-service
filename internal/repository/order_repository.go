package repository

import (
	"order-service/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(db *gorm.DB, entity *entity.Order) error {
	return db.Create(entity).Error
}

func (r *OrderRepository) FindByProductId(db *gorm.DB, order *entity.Order, id string) error {
	return db.Where("product_id = ?", id).Take(order).Error
}
