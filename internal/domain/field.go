package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Field struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FormID     primitive.ObjectID `json:"form_id" bson:"form_id"`
	Name       string             `json:"name" bson:"name"`
	Type       string             `json:"type" bson:"type"`
	Options    []string           `json:"options,omitempty" bson:"options"`
	MinChars   int                `json:"min_chars,omitempty" bson:"min_chars"`
	MaxChars   int                `json:"max_chars,omitempty" bson:"max_chars"`
	MinValue   int                `json:"min_value,omitempty" bson:"min_value"`
	MaxValue   int                `json:"max_value,omitempty" bson:"max_value"`
	IsRequired bool               `json:"is_required" bson:"is_required"`
	IsUnique   bool               `json:"is_unique" bson:"is_unique"`
	IsHidden   bool               `json:"is_hidden" bson:"is_hidden"`
	Order      int                `json:"order" bson:"order"`
}
