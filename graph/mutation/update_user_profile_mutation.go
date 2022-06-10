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

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func UpdateUserProfileMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.UpdateUserProfileOutput(),
		Description: "User Update",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.UserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["user"] == nil {
				err = errors.New("user is required")
				return
			}
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["user"].(map[string]interface{})
			updateUserProfileReq := dto.UpdateUserProfileRequest{
				FirstName: req["first_name"].(string),
				LastName:  req["last_name"].(string),
			}
			err = utils.CheckValidate(updateUserProfileReq)
			if err != nil {
				return
			}
			userRepo := containerRepo["user_repository"].(service.UserRepository)
			newUser := entity.Users{FirstName: updateUserProfileReq.FirstName, LastName: updateUserProfileReq.LastName}
			file, _ := ctx.FormFile("file")
			if file != nil {
				ioFile, errFile := file.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, file.Filename, []string{"avt_user"})
				if errUpload != nil {
					err = errUpload
					return
				}
				newUser.AvatarUrl = &url
			}
			user, err = userRepo.UpdateUser(newUser, user)
			result = dto.UpdateUserProfileResponse{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				AvatarURL: user.AvatarUrl,
			}
			return
		},
	}
}
