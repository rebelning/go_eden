package model

type Response struct {
	ResCode int         `json:"resCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
