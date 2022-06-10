package mutation

import (
	"errors"
	"music-api/graph/input"
	"music-api/graph/output"
	"music-api/internal/pkg/domain/domain_model/dto"
	"music-api/internal/pkg/domain/domain_model/entity"
	"music-api/internal/pkg/domain/service"
	"music-api/pkg/share/middleware"
	"music-api/pkg/share/utils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func UpdateSongMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.UpdateSongOutput(),
		Description: "update a song",

		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["song"] == nil {
				err = errors.New("song is required")
				return
			}
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["song"].(map[string]interface{})
			updateSongReq := dto.UpdateSongRequest{
				SongID:      req["id"].(int),
				Title:       req["title"].(string),
				Description: req["decription"].(string),
				Singer:      req["singer"].(string),
			}
			err = utils.CheckValidate(updateSongReq)
			if err != nil {
				return
			}
			songRepo := containerRepo["song_repository"].(service.SongRepository)
			newSong := entity.Song{
				Title:      updateSongReq.Title,
				Decription: updateSongReq.Description,
				Singer:     updateSongReq.Singer,
			}
			imgFile, _ := ctx.FormFile("image_file")
			if imgFile != nil {
				ioFile, errFile := imgFile.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, imgFile.Filename, []string{"img_music"})
				if errUpload != nil {
					err = errUpload
					return
				}
				newSong.ImageURL = &url
			}

			contentFile, _ := ctx.FormFile("music_file")
			if contentFile != nil {
				ioFile, errFile := contentFile.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, contentFile.Filename, []string{"music"})
				if errUpload != nil {
					err = errUpload
					return
				}
				newSong.ContentURL = &url
			}

			song, err := songRepo.UpdateSong(newSong, entity.Song{
				ID:     updateSongReq.SongID,
				UserID: user.ID,
			})
			if err != nil {
				return
			}
			song, err = songRepo.FirstSong(entity.Song{
				ID:     updateSongReq.SongID,
				UserID: user.ID,
			})
			if err != nil {
				return
			}
			result = dto.UpdateSongResponse{
				ID:          song.ID,
				Title:       song.Title,
				ContentURL:  song.ContentURL,
				ImageURL:    song.ImageURL,
				Description: song.Decription,
				Singer:      song.Singer,
				UpdatedAt:   utils.FormatTime(song.UpdatedAt),
			}
			return
		},
	}
}
