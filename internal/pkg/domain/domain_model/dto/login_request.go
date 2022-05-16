package dto

type LoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required,password"`
}
