package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"
)

type foodPostRepository struct {
	DB db.Database
}

func (r *foodPostRepository) CreateFoodPost(foodPost entity.FoodPost) (entity.FoodPost, error) {
	err := r.DB.Create(&foodPost)
	return foodPost, err
}
func NewFoodPostRepository(db db.Database) *foodPostRepository {
	return &foodPostRepository{
		DB: db,
	}
}
