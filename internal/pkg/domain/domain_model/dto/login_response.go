package dto

// LoginResponse struct
type LoginResponse struct {
	Token          string `json:"token"`
	TokenExpiredAt string `json:"token_expired_at"`
	Role           string `json:"role"`
}
