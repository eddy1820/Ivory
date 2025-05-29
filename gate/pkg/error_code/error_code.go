package error_code

import (
	"fmt"
)

type ErrorData struct {
	HttpStatus int      `json:"httpStatus"`
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Details    []string `json:"details,omitempty"`
}

var codes = map[int]string{}

func NewErrorData(statusCode int, code int, msg string) *ErrorData {
	if _, isMapped := codes[code]; isMapped {
		panic(fmt.Sprintf("錯誤 %d 已經存在, 請更換一個", code))
	}
	codes[code] = msg
	return &ErrorData{HttpStatus: statusCode, Code: code, Message: msg}
}

func (e *ErrorData) Error() string {
	return fmt.Sprintf("錯誤: %d, 錯誤訊息: %s", e.Code, e.Message)
}

func (e *ErrorData) WithDetails(details ...string) *ErrorData {
	newError := *e
	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}

	return &newError
}
