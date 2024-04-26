package dto

type CreateFieldInput struct {
	Name       string   `json:"name" validate:"required"`
	Type       string   `json:"type" validate:"required"`
	Options    []string `json:"options,omitempty" validate:"omitempty,max=10"`
	MinChars   int      `json:"minChars,omitempty" validate:"omitempty,min=0"`
	MaxChars   int      `json:"maxChars,omitempty" validate:"omitempty,min=0"`
	MinValue   int      `json:"minValue,omitempty" validate:"omitempty,min=0"`
	MaxValue   int      `json:"maxValue,omitempty" validate:"omitempty,min=0"`
	IsRequired bool     `json:"isRequired"`
	IsUnique   *bool    `json:"isUnique"`
	IsHidden   *bool    `json:"isHidden"`
	Order      int      `json:"order" validate:"required"`
}

type UpdateFieldInput struct {
	Name       string   `json:"name" validate:"required"`
	Type       string   `json:"type" validate:"required"`
	Options    []string `json:"options,omitempty" validate:"omitempty,max=10"`
	MinChars   int      `json:"minChars,omitempty" validate:"omitempty,min=0"`
	MaxChars   int      `json:"maxChars,omitempty" validate:"omitempty,min=0"`
	MinValue   int      `json:"minValue,omitempty" validate:"omitempty,min=0"`
	MaxValue   int      `json:"maxValue,omitempty" validate:"omitempty,min=0"`
	IsRequired bool     `json:"isRequired"`
	IsUnique   bool     `json:"isUnique" validate:"required"`
	IsHidden   bool     `json:"isHidden" validate:"required"`
	Order      int      `json:"order" validate:"required"`
}
