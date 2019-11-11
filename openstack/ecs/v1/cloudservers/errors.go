package cloudservers

import (
	"encoding/json"
	"fmt"
)

type errorObj struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type internalErr struct {
	ID           string `json:"id"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type BatchOperateError struct {
	BasicError    errorObj      `json:"error"`
	InternalError []internalErr `json:"internalError"`
}

//Error,Implement the Error() interface.
func (e BatchOperateError) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return fmt.Sprintf(string(data))
}

//ErrorCode,Error code converted to string type.
func (e BatchOperateError) ErrorCode() string {
	return e.BasicError.Code
}

//Message,Return error message.
func (e BatchOperateError) Message() string {
	return e.BasicError.Message
}
