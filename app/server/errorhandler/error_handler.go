package errorhandler

import (
	"github.com/gofiber/fiber/v2"
)

func InternalServerErrorHandler(ctx *fiber.Ctx, err error, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message + " " + err.Error(),
	})
}
