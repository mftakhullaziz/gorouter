package model

type DefaultResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"is_success"`
	Data       interface{} `json:"data"`
}
