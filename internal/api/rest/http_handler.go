package rest

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/internal/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestHandler struct {
	App       *fiber.App
	DB        *mongo.Database
	Auth      helper.Auth
	Config    config.AppConfig
	Validator validator.Validate
}
