package dto

type CreateFormInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateFormInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
