package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/utils"
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

func LoginQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.LoginOutput(),
		Description: "User Login",
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.UserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

			req := p.Args["user"].(map[string]interface{})
			loginReq := dto.LoginRequest{
				Username: req["username"].(string),
				Password: req["password"].(string),
			}

			err = utils.CheckValidate(loginReq)
			if err != nil {
				return
			}
			loginReq.Password = utils.EncryptPassword(loginReq.Password)
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			user, err := userRepo.FirstUser(entity.Users{
				Username: loginReq.Username,
				Password: loginReq.Password,
			})
			if err != nil {
				return
			}
			if user.ID == 0 {
				err = errors.New("Login fail")
				return
			}
			timeNow := time.Now()

			timeExpriedAt := timeNow.Add(time.Hour * 2)
			// generate uuid
			uuid := uuid.Must(uuid.NewV4(), nil)
			tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
				UUID:       uuid,
				Authorized: true,
				ExpriedAt:  timeExpriedAt,
			})

			if err != nil {
				return
			}

			newUser := entity.Users{
				Token:          &tokenString,
				TokenExpriedAt: &timeExpriedAt,
			}
			_, err = userRepo.UpdateUser(newUser, user)
			result = map[string]interface{}{
				"token":            newUser.Token,
				"token_expried_at": newUser.TokenExpriedAt,
				"role":             user.Role,
			}

			return
		},
	}
}
