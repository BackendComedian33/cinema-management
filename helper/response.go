package helper

import "technical-test/dto"

func BuildResponse(status_code int, success bool, data interface{}, err error, errorCode int) dto.ApiResponse {

	return dto.ApiResponse{
		StatusCode: status_code,
		Success:    success,
		Data:       data,
		Error:      err,
		ErrorCode:  errorCode,
	}
}
