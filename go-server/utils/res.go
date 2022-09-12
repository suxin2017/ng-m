package utils

var (
	OkCode    = 1
	ErrorCode = -1
)

type ErrorRes struct {
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type OkRes struct {
	Code int `json:"code,omitempty"`
	Data any `json:"data,omitempty"`
}
type OkMessageRes struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func Ok(data any) OkRes {
	return OkRes{
		Code: OkCode,
		Data: data,
	}

}
func OkMessage() OkMessageRes {
	return OkMessageRes{
		Code:    OkCode,
		Message: "操作成功",
	}

}
func Error(message string) ErrorRes {
	return ErrorRes{
		Code:    ErrorCode,
		Message: message,
	}
}
func ErrorWithTrack(message string, err error) ErrorRes {
	return ErrorRes{
		Code:         OkCode,
		Message:      message,
		ErrorMessage: err.Error(),
	}
}
