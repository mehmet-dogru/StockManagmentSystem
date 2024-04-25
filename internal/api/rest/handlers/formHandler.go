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

type FormHandler struct {
	svc       service.FormService
	validator validator.Validate
}

func SetupFormRoutes(rh *rest.RestHandler) {
	app := rh.App
	versionRoutes := app.Group("/api/v1")

	repo := repository.NewFormRepository(rh.DB)
	svc := service.NewFormService(repo, rh.Auth, rh.Config)

	handler := &FormHandler{
		svc:       svc,
		validator: rh.Validator,
	}

	pvtRoutes := versionRoutes.Group("/forms", rh.Auth.Authorize)

	pvtRoutes.Post("/", handler.CreateForm)
	pvtRoutes.Get("/", handler.GetFormList)
}

func (h *FormHandler) CreateForm(ctx *fiber.Ctx) error {
	form := dto.CreateFormInput{}

	err := ctx.BodyParser(&form)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(form); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	currentUser := h.svc.Auth.GetCurrentUser(ctx)
	res, err := h.svc.CreateForm(currentUser.ID, form)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) GetFormList(ctx *fiber.Ctx) error {
	currentUser := h.svc.Auth.GetCurrentUser(ctx)
	forms, err := h.svc.FindForms(currentUser.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, forms)
}
