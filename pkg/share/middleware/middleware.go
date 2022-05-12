package middleware

import (
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/repository"
	"be_soc/pkg/infrastucture/db"
	"net/http"
	"strings"
	"time"

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

		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Authorization Token is required",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		clientToken = strings.TrimSpace(extractedToken[1])
		repo := repository.UserRepository{
			DB: db,
		}
		user, err := repo.FirstUser(entity.Users{
			Token: &clientToken,
		})

		if err != nil {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  err.Error(),
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		if user.ID == 0 {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Token",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		timeNow := time.Now()
		if timeNow.After(*user.TokenExpriedAt) {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Token Expired",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func AuthAdminMiddleware(db db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Authorization Token is required",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		clientToken = strings.TrimSpace(extractedToken[1])
		repo := repository.UserRepository{
			DB: db,
		}
		user, err := repo.FirstUser(entity.Users{
			Token: &clientToken,
		})

		if err != nil {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  err.Error(),
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		if user.ID == 0 {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Token",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		timeNow := time.Now()
		if timeNow.After(*user.TokenExpriedAt) {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Token Expired",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		if user.Role != entity.AdminRole {
			data := dto.BaseResponse{
				Status: http.StatusForbidden,
				Error:  "Required Admin Authorization",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
