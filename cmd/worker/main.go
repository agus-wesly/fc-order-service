package main

import (
	"context"
	"log"
	"order-service/config"
	"os"
	"os/signal"
	"syscall"
	"time"
	"order-service/internal/delivery/messaging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go RunOrderConsumer(ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			log.Println("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}

func RunOrderConsumer(ctx context.Context) {
	channel := config.NewRabbitMqConsumer()
	orderHandler := messaging.NewOrderHandler()
	messaging.ConsumeTopic(ctx, channel, "order.created", orderHandler.OrderCreated)
}
