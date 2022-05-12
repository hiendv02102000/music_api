package schema

import (
	"backend-food/internal/pkg/repository"
	"backend-food/pkg/infrastucture/db"
)

func GetContainerRepo(data db.Database) map[string]interface{} {

	return map[string]interface{}{
		"user_repository": repository.NewUserRepository(data),
	}
}
