package query

import (
	"backend-food/internal/pkg/domain/domain_model/entity"

	"github.com/graphql-go/graphql"
)

func LoginQuery(outputType map[string]*graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        outputType["LoginOutput"],
		Description: "User Login",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token := "12356"
			return entity.Users{Token: &token}, nil
		},
	}
}
