package dto

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,max=50"`
	Email    string `json:"email" validate:"omitempty,email,max=50"`
	Password string `json:"passoword" validate:"omitempty,min=6,max=100"`
}
