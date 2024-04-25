package dto

type CreateFieldInput struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Options    []string `json:"options,omitempty"`
	MinChars   int      `json:"min_chars,omitempty"`
	MaxChars   int      `json:"max_chars,omitempty"`
	MinValue   int      `json:"min_value,omitempty"`
	MaxValue   int      `json:"max_value,omitempty"`
	IsRequired bool     `json:"is_required"`
	IsUnique   bool     `json:"is_unique"`
	IsHidden   bool     `json:"is_hidden"`
	Order      int      `json:"order"`
}

type UpdateFieldInput struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Options    []string `json:"options,omitempty"`
	MinChars   int      `json:"min_chars,omitempty"`
	MaxChars   int      `json:"max_chars,omitempty"`
	MinValue   int      `json:"min_value,omitempty"`
	MaxValue   int      `json:"max_value,omitempty"`
	IsRequired bool     `json:"is_required"`
	IsUnique   bool     `json:"is_unique"`
	IsHidden   bool     `json:"is_hidden"`
	Order      int      `json:"order"`
}
