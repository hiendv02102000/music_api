package dto

type BaseRequest struct {
	Query string `json:"query" form:"query"`
}
