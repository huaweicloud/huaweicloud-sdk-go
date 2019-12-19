package customercredit

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
)

type QueryCreditResp struct {
	//Error Code
	ErrorCode string `json:"error_code"`

	//Description
	ErrorMsg string `json:"error_msg"`

	//Credit limit
	CreditAmount float64 `json:"creditAmount,omitempty"`

	//Used credit
	UsedAmount *float64 `json:"usedAmount,omitempty"`

	//Unit
	MeasureId *int `json:"measureId,omitempty"`

	//Currency
	Currency string `json:"currency"`

}

type SetCreditResp struct {
	//Error Code
	ErrorCode string `json:"error_code"`

	//Description
	ErrorMsg string `json:"error_msg"`
}

type errorObj struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type internalErr struct {
	ID           string `json:"id"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

type SetCreditError struct {
	BasicError    errorObj      `json:"error"`
	InternalError []internalErr `json:"internalError"`
}

//Error,Implement the Error() interface.
func (e SetCreditError) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return fmt.Sprintf(string(data))
}

//ErrorCode,Error code converted to string type.
func (e SetCreditError) ErrorCode() string {
	return e.BasicError.Code
}

//Message,Return error message.
func (e SetCreditError) Message() string {
	return e.BasicError.Message
}

type QueryCreditResult struct {
	gophercloud.Result
}

type SetCreditResult struct {
	gophercloud.Result
}

func (r QueryCreditResult) Extract() (*QueryCreditResp, error) {
	var res *QueryCreditResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r SetCreditResult) Extract() (*SetCreditResp, error) {
	var res *SetCreditResp
	err := r.ExtractInto(&res)
	return res, err
}
