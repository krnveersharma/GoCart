package handlers

import (
	"GoCart/internal/api/rest"
	"GoCart/internal/repository"
	"GoCart/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := CatalogHandler{
		svc: svc,
	}

	app.Get("/products")
	app.Get("/products/:id")
	app.Get("/categories")
	app.Get("/categories/:id")

	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Patch("/products/:id", handler.UpdateStock)
	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Delete("/products/:id", handler.DeleteProduct)

}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Printf("seller is", user)

	return rest.SuccessResponse(ctx, "create category endpoint", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "edit category endpoint", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "delete category endpoint", nil)
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "create product endpoint", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "get products endpoint", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "get product endpoint", nil)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "edit product endpoint", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "update stock endpoint", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "delete product endpoint", nil)
}
