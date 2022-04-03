package schema

import (
	"backend-food/graph/output"

	"github.com/graphql-go/graphql"
)

var OutputTypes = map[string]*graphql.Object{
	"LoginOutput": output.LoginOutput(),
}
