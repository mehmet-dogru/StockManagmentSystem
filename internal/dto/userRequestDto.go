package dto

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=16"`
}

type UserSignup struct {
	UserLogin
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
}

type ProfileInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"userName"`
}
