package service

import "backend-food/internal/pkg/domain/domain_model/entity"

type CommentRepository interface {
	CreateComment(entity.Comment) (entity.Comment, error)
}
