package errorhandler

import (
	"github.com/gofiber/fiber/v2"
)

func genericHandler(ctx *fiber.Ctx, err error, message string, errorCode int) error {
	return ctx.Status(errorCode).JSON(fiber.Map{
		"message": message + " " + err.Error(),
	})
}

func InternalServerError(ctx *fiber.Ctx, err error, message string) error {
	return genericHandler(ctx, err, message, fiber.StatusInternalServerError)
}

func BadRequest(ctx *fiber.Ctx, err error, message string) error {
	return genericHandler(ctx, err, message, fiber.StatusBadRequest)
}
