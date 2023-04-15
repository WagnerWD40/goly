package controllers

import (
	"app/model"
	"app/server/errorhandler"
	"app/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRedirects(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGollies()

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "error getting all goly links")
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func GetGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "Could not parse url id")
	}

	goly, err := model.GetGoly(id)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "error getting goly link")
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func CreateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goly model.Goly

	err := ctx.BodyParser(&goly)

	if err != nil {
		return errorhandler.BadRequest(ctx, err, "error parsing JSON")
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "could not create goly in DB")
	}

	return ctx.Status(fiber.StatusCreated).JSON(goly)

}

func UpdateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goly model.Goly

	err := ctx.BodyParser(&goly)

	if err != nil {
		return errorhandler.BadRequest(ctx, err, "error parsing JSON")
	}

	err = model.UpdateGoly(goly)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "could not update link in DB")
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func DeleteGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)

	if err != nil {
		return errorhandler.BadRequest(ctx, err, "could not parse ir from url")
	}

	err = model.DeleteGoly(id)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "could not delete goly in DB")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted",
	})
}

func Redirect(ctx *fiber.Ctx) error {
	golyurl := ctx.Params("redirect")

	goly, err := model.FindByGolyUrl(golyurl)

	if err != nil {
		return errorhandler.InternalServerError(ctx, err, "could not find goly in DB")
	}

	goly.Clicked += 1

	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Println("error updating goly status " + golyurl)
	}

	return ctx.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}
