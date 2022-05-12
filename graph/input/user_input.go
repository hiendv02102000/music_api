package input

import "github.com/graphql-go/graphql"

func UserInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateUserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"username": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"first_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
