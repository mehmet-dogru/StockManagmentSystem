package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort  string
	MongoURI    string
	MongoDbName string
	AppSecret   string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if len(mongoURI) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	mongoDbName := os.Getenv("MONGO_DB_NAME")
	if len(mongoDbName) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	return AppConfig{
		ServerPort:  httpPort,
		MongoURI:    mongoURI,
		AppSecret:   appSecret,
		MongoDbName: mongoDbName,
	}, nil
}
