package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Order struct {
	ProductId  string `json:"product_id"`
	TotalPrice int `json:"total_price"`
	Status     string `json:"status"`
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	app.Post("/orders", func(c *fiber.Ctx) error {
		order := new(Order)
		if err := c.BodyParser(order); err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(order)
	})

	app.Listen(":5959")
}
