package dto

// UpdateSongRequest struct
type UpdateSongRequest struct {
	SongID      int `validate:"required"`
	Title       string
	Description string
	Singer      string
}
