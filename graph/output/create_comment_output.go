package output

import "github.com/graphql-go/graphql"

func CreateCommentOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CreateCommentOutput",
			Fields: graphql.Fields{
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"song_id": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
