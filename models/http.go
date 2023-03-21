package model

type TSuccessResponse struct {
	Data    interface{} `json:"data"`
}
type TErrorResponse struct {
	Error    interface{} `json:"errors"`
}
