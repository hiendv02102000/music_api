package dto

// LoginResponse struct
type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,password"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
