package api

import (
	"GoCart/config"
	"net/http"

	"github.com/gofiber/fiber"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)
	app.Listen("localhost:9000")
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "I am Breathing!",
	})
}
