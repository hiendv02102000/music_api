package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"

	"github.com/graphql-go/graphql"
)

func GetCommentSongQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetSongDetailOutput(),
		Description: "get comments of song",
		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

			return
		},
	}
}
