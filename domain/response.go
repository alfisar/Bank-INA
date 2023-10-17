package domain

type Response struct {
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}
