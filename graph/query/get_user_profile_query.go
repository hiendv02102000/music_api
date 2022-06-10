package query

import (
	"music-api/graph/input"
	"music-api/graph/output"
	"music-api/internal/pkg/domain/domain_model/dto"
	"music-api/internal/pkg/domain/domain_model/entity"
	"music-api/internal/pkg/domain/service"
	"music-api/pkg/share/middleware"
	"music-api/pkg/share/utils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GetProfileUserQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetUserProfileOutput(),
		Description: "Get User Profile",
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.UserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			user := entity.Users{}
			ctx := p.Context.(*gin.Context)

			if p.Args["user"] != nil {
				req := p.Args["user"].(map[string]interface{})
				if req["id"] != nil {
					user.ID = req["id"].(int)
					userRepo := containerRepo["user_repository"].(service.UserRepository)
					user, err = userRepo.FirstUser(user)
					if err != nil {
						return
					}
				} else {
					user = middleware.GetUserFromContext(ctx)
				}
			} else {
				user = middleware.GetUserFromContext(ctx)
			}
			result = dto.GetUserProfileResponse{
				ID:        user.ID,
				Username:  user.Username,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Role:      string(user.Role),
				CreatedAt: utils.FormatTime(user.CreatedAt),
				AvatarURL: user.AvatarUrl,
			}
			return
		},
	}
}
