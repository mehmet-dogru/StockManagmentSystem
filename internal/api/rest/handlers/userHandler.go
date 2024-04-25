package handlers

import (
	"DynamicStockManagmentSystem/internal/api/rest"
	"DynamicStockManagmentSystem/internal/api/rest/responses"
	"DynamicStockManagmentSystem/internal/dto"
	"DynamicStockManagmentSystem/internal/repository"
	"DynamicStockManagmentSystem/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	svc       service.UserService
	validator validator.Validate
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App
	versionRoutes := app.Group("/api/v1")

	repo := repository.NewUserRepository(rh.DB)
	svc := service.NewUserService(repo, rh.Auth, rh.Config)

	handler := &UserHandler{
		svc:       svc,
		validator: rh.Validator,
	}

	pubRoutes := versionRoutes.Group("/users")

	//Public Endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	//Protected Endpoints
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	pvtRoutes.Get("/account", handler.GetAccount)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(user); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(loginInput); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	token, err := h.svc.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

func (h *UserHandler) GetAccount(ctx *fiber.Ctx) error {
	currentUser := h.svc.Auth.GetCurrentUser(ctx)
	accountInfo, err := h.svc.GetProfile(currentUser.ID)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, accountInfo)
}
