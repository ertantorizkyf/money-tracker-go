package dto

type GeneralResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
	Data       any    `json:"data"`
}

func SetGeneralResp(statusCode int, message string, isError bool, data any) *GeneralResp {
	return &GeneralResp{
		StatusCode: statusCode,
		Message:    message,
		IsError:    isError,
		Data:       data,
	}
}
