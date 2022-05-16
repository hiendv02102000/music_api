package dto

// CreateFoodPostRequest struct
type CreateFoodPostRequest struct {
	Title   string   `validate:"required"`
	Content string   `validate:"required"`
	Dish    string   `validate:"required"`
	Tags    []string ``
}
