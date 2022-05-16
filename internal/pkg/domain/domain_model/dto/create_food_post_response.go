package dto

// CreateFoodPostResponse struct
type CreateFoodPostResponse struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Dish      string   `json:"dish"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}
