package schema

import (
	"backend-food/graph/output"

	"github.com/graphql-go/graphql"
)

func InitOutputType() map[string]*graphql.Object {
	return map[string]*graphql.Object{
		"LoginOutput": output.LoginOutput(),
	}
}
