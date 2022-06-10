package mutation

import (
	"errors"
	"music-api/graph/input"
	"music-api/graph/output"
	"music-api/internal/pkg/domain/domain_model/dto"
	"music-api/internal/pkg/domain/domain_model/entity"
	"music-api/internal/pkg/domain/service"
	"music-api/pkg/share/middleware"
	"music-api/pkg/share/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

func ChangePasswordMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.ChangePasswordOutput(),
		Description: "Change pasword",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.ChangePasswordInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["user"].(map[string]interface{})
			changePasswordReq := dto.ChangePasswordRequest{
				OldPassword: req["old_password"].(string),
				Password:    req["password"].(string),
			}
			err = utils.CheckValidate(changePasswordReq)
			if err != nil {
				return
			}
			if utils.EncryptPassword(changePasswordReq.OldPassword) != user.Password {
				err = errors.New("wrong password")
				return
			}

			userRepo := containerRepo["user_repository"].(service.UserRepository)
			changePasswordReq.Password = utils.EncryptPassword(changePasswordReq.Password)

			timeNow := time.Now()
			timeExpiredAt := timeNow.Add(time.Hour * 2)
			// generate uuid
			uuid := uuid.Must(uuid.NewV4(), nil)
			tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
				UUID:       uuid,
				Authorized: true,
				ExpriedAt:  timeExpiredAt,
			})

			if err != nil {
				return
			}

			newUser := entity.Users{
				Token:          &tokenString,
				TokenExpiredAt: &timeExpiredAt,
				Password:       changePasswordReq.Password,
			}
			_, err = userRepo.UpdateUser(newUser, user)
			loginRes := dto.LoginResponse{
				Token:          tokenString,
				TokenExpiredAt: timeExpiredAt,
			}
			result = loginRes

			return
		},
	}
}
