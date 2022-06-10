package schema

import (
	"music-api/internal/pkg/repository"
	"music-api/pkg/infrastucture/db"
)

func GetContainerRepo(data db.Database) map[string]interface{} {

	return map[string]interface{}{
		"user_repository":    repository.NewUserRepository(data),
		"song_repository":    repository.NewSongRepository(data),
		"comment_repository": repository.NewCommentRepository(data),
		//"food_post_repository": repository.NewFoodPostRepository(data),
	}
}
