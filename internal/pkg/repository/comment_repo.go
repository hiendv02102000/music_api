package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"
)

type commentRepository struct {
	DB db.Database
}

func (u *commentRepository) FindCommentList(condition entity.Comment) (comments []entity.Comment, err error) {
	err = u.DB.Find(&comments, condition)
	return
}
func (u *commentRepository) FindCommentListWithPagination(condition entity.Comment, pageNum int, pageSize int) (comments []entity.Comment, err error) {
	err = u.DB.FindWithPagination(&comments, (pageNum-1)*pageSize, pageSize, condition)
	return
}
func (u *commentRepository) CreateComment(comment entity.Comment) (entity.Comment, error) {
	err := u.DB.Create(&comment)
	return comment, err
}
func NewCommentRepository(db db.Database) *commentRepository {
	return &commentRepository{
		DB: db,
	}
}
