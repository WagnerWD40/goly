package controllers

import (
	"app/model"
	"app/server/errorhandler"
	"app/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRedirects(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGollies()

	if err != nil {
		return errorhandler.InternalServerErrorHandler(ctx, err, "error getting all goly links")
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func GetGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)

	if err != nil {
		return errorhandler.InternalServerErrorHandler(ctx, err, "Could not parse url id")
	}

	goly, err := model.GetGoly(id)

	if err != nil {
		return errorhandler.InternalServerErrorHandler(ctx, err, "error getting goly link")
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func CreateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goly model.Goly

	err := ctx.BodyParser(&goly)

	if err != nil {
		return errorhandler.InternalServerErrorHandler(ctx, err, "error parsing JSON")
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)

	if err != nil {
		return errorhandler.InternalServerErrorHandler(ctx, err, "could not create goly in DB")
	}

	return ctx.Status(fiber.StatusCreated).JSON(goly)

}
