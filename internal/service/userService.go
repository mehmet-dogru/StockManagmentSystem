package service

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/internal/domain"
	"DynamicStockManagmentSystem/internal/dto"
	"DynamicStockManagmentSystem/internal/helper"
	"DynamicStockManagmentSystem/internal/repository"
	"errors"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func NewUserService(repo repository.UserRepository, auth helper.Auth, config config.AppConfig) UserService {
	return UserService{
		Repo:   repo,
		Auth:   auth,
		Config: config,
	}
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Password:  hPassword,
	})

	return s.Auth.GenerateToken(user.ID.Hex(), user.Username, user.FirstName, user.LastName)
}

func (s UserService) findUserByUsername(username string) (*domain.User, error) {
	user, err := s.Repo.FindUser(username)
	return &user, err
}

func (s UserService) Login(username string, password string) (string, error) {
	user, err := s.findUserByUsername(username)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID.Hex(), user.Username, user.FirstName, user.LastName)
}
