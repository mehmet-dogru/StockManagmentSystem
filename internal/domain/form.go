package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Form struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	Title       string             `json:"title" bson:"title" validate:"required,unique"`
	Description string             `json:"description" bson:"description" validate:"required,unique"`
}
