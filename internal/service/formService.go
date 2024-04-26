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
	existingForm, err := s.Repo.GetFormByTitle(input.Title, userID)
	if err != nil {
		return "", err
	}
	if existingForm.ID != primitive.NilObjectID {
		return "", errors.New("a form with the same title already exists")
	}

	_, err = s.Repo.CreateForm(domain.Form{
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
	deleteResult, err := s.Repo.DeleteForm(formID, userID)
	if err != nil {
		return "", err
	}

	if deleteResult == 0 {
		return "", errors.New("form does not exist")
	}

	return "form deleted successfully", nil
}

func (s FormService) CountForms(userID primitive.ObjectID) (int64, error) {
	count, err := s.Repo.CountForms(userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s FormService) FindFormsPaginated(userID primitive.ObjectID, offset, limit int) ([]domain.Form, error) {
	forms, err := s.Repo.FindFormsPaginated(userID, offset, limit)
	if err != nil {
		return nil, err
	}

	return forms, nil
}
