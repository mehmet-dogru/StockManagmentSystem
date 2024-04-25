package repository

import (
	"DynamicStockManagmentSystem/internal/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(username string) (domain.User, error)
	FindUserByID(id primitive.ObjectID) (domain.User, error)
	UpdateUser(id primitive.ObjectID, u domain.User) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	userRepository := &userRepository{
		collection: db.Collection("users"),
	}

	err := userRepository.ensureUniqueIndex()
	if err != nil {
		log.Printf("error creating unique index: %v", err)
	}

	return userRepository
}

func (r *userRepository) ensureUniqueIndex() error {
	_, err := r.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return err
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()

	res, err := r.collection.InsertOne(ctx, user)
	if res.InsertedID == nil || err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (r *userRepository) FindUser(username string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username}

	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user does not exist")
		}
		log.Printf("find user error %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindUserByID(id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user does not exist")
		}
		log.Printf("find user error %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(id primitive.ObjectID, u domain.User) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": u}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("error on update %v", err)
		return errors.New("failed to update user")
	}

	return nil
}
