package handler

import (
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/repository"
	"backend-food/internal/pkg/usecase"
	"backend-food/pkg/infrastucture/db"
	"backend-food/pkg/infrastucture/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type HTTPHandler struct {
	Schema  *graphql.Schema
	usecase *usecase.UserUsecase
}

func NewHTTPHandler(db db.Database) *HTTPHandler {
	usersRepository := repository.NewUserRepository(db)
	usersUsecase := usecase.NewUserUsecase(usersRepository)
	schema := schema.NewSchema()
	return &HTTPHandler{usecase: usersUsecase, Schema: schema}
}

func (h *HTTPHandler) Handle(c *gin.Context) {
	req := dto.BaseRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	fmt.Println(req.Query)
	data := graphql.Do(graphql.Params{
		Schema:        *h.Schema,
		RequestString: req.Query,
	})
	c.JSON(http.StatusOK, data)
}
