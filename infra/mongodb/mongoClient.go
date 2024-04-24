package mongodb

import (
	"DynamicStockManagmentSystem/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectMongoDB(cfg config.AppConfig) *mongo.Database {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("mongo connection error: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("mongo ping error: %v", err)
	}

	log.Println("mongo connection success")

	return client.Database(cfg.MongoDbName)
}
