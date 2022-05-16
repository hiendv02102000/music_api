package input

import "github.com/graphql-go/graphql"

func FoodPostInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "Food Post Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"content": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dish": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"user": &graphql.InputObjectFieldConfig{
				Type: UserInput(),
			},
			"tags": &graphql.InputObjectFieldConfig{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
		},
	})
}
