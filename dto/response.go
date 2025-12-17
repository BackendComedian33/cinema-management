package dto

type ApiResponse struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      error       `json:"error,omitempty"`
	ErrorCode  int         `json:"error_code,omitempty"`
}
