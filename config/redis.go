package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"order-service/pkg/dotenv"
)

func NewRedis() *redis.Client {
	val := fmt.Sprintf("redis://%s:%s", dotenv.Getenv("REDIS_HOST"), dotenv.Getenv("REDIS_PORT"))
	opt, err := redis.ParseURL(
		val,
	)
	if err != nil {
		log.Fatalln("Failed to setup redis")
	}

	client := redis.NewClient(opt)
	return client
}
