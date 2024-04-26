package service

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/internal/domain"
	"DynamicStockManagmentSystem/internal/dto"
	"DynamicStockManagmentSystem/internal/helper"
	"DynamicStockManagmentSystem/internal/repository"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StockService struct {
	Repo   repository.StockRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func NewStockService(repo repository.StockRepository, auth helper.Auth, config config.AppConfig) StockService {
	return StockService{
		Repo:   repo,
		Auth:   auth,
		Config: config,
	}
}

func (s StockService) CreateStock(formID primitive.ObjectID, input dto.AddStockRequestDto) (string, error) {
	_, err := s.Repo.AddStock(formID, domain.Stock{
		ProductName: input.ProductName,
		Quantity:    input.Quantity,
		Price:       input.Price,
		Currency:    input.Currency,
		IsAvailable: input.IsAvailable,
	})
	if err != nil {
		return "", err
	}

	return "stock created successfully", nil
}

func (s StockService) FindStocks(formID primitive.ObjectID) ([]domain.Stock, error) {
	stocks, err := s.Repo.GetStockList(formID)
	if err != nil {
		return []domain.Stock{}, err
	}

	return stocks, nil
}

func (s StockService) FindStockByID(stockID primitive.ObjectID, formID primitive.ObjectID) (domain.Stock, error) {
	stock, err := s.Repo.GetStock(stockID, formID)
	if err != nil {
		return domain.Stock{}, err
	}

	return stock, nil
}

func (s StockService) UpdateStock(stockID primitive.ObjectID, formID primitive.ObjectID, input dto.UpdateStockRequestDto) (string, error) {
	err := s.Repo.UpdateStock(stockID, domain.Stock{
		ProductName: input.ProductName,
		Quantity:    input.Quantity,
		Price:       input.Price,
		Currency:    input.Currency,
		IsAvailable: input.IsAvailable,
	}, formID)
	if err != nil {
		return "", err
	}

	return "stock updated successfully", nil
}

func (s StockService) DeleteStock(stockID primitive.ObjectID, formID primitive.ObjectID) (string, error) {
	deleteResult, err := s.Repo.DeleteStock(stockID, formID)
	if err != nil {
		return "", err
	}

	if deleteResult == 0 {
		return "", errors.New("stock does not exist")
	}

	return "stock deleted successfully", nil
}
