package dto

type ChangePasswordRequest struct {
	Password    string `validate:"required,password"`
	OldPassword string `validate:"required,password"`
}
