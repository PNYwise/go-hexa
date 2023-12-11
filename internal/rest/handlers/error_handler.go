package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type badRequestResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Message
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	switch code {
	case 400:
		badRequest := &badRequestResponse{
			Code:    code,
			Message: fiber.ErrBadGateway.Message,
			Error:   message,
		}
		ctx.Status(code).JSON(badRequest)
	default:
		ctx.Status(code).JSON(e)
	}
	return nil
}
