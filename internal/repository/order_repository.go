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

type OrderRepository struct {
	cache *redis.Client
}

func NewOrderRepository(cache *redis.Client) *OrderRepository {
	return &OrderRepository{
		cache: cache,
	}
}

func (r *OrderRepository) Create(db *gorm.DB, entity *entity.Order) error {
	return db.Create(entity).Error
}

func (r *OrderRepository) FindByProductId(db *gorm.DB, ctx context.Context, orders *[]entity.Order, productId string) error {
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

func (r *OrderRepository) getKeyWithPrefix(category string, key string) string {
	return fmt.Sprintf("%s-%s: %s", basePrefix, category, key)
}
