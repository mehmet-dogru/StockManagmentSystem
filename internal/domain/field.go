package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Field struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FormID     primitive.ObjectID `json:"formId" bson:"formId"`
	Name       string             `json:"name" bson:"name"`
	Type       string             `json:"type" bson:"type"`
	Options    []string           `json:"options,omitempty" bson:"options"`
	MinChars   int                `json:"minChars,omitempty" bson:"minChars"`
	MaxChars   int                `json:"maxChars,omitempty" bson:"maxChars"`
	MinValue   int                `json:"minValue,omitempty" bson:"minValue"`
	MaxValue   int                `json:"maxValue,omitempty" bson:"maxValue"`
	IsRequired bool               `json:"isRequired" bson:"isRequired"`
	IsUnique   bool               `json:"isUnique" bson:"isUnique"`
	IsHidden   bool               `json:"isHidden" bson:"isHidden"`
	Order      int                `json:"order" bson:"order"`
}
