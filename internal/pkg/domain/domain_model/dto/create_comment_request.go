package dto

// CreateFoodPostRequest struct
type CreateCommentRequest struct {
	SongID  int    `validate:"required"`
	Content string `validate:"required"`
}
