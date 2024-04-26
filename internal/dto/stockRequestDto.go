package dto

type AddStockRequestDto struct {
	ProductName string  `json:"productName" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	IsAvailable bool    `json:"isAvailable" validate:"required"`
}

type UpdateStockRequestDto struct {
	ProductName string  `json:"productName" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	IsAvailable bool    `json:"isAvailable" validate:"required"`
}
