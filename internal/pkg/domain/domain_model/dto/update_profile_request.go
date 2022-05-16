package dto

// UpdateUserProfileResponse struct
type UpdateUserProfileRequest struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
}
