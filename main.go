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
	producer := config.NewRabbitMqProducer()

	{
		log.Println("Successfully connected to RabbitMQ")
		log.Println("Waiting for messages")
		messages, err := producer.Channel.Consume(
			"app_queue", // queue name
			"",          // consumer
			true,        // auto-ack
			false,       // exclusive
			false,       // no local
			false,       // no wait
			nil,         // arguments
		)
		if err != nil {
			log.Println(err)
		}
		go func() {
			for message := range messages {
				// For example, show received message in a console.
				log.Printf(" > Received message: %s with type : %s\n", message.Body, message.Type)
			}
		}()
	}

	migrations.Start()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validate,
		Cache:    cache,
		Producer: producer,
	})

	err := app.Listen(fmt.Sprintf(":%d", config.APP_PORT))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
