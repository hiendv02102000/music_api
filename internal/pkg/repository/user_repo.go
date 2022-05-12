package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"
)

type userRepository struct {
	DB db.Database
}

func (u *userRepository) FirstUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(condition, &user)
	return user, err
}
func (u *userRepository) FindUserList(condition entity.Users) (user []entity.Users, err error) {
	err = u.DB.Find(condition, &user)
	return
}
func (u *userRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	return user, err
}
func (u *userRepository) DeleteUser(user entity.Users) error {
	err := u.DB.Delete(&user)
	return err
}
func (u *userRepository) UpdateUser(user, oldUser entity.Users) error {
	return u.DB.Update(entity.Users{}, &oldUser, &user)
}
func NewUserRepository(db db.Database) *userRepository {
	return &userRepository{
		DB: db,
	}
}
