package models

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
