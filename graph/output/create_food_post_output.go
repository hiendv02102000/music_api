package output

import "github.com/graphql-go/graphql"

func CreateFoodPostOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CreateFoodPostOutput", // string not space
			Fields: graphql.Fields{
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"dish": &graphql.Field{
					Type: graphql.String,
				},
				"tags": &graphql.Field{
					Type: &graphql.List{
						OfType: graphql.String,
					},
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
