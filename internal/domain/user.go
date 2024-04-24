package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required,unique"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required,unique"`
	Username  string             `json:"username" bson:"username" validate:"required,unique"`
	Password  string             `json:"password" bson:"password" validate:"required,min=6,max=16"`
}
