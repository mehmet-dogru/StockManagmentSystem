package handlers

import (
	"DynamicStockManagmentSystem/internal/api/rest"
	"DynamicStockManagmentSystem/internal/api/rest/responses"
	"DynamicStockManagmentSystem/internal/dto"
	"DynamicStockManagmentSystem/internal/repository"
	"DynamicStockManagmentSystem/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type FieldHandler struct {
	svc       service.FieldService
	validator validator.Validate
}

func SetupFieldRoutes(rh *rest.RestHandler) {
	app := rh.App
	versionRoutes := app.Group("/api/v1")

	repo := repository.NewFieldRepository(rh.DB)
	svc := service.NewFieldService(repo, rh.Auth, rh.Config)

	handler := &FieldHandler{
		svc:       svc,
		validator: rh.Validator,
	}

	pvtRoutes := versionRoutes.Group("/forms/:id", rh.Auth.Authorize)

	pvtRoutes.Post("/field", handler.CreateField)
	pvtRoutes.Get("/field", handler.GetFieldList)
	pvtRoutes.Get("/field/:field_id", handler.GetField)
	pvtRoutes.Put("/field/:id", handler.UpdateField)
	pvtRoutes.Delete("/field/:id", handler.DeleteField)
}

func (h *FieldHandler) CreateField(ctx *fiber.Ctx) error {
	field := dto.CreateFieldInput{}

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	err := ctx.BodyParser(&field)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(field); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	res, err := h.svc.CreateField(formObjectID, field)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FieldHandler) GetFieldList(ctx *fiber.Ctx) error {
	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.svc.FindFields(formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FieldHandler) GetField(ctx *fiber.Ctx) error {
	fieldID := ctx.Params("field_id")
	fieldObjectID, _ := primitive.ObjectIDFromHex(fieldID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.svc.FindFieldByID(fieldObjectID, formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FieldHandler) UpdateField(ctx *fiber.Ctx) error {
	fieldID := ctx.Params("id")
	fieldObjectID, _ := primitive.ObjectIDFromHex(fieldID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	field := dto.UpdateFieldInput{}

	err := ctx.BodyParser(&field)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(field); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	res, err := h.svc.UpdateField(fieldObjectID, formObjectID, field)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FieldHandler) DeleteField(ctx *fiber.Ctx) error {
	fieldID := ctx.Params("id")
	fieldObjectID, _ := primitive.ObjectIDFromHex(fieldID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.svc.DeleteField(fieldObjectID, formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}
