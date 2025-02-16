package api

import (
	"GoCart/config"
	"GoCart/internal/api/rest"
	"GoCart/internal/api/rest/handlers"
	"GoCart/internal/domain"
	"GoCart/internal/helper"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database is not connected %s", err.Error())
	}
	log.Println("database connected")

	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{})
	if err != nil {
		log.Fatalf("Database is not migrated %s", err.Error())
	}

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
	}
	SetupRoutes(rh)

	err = app.Listen(config.ServerPort)
	if err != nil {
		log.Fatalf("Server is not running %s", err.Error())
	}
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
