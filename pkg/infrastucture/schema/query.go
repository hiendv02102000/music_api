package schema

import (
	"backend-food/graph/query"

	"github.com/graphql-go/graphql"
)

var QueryTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"login": query.LoginQuery(OutputTypes),
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	},
})
