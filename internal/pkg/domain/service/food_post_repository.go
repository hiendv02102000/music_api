package service

import "backend-food/internal/pkg/domain/domain_model/entity"

type FoodPostRepositoryInterface interface {
	CreateFoodPost(entity.FoodPost) (entity.FoodPost, error)
}
