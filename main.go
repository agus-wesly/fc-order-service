package main

import (
	"fmt"
	"log"
	"order-service/config"

	"order-service/db/migrations"
)

func main() {
	app := config.NewFiber()
	db := config.NewDatabase()
	// TODO : setup defer to close db
	validate := config.NewValidator()
	cache := config.NewRedis()
	// TODO : setup defer to close cache

	migrations.Start()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validate,
		Cache:    cache,
	})

	err := app.Listen(fmt.Sprintf(":%d", config.APP_PORT))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
