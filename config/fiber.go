package config

import (
	"github.com/gofiber/fiber/v2"
)

const APP_NAME = "order-service"

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: APP_NAME,
		ErrorHandler: NewErrorHandler(),
	})
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
