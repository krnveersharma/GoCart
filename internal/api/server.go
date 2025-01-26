package api

import (
	"GoCart/config"
	"GoCart/internal/api/rest"
	"GoCart/internal/api/rest/handlers"
	"net/http"

	"github.com/gofiber/fiber"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	rh := &rest.RestHandler{
		App: app,
	}
	SetupRoutes(rh)
	app.Get("/health", HealthCheck)
	app.Listen("localhost:9000")
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "I am Breathing!",
	})
}

func SetupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)
	// transactions

	// catalog
}
