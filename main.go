package main

import (
	"fmt"
	"log"
	"order-service/config"
)

func main() {
	app := config.NewFiber()

	config.Bootstrap(&config.BootstrapConfig{
		App: app,
	})

	err := app.Listen(fmt.Sprintf(":%d", config.APP_PORT))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
