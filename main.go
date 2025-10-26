package main

import (
	"fmt"
	"log"
	"order-service/config"
)

func main() {
	app := config.NewFiber()
	db := config.NewDatabase()

	config.Bootstrap(&config.BootstrapConfig{
		DB: db,
		App: app,
	})

	err := app.Listen(fmt.Sprintf(":%d", config.APP_PORT))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
