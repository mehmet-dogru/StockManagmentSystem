package repository

import (
	"DynamicStockManagmentSystem/internal/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FormRepository interface {
	CreateForm(form domain.Form) (domain.Form, error)
	GetForms(userID primitive.ObjectID) ([]domain.Form, error)
	GetFormByID(formID primitive.ObjectID, userID primitive.ObjectID) (domain.Form, error)
	UpdateForm(formID primitive.ObjectID, form domain.Form, userID primitive.ObjectID) error
	DeleteForm(formID primitive.ObjectID, userID primitive.ObjectID) (int64, error)
}

type formRepository struct {
	collection *mongo.Collection
}

func NewFormRepository(db *mongo.Database) FormRepository {
	return &formRepository{
		collection: db.Collection("forms"),
	}
}

func (f formRepository) CreateForm(form domain.Form) (domain.Form, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	form.ID = primitive.NewObjectID()

	res, err := f.collection.InsertOne(ctx, form)
	if res.InsertedID == nil || err != nil {
		log.Printf("create form error %v", err)
		return domain.Form{}, errors.New("failed to create form")
	}

	return form, nil
}

func (f formRepository) GetForms(userID primitive.ObjectID) ([]domain.Form, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := f.collection.Find(ctx, primitive.M{"userId": userID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []domain.Form{}, errors.New("form does not exist")
		}
		log.Printf("get forms error %v", err)
		return []domain.Form{}, errors.New("failed to get forms")
	}

	var forms []domain.Form
	if err = cursor.All(ctx, &forms); err != nil {
		log.Printf("get forms error %v", err)
		return []domain.Form{}, errors.New("failed to get forms")
	}

	return forms, nil
}

func (f formRepository) GetFormByID(formID primitive.ObjectID, userID primitive.ObjectID) (domain.Form, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var form domain.Form
	err := f.collection.FindOne(ctx, primitive.M{"_id": formID, "userId": userID}).Decode(&form)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Form{}, errors.New("form does not exist")
		}
		log.Printf("get form error %v", err)
		return domain.Form{}, errors.New("failed to get form")
	}

	return form, nil
}

func (f formRepository) UpdateForm(formID primitive.ObjectID, form domain.Form, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := primitive.M{"_id": formID, "userId": userID}
	update := primitive.M{"$set": bson.M{
		"title":       form.Title,
		"description": form.Description,
	}}

	_, err := f.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("update form error %v", err)
		return errors.New("failed to update form")
	}

	return nil
}

func (f formRepository) DeleteForm(formID primitive.ObjectID, userID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := primitive.M{"_id": formID, "userId": userID}

	deleteResult, err := f.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("delete form error %v", err)
		return 0, errors.New("failed to delete form")
	}

	return deleteResult.DeletedCount, nil
}
