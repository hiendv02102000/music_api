package mutation

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateFoodPostMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateFoodPostOutput(),
		Description: "Create Food Post",

		Args: graphql.FieldConfigArgument{
			"food_post": &graphql.ArgumentConfig{
				Type: input.FoodPostInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["food_post"] == nil {
				err = errors.New("food_post is required")
				return
			}
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["food_post"].(map[string]interface{})
			tags, err := utils.ConverInterfaceToString(req["tags"].([]interface{}))
			if err != nil {
				err = errors.New("tags must be a string array")
				return
			}
			createFoodPostRequest := dto.CreateFoodPostRequest{
				Title:   req["title"].(string),
				Content: req["content"].(string),
				Dish:    req["dish"].(string),
				Tags:    tags,
			}
			err = utils.CheckValidate(createFoodPostRequest)
			if err != nil {
				return
			}
			foodPostRepo := containerRepo["food_post_repository"].(service.FoodPostRepositoryInterface)
			postTags := ""
			for _, tag := range createFoodPostRequest.Tags {
				postTags += utils.FormatStringSpace(tag) + ","
			}
			foodPost := entity.FoodPost{
				Title:   createFoodPostRequest.Title,
				Content: createFoodPostRequest.Content,
				Dish:    createFoodPostRequest.Dish,
				Tags:    postTags,
				UserID:  user.ID,
			}
			file, _ := ctx.FormFile("file")
			if file != nil {
				ioFile, errFile := file.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, file.Filename, []string{"food_post"})
				if errUpload != nil {
					err = errUpload
					return
				}
				foodPost.ImageURL = &url
			}
			foodPost, err = foodPostRepo.CreateFoodPost(foodPost)
			if err != nil {
				return
			}
			result = dto.CreateFoodPostResponse{
				Title:     foodPost.Title,
				Content:   foodPost.Content,
				Dish:      foodPost.Dish,
				Tags:      createFoodPostRequest.Tags,
				CreatedAt: utils.FormatTime(foodPost.CreatedAt),
			}
			return
		},
	}
}
