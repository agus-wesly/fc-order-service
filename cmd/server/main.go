package main

import (
	"fmt"
	"log"
	"order-service/pkg/dotenv"
	"order-service/config"

	"order-service/db/migrations"
)

func main() {
	app := config.NewFiber()
	db := config.NewDatabase()
	validate := config.NewValidator()
	cache := config.NewRedis()
	producer := config.NewRabbitMqProducer()

	migrations.Start()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validate,
		Cache:    cache,
		Producer: producer,
	})

	// Resource cleanup
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}

		cache.Close()
		producer.Connection.Close()
		producer.Channel.Close()
	}()

	err := app.Listen(fmt.Sprintf(":%s", dotenv.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
