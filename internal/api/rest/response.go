package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(err.Error())
}

func InternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func BadRequest(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": msg,
	})
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data interface{}) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": msg,
		"data":    data,
	})
}
