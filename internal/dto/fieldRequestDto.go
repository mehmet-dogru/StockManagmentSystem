package dto

type CreateFieldInput struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Options    []string `json:"options,omitempty"`
	MinChars   int      `json:"minChars,omitempty"`
	MaxChars   int      `json:"maxChars,omitempty"`
	MinValue   int      `json:"minValue,omitempty"`
	MaxValue   int      `json:"maxValue,omitempty"`
	IsRequired bool     `json:"isRequired"`
	IsUnique   bool     `json:"isUnique"`
	IsHidden   bool     `json:"isHidden"`
	Order      int      `json:"order"`
}

type UpdateFieldInput struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Options    []string `json:"options,omitempty"`
	MinChars   int      `json:"minChars,omitempty"`
	MaxChars   int      `json:"maxChars,omitempty"`
	MinValue   int      `json:"minValue,omitempty"`
	MaxValue   int      `json:"maxValue,omitempty"`
	IsRequired bool     `json:"isRequired"`
	IsUnique   bool     `json:"isUnique"`
	IsHidden   bool     `json:"isHidden"`
	Order      int      `json:"order"`
}
