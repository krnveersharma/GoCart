package api

import (
	"GoCart/config"
	"GoCart/internal/api/rest"
	"GoCart/internal/api/rest/handlers"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	fmt.Println("config is %v", config.Dsn)
	rh := &rest.RestHandler{
		App: app,
	}
	SetupRoutes(rh)

	app.Listen(config.ServerPort)
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
