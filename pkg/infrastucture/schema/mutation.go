package schema

import (
	"backend-food/graph/mutation"

	"github.com/graphql-go/graphql"
)

func GetAnonymousMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"create_user": mutation.CreateUserMutation(containerRepo),
		},
	})
}

func GetClientMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"create_food": mutation.CreateFoodPostMutation(containerRepo),
		},
	})
}
func GetAdminMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "mutation",
		Fields: graphql.Fields{},
	})
}
