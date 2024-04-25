package service

import (
	"DynamicStockManagmentSystem/config"
	"DynamicStockManagmentSystem/internal/domain"
	"DynamicStockManagmentSystem/internal/dto"
	"DynamicStockManagmentSystem/internal/helper"
	"DynamicStockManagmentSystem/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormService struct {
	Repo   repository.FormRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func NewFormService(repo repository.FormRepository, auth helper.Auth, config config.AppConfig) FormService {
	return FormService{
		Repo:   repo,
		Auth:   auth,
		Config: config,
	}
}

func (s FormService) CreateForm(userID primitive.ObjectID, input dto.CreateFormInput) (string, error) {
	_, err := s.Repo.CreateForm(domain.Form{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
	})
	if err != nil {
		return "", err
	}

	return "form created successfully", nil
}

func (s FormService) FindForms(userID primitive.ObjectID) ([]domain.Form, error) {
	forms, err := s.Repo.GetForms(userID)
	if err != nil {
		return []domain.Form{}, err
	}

	return forms, nil
}

func (s FormService) FindFormByID(formID primitive.ObjectID, userID primitive.ObjectID) (domain.Form, error) {
	form, err := s.Repo.GetFormByID(formID, userID)
	if err != nil {
		return domain.Form{}, err
	}

	return form, nil
}

func (s FormService) UpdateForm(formID primitive.ObjectID, userID primitive.ObjectID, input dto.UpdateFormInput) (string, error) {
	err := s.Repo.UpdateForm(formID, domain.Form{
		Title:       input.Title,
		Description: input.Description,
	}, userID)
	if err != nil {
		return "", err
	}

	return "form updated successfully", nil
}

func (s FormService) DeleteForm(formID primitive.ObjectID, userID primitive.ObjectID) (string, error) {
	err := s.Repo.DeleteForm(formID, userID)
	if err != nil {
		return "", err
	}

	return "form deleted successfully", nil
}
