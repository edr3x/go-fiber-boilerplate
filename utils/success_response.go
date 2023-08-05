package utils

type SuccessResponse struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}
