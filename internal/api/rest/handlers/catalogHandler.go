package handlers

import (
	"GoCart/internal/api/rest"
	"GoCart/internal/dto"
	"GoCart/internal/repository"
	"GoCart/internal/service"
	"strconv"

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

	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

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

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cats, err := h.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "categories", cats)
}

func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	cat, err := h.svc.GetCategory(id)
	if err != nil {
		return err
	}

	return rest.SuccessResponse(ctx, "get category by id endpoint", nil)
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {

	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(req)
	if err != nil {
		return rest.BadRequest(ctx, "create category request is not valid")
	}

	err = h.svc.CreateCategory(req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category created successfuly", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(req)
	if err != nil {
		return rest.BadRequest(ctx, "create category request is not valid")
	}

	updatedCategory, err = h.svc.EditCategory(id, req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "edit category endpoint", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
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
