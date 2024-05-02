package api

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/infra/mongodb"
	"DynamicStockManagmentSystem/internal/api/rest"
	"DynamicStockManagmentSystem/internal/api/rest/handlers"
	"DynamicStockManagmentSystem/internal/helper"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("gracefully shutting down")
		_ = app.Shutdown()
	}()

	if err := app.Listen(config.ServerPort); err != nil {
		log.Panic(err)
	}

	fmt.Println("running cleanup tasks...")
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	handlers.SetupFormRoutes(rh)
	handlers.SetupFieldRoutes(rh)
}
