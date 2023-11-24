package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	switch code {
	case 404:
		ctx.Status(code).JSON(fiber.Map{"message": message})
	default:
		ctx.Status(code).JSON(fiber.Map{"message": message})
	}
	return nil
}
