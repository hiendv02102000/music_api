package handler

import (
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/pkg/infrastucture/db"
	"backend-food/pkg/infrastucture/schema"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type HTTPAdminHandler struct {
	Schema *graphql.Schema
}

func NewHTTPAdminHandler(db db.Database) *HTTPAdminHandler {

	schema := schema.NewAdminSchema(db)
	return &HTTPAdminHandler{Schema: schema}
}

func (h *HTTPAdminHandler) Handle(c *gin.Context) {
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
	//fmt.Println(req.Query)
	//fmt.Println(req.Mutation)
	exce := ""
	if len(req.Query) > 0 {
		exce = req.Query
	} else {
		exce = req.Mutation
	}
	//fmt.Println(exce)
	data := graphql.Do(graphql.Params{
		Context:       c,
		Schema:        *h.Schema,
		RequestString: exce,
	})
	code := http.StatusOK
	if len(data.Errors) > 0 {
		code = http.StatusBadRequest
	}
	resp := dto.BaseResponse{
		Status: code,
		Error:  data.Errors,
		Result: data.Data,
	}
	c.JSON(code, resp)
}
