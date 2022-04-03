package schema

import (
	"github.com/graphql-go/graphql"
)

func NewSchema() *graphql.Schema {
	myschema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: QueryTypes,
		},
	)
	return &myschema
}
