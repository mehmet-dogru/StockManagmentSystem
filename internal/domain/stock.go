package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FormID      primitive.ObjectID `json:"formId" bson:"formId"`
	ProductName string             `json:"productName" bson:"productName"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Price       float64            `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	IsAvailable bool               `json:"isAvailable" bson:"isAvailable"`
}
