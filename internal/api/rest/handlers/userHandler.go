package handlers

import (
	"GoCart/internal/api/rest"
	"GoCart/internal/dto"
	"GoCart/internal/repository"
	"GoCart/internal/service"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")

	//Public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	//Private Endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Post("/profile", handler.CreateProfile)

	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)
	pvtRoutes.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid details",
		})
	}
	_, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	if err := ctx.BodyParser(&loginInput); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid details",
		})
	}

	email, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": email,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	err := h.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Code sent to your registered Email",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	var req dto.VerificationCodeInput
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfuly",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile ",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create order",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		fmt.Printf("BecomeSeller:error in parsing request body %v", err)
		return ctx.Status((http.StatusBadRequest)).JSON(&fiber.Map{
			"message": "request parameters are not valid",
		})
	}

	token, err := h.svc.BecomeSeller(user.ID, req)
	if err != nil {
		fmt.Printf("BecomeSeller:error in generating token %v", err)
		return ctx.Status((http.StatusBadRequest)).JSON(&fiber.Map{
			"message": "failed to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"token":   token,
	})
}
