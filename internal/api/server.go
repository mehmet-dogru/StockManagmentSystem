package api

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/infra/mongodb"
	"DynamicStockManagmentSystem/internal/api/rest"
	"DynamicStockManagmentSystem/internal/api/rest/handlers"
	"DynamicStockManagmentSystem/internal/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	db := mongodb.ConnectMongoDB(config)

	auth := helper.SetupAuth(config.AppSecret)
	validate := validator.New()

	rh := &rest.RestHandler{
		App:       app,
		DB:        db,
		Auth:      auth,
		Config:    config,
		Validator: *validate,
	}

	setupRoutes(rh)

	log.Fatal(app.Listen(config.ServerPort))
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	handlers.SetupFormRoutes(rh)
	handlers.SetupFieldRoutes(rh)
}
