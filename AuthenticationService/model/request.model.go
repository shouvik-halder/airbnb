package model

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
