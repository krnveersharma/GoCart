package api

import (
	"GoCart/config"
	"GoCart/internal/api/rest"
	"GoCart/internal/api/rest/handlers"
	"GoCart/internal/domain"
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
	log.Print(db)

	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
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
