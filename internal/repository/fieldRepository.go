package repository

import (
	"DynamicStockManagmentSystem/internal/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FieldRepository interface {
	CreateField(field domain.Field) (domain.Field, error)
	GetFields(formID primitive.ObjectID) ([]domain.Field, error)
	GetFieldByID(fieldID primitive.ObjectID, formID primitive.ObjectID) (domain.Field, error)
	UpdateField(fieldID primitive.ObjectID, field domain.Field, formID primitive.ObjectID) error
	DeleteField(fieldID primitive.ObjectID, formID primitive.ObjectID) (int64, error)
}

type fieldRepository struct {
	collection *mongo.Collection
}

func NewFieldRepository(db *mongo.Database) FieldRepository {
	return &fieldRepository{
		collection: db.Collection("fields"),
	}
}

func (f fieldRepository) CreateField(field domain.Field) (domain.Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	field.ID = primitive.NewObjectID()

	res, err := f.collection.InsertOne(ctx, field)
	if res.InsertedID == nil || err != nil {
		log.Printf("create field error %v", err)
		return domain.Field{}, errors.New("failed to create field")
	}

	return field, nil
}

func (f fieldRepository) GetFields(formID primitive.ObjectID) ([]domain.Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := f.collection.Find(ctx, primitive.M{"formId": formID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []domain.Field{}, errors.New("field does not exist")
		}
		log.Printf("get fields error %v", err)
		return []domain.Field{}, errors.New("failed to get fields")
	}

	var fields []domain.Field
	for cursor.Next(ctx) {
		var field domain.Field
		err := cursor.Decode(&field)
		if err != nil {
			log.Printf("decode field error %v", err)
			return []domain.Field{}, errors.New("failed to decode field")
		}
		fields = append(fields, field)
	}

	return fields, nil

}

func (f fieldRepository) GetFieldByID(fieldID primitive.ObjectID, formID primitive.ObjectID) (domain.Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var field domain.Field
	err := f.collection.FindOne(ctx, primitive.M{"_id": fieldID, "formId": formID}).Decode(&field)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Field{}, errors.New("field does not exist")
		}
		log.Printf("get field error %v", err)
		return domain.Field{}, errors.New("failed to get field")
	}

	return field, nil
}

func (f fieldRepository) UpdateField(fieldID primitive.ObjectID, field domain.Field, formID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := primitive.M{
		"$set": primitive.M{
			"name":       field.Name,
			"type":       field.Type,
			"options":    field.Options,
			"minChars":   field.MinChars,
			"maxChars":   field.MaxChars,
			"minValue":   field.MinValue,
			"maxValue":   field.MaxValue,
			"isRequired": field.IsRequired,
			"isUnique":   field.IsUnique,
			"isHidden":   field.IsHidden,
			"order":      field.Order,
		},
	}

	_, err := f.collection.UpdateOne(ctx, primitive.M{"_id": fieldID, "formId": formID}, update)
	if err != nil {
		log.Printf("update field error %v", err)
		return errors.New("failed to update field")
	}

	return nil
}

func (f fieldRepository) DeleteField(fieldID primitive.ObjectID, formID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := f.collection.DeleteOne(ctx, primitive.M{"_id": fieldID, "formId": formID})
	if err != nil {
		log.Printf("delete field error %v", err)
		return 0, errors.New("failed to delete field")
	}

	return res.DeletedCount, nil
}
