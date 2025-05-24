package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
