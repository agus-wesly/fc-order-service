package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"order-service/internal/entity"
	"time"

	"gorm.io/gorm"
)

const (
	basePrefix = "order"
	tTLSecond  = 5
)

type OrderRepository interface {
	Create(db *gorm.DB, entity *entity.Order) error
	FindByProductId(db *gorm.DB, ctx context.Context, orders *[]entity.Order, productId string) error
	List(db *gorm.DB, ctx context.Context, orders *[]entity.Order) error
	FindById(db *gorm.DB, ctx context.Context, order *entity.Order, id string) error
}

type orderRepository struct {
	cache *redis.Client
}

func NewOrderRepository(cache *redis.Client) OrderRepository {
	return &orderRepository{
		cache: cache,
	}
}

func (r *orderRepository) Create(db *gorm.DB, entity *entity.Order) error {
	return db.Create(entity).Error
}

func (r *orderRepository) FindByProductId(db *gorm.DB, ctx context.Context, orders *[]entity.Order, productId string) error {
	key := r.getKeyWithPrefix("productId", productId)
	ordersBytes, err := r.cache.Get(ctx, key).Bytes()
	if err == nil {
		if err = json.Unmarshal(ordersBytes, orders); err != nil {
			return err
		}
		return nil
	}

	if err := db.Where("product_id = ?", productId).Find(orders).Error; err != nil {
		return err
	}
	if ordersBytes, err = json.Marshal(orders); err != nil {
		return err
	}

	if err = r.cache.Set(ctx, key, ordersBytes, time.Second*time.Duration(tTLSecond)).Err(); err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) List(db *gorm.DB, ctx context.Context, orders *[]entity.Order) error {
	return db.Find(orders).Error
}

func (r *orderRepository) FindById(db *gorm.DB, ctx context.Context, order *entity.Order, id string) error {
	key := r.getKeyWithPrefix("id", id)
	orderBytes, err := r.cache.Get(ctx, key).Bytes()
	if err == nil {
		if err = json.Unmarshal(orderBytes, order); err != nil {
			return err
		}
		return nil
	}

	if err := db.Where("id = ?", id).Take(order).Error; err != nil {
		return err
	}
	if orderBytes, err = json.Marshal(order); err != nil {
		return err
	}

	if err = r.cache.Set(ctx, key, orderBytes, time.Second*time.Duration(tTLSecond)).Err(); err != nil {
		return err
	}
	return nil

}

func (r *orderRepository) getKeyWithPrefix(category string, key string) string {
	return fmt.Sprintf("%s-%s: %s", basePrefix, category, key)
}
