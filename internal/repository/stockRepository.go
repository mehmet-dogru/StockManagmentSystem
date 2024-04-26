package repository

import (
	"DynamicStockManagmentSystem/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type StockRepository interface {
	AddStock(formID primitive.ObjectID, stock domain.Stock) (domain.Stock, error)
	GetStockList(formID primitive.ObjectID) ([]domain.Stock, error)
	GetStock(stockID primitive.ObjectID, formID primitive.ObjectID) (domain.Stock, error)
	UpdateStock(stockID primitive.ObjectID, stock domain.Stock, formID primitive.ObjectID) error
	DeleteStock(stockID primitive.ObjectID, formID primitive.ObjectID) (int64, error)
}

type stockRepository struct {
	collection *mongo.Collection
}

func NewStockRepository(db *mongo.Database) StockRepository {
	return &stockRepository{
		collection: db.Collection("stocks"),
	}
}

func (s stockRepository) AddStock(formID primitive.ObjectID, stock domain.Stock) (domain.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stock.ID = primitive.NewObjectID()
	stock.FormID = formID

	_, err := s.collection.InsertOne(ctx, stock)
	if err != nil {
		return domain.Stock{}, err
	}

	return stock, nil
}

func (s stockRepository) GetStockList(formID primitive.ObjectID) ([]domain.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, primitive.M{"formId": formID})
	if err != nil {
		return []domain.Stock{}, err
	}

	var stocks []domain.Stock
	for cursor.Next(ctx) {
		var stock domain.Stock
		err := cursor.Decode(&stock)
		if err != nil {
			return []domain.Stock{}, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (s stockRepository) GetStock(stockID primitive.ObjectID, formID primitive.ObjectID) (domain.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var stock domain.Stock
	err := s.collection.FindOne(ctx, primitive.M{"_id": stockID, "formId": formID}).Decode(&stock)
	if err != nil {
		return domain.Stock{}, err
	}

	return stock, nil
}

func (s stockRepository) UpdateStock(stockID primitive.ObjectID, stock domain.Stock, formID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.collection.UpdateOne(ctx, primitive.M{"_id": stockID, "formId": formID}, primitive.M{"$set": bson.M{
		"productName": stock.ProductName,
		"quantity":    stock.Quantity,
		"price":       stock.Price,
		"currency":    stock.Currency,
		"isAvailable": stock.IsAvailable,
	}})
	if err != nil {
		return err
	}

	return nil
}

func (s stockRepository) DeleteStock(stockID primitive.ObjectID, formID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.collection.DeleteOne(ctx, primitive.M{"_id": stockID, "formId": formID})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
