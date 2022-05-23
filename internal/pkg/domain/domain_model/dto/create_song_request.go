package dto

// CreateSongRequest struct
type CreateSongRequest struct {
	Title       string `validate:"required"`
	Description string
	Singer      string
}
