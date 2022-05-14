package input

import "github.com/graphql-go/graphql"

func CommentInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "Comment Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"content": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"food_post_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
