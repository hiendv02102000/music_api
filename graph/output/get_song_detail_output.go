package output

import "github.com/graphql-go/graphql"

func GetSongDetailOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetSongDetailOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content_url": &graphql.Field{
					Type: graphql.String,
				},
				"view": &graphql.Field{
					Type: graphql.Int,
				},
				"image_url": &graphql.Field{
					Type: graphql.String,
				},
				"decription": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{
						Name:   "user_detail_song",
						Fields: userSongOutput(),
					}),
				},
				// "comments": &graphql.Field{
				// 	Type: &graphql.List{
				// 		OfType: graphql.NewObject(graphql.ObjectConfig{
				// 			Name: "commentsDetailSongOutput",
				// 			Fields: graphql.Fields{
				// 				"id": &graphql.Field{
				// 					Type: graphql.Int,
				// 				},
				// 				"user": &graphql.Field{
				// 					Type: userSongOutput(),
				// 				},
				// 				"content": &graphql.Field{
				// 					Type: graphql.String,
				// 				},
				// 			},
				// 		}),
				// 	},
				// },
			},
		},
	)
}
func userSongOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "userSongOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"avatar_url": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
