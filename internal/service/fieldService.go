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

type FieldService struct {
	Repo   repository.FieldRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func NewFieldService(repo repository.FieldRepository, auth helper.Auth, config config.AppConfig) FieldService {
	return FieldService{
		Repo:   repo,
		Auth:   auth,
		Config: config,
	}
}

func (f FieldService) CreateField(formID primitive.ObjectID, input dto.CreateFieldInput) (string, error) {
	var err error
	var fieldType string

	switch input.Type {
	case "Combobox":
		fieldType = "Combobox"
	case "Text":
		fieldType = "Text"
	case "Checkbox":
		fieldType = "Checkbox"
	case "Number", "NumberDecimal":
		fieldType = "Number"
	default:
		return "", errors.New("invalid field type")
	}

	switch fieldType {
	case "Combobox":
		if len(input.Options) == 0 || len(input.Options) > 10 {
			return "", errors.New("invalid options for Combobox field")
		}
	case "Text":
		if input.MinChars < 0 || input.MaxChars < 0 || input.MinChars > input.MaxChars {
			return "", errors.New("invalid min-max characters for Text field")
		}
	case "Number":
		if input.MinValue < 0 || input.MaxValue < 0 || input.MinValue > input.MaxValue {
			return "", errors.New("invalid min-max values for Number field")
		}
	default:
		if input.MinValue < 0 || input.MaxValue < 0 || input.MinValue > input.MaxValue {
			return "", errors.New("invalid min-max values for NumberDecimal field")
		}
	}

	_, err = f.Repo.CreateField(domain.Field{
		FormID:     formID,
		Type:       input.Type,
		Name:       input.Name,
		Options:    input.Options,
		MinChars:   input.MinChars,
		MaxChars:   input.MaxChars,
		MinValue:   input.MinValue,
		MaxValue:   input.MaxValue,
		IsRequired: input.IsRequired,
		IsUnique:   *input.IsUnique,
		IsHidden:   *input.IsHidden,
		Order:      input.Order,
	})
	if err != nil {
		return "", err
	}

	return "field created successfully", nil
}

func (f FieldService) FindFields(formID primitive.ObjectID) ([]domain.Field, error) {
	fields, err := f.Repo.GetFields(formID)
	if err != nil {
		return []domain.Field{}, err
	}

	return fields, nil
}

func (f FieldService) FindFieldByID(fieldID primitive.ObjectID, formID primitive.ObjectID) (domain.Field, error) {
	field, err := f.Repo.GetFieldByID(fieldID, formID)
	if err != nil {
		return domain.Field{}, err
	}

	return field, nil
}

func (f FieldService) UpdateField(fieldID primitive.ObjectID, formID primitive.ObjectID, input dto.UpdateFieldInput) (string, error) {
	err := f.Repo.UpdateField(fieldID, domain.Field{
		Type:       input.Type,
		Name:       input.Name,
		Options:    input.Options,
		MinChars:   input.MinChars,
		MaxChars:   input.MaxChars,
		MinValue:   input.MinValue,
		MaxValue:   input.MaxValue,
		IsRequired: input.IsRequired,
		IsUnique:   input.IsUnique,
		IsHidden:   input.IsHidden,
		Order:      input.Order,
	}, formID)
	if err != nil {
		return "", err
	}

	return "field updated successfully", nil
}

func (f FieldService) DeleteField(fieldID primitive.ObjectID, formID primitive.ObjectID) (string, error) {
	deleteResult, err := f.Repo.DeleteField(fieldID, formID)
	if err != nil {
		return "", err
	}

	if deleteResult == 0 {
		return "", errors.New("field does not exist")
	}

	return "field deleted successfully", nil
}
