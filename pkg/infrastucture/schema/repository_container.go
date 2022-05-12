package schema

import (
	"be_soc/internal/pkg/repository"
	"be_soc/pkg/infrastucture/db"
)

func GetContainerRepo(data db.Database) map[string]interface{} {

	return map[string]interface{}{
		"user_repository": repository.NewUserRepository(data),
	}
}
