package input

import "github.com/graphql-go/graphql"

func FoodPostInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "Food Post Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"content": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"first_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dish": &graphql.InputObjectFieldConfig{
				Type: DishInput(),
			},
			"user": &graphql.InputObjectFieldConfig{
				Type: UserInput(),
			},
			"tag": &graphql.InputObjectFieldConfig{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
		},
	})
}
