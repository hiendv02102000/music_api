package middleware

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) entity.Users {

	value, exist := c.Get("user")
	if !exist {
		return entity.Users{}
	}
	return value.(entity.Users)
}

func AuthClientMiddleware(db db.Database) gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}

func AuthAdminMiddleware(db db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
