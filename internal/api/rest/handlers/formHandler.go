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

type FormHandler struct {
	formService  service.FormService
	stockService service.StockService
	validator    validator.Validate
}

func SetupFormRoutes(rh *rest.RestHandler) {
	app := rh.App
	versionRoutes := app.Group("/api/v1")

	formRepository := repository.NewFormRepository(rh.DB)
	formService := service.NewFormService(formRepository, rh.Auth, rh.Config)

	stockRepository := repository.NewStockRepository(rh.DB)
	stockService := service.NewStockService(stockRepository, rh.Auth, rh.Config)

	handler := &FormHandler{
		formService:  formService,
		stockService: stockService,
		validator:    rh.Validator,
	}

	pvtRoutes := versionRoutes.Group("/forms", rh.Auth.Authorize)

	pvtRoutes.Post("/", handler.CreateForm)
	pvtRoutes.Get("/", handler.GetFormList)
	pvtRoutes.Get("/:id", handler.GetForm)
	pvtRoutes.Put("/:id", handler.UpdateForm)
	pvtRoutes.Delete("/:id", handler.DeleteForm)

	pvtRoutes.Post("/:id/stocks", handler.CreateStock)
	pvtRoutes.Get("/:id/stocks", handler.GetStockList)
	pvtRoutes.Get("/:id/stocks/:stock_id", handler.GetStock)
	pvtRoutes.Put("/:id/stocks/:stock_id", handler.UpdateStock)
	pvtRoutes.Delete("/:id/stocks/:stock_id", handler.DeleteStock)
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

	currentUser := h.formService.Auth.GetCurrentUser(ctx)
	res, err := h.formService.CreateForm(currentUser.ID, form)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) GetFormList(ctx *fiber.Ctx) error {
	currentUser := h.formService.Auth.GetCurrentUser(ctx)
	forms, err := h.formService.FindForms(currentUser.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, forms)
}

func (h *FormHandler) GetForm(ctx *fiber.Ctx) error {
	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)
	currentUser := h.formService.Auth.GetCurrentUser(ctx)

	form, err := h.formService.FindFormByID(formObjectID, currentUser.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, form)
}

func (h *FormHandler) UpdateForm(ctx *fiber.Ctx) error {
	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	form := dto.UpdateFormInput{}

	err := ctx.BodyParser(&form)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(form); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	currentUser := h.formService.Auth.GetCurrentUser(ctx)
	res, err := h.formService.UpdateForm(formObjectID, currentUser.ID, form)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) DeleteForm(ctx *fiber.Ctx) error {
	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	currentUser := h.formService.Auth.GetCurrentUser(ctx)
	res, err := h.formService.DeleteForm(formObjectID, currentUser.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) CreateStock(ctx *fiber.Ctx) error {
	stock := dto.AddStockRequestDto{}

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	err := ctx.BodyParser(&stock)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(stock); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	res, err := h.stockService.CreateStock(formObjectID, stock)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) GetStockList(ctx *fiber.Ctx) error {
	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.stockService.FindStocks(formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) GetStock(ctx *fiber.Ctx) error {
	stockID := ctx.Params("stock_id")
	stockObjectID, _ := primitive.ObjectIDFromHex(stockID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.stockService.FindStockByID(stockObjectID, formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) UpdateStock(ctx *fiber.Ctx) error {
	stockID := ctx.Params("stock_id")
	stockObjectID, _ := primitive.ObjectIDFromHex(stockID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	stock := dto.UpdateStockRequestDto{}

	err := ctx.BodyParser(&stock)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	if err := h.validator.Struct(stock); err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	res, err := h.stockService.UpdateStock(stockObjectID, formObjectID, stock)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}

func (h *FormHandler) DeleteStock(ctx *fiber.Ctx) error {
	stockID := ctx.Params("stock_id")
	stockObjectID, _ := primitive.ObjectIDFromHex(stockID)

	formID := ctx.Params("id")
	formObjectID, _ := primitive.ObjectIDFromHex(formID)

	res, err := h.stockService.DeleteStock(stockObjectID, formObjectID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, res)
}
