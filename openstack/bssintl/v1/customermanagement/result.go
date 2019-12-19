package customermanagement

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
)

type CheckCustomerRegisterInfoResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Status
	Status string `json:"status"`

	//Whether the number of verification code sending times reaches the upper limit (15 times per hour, 60 times per day)
	UpLimit string `json:"uplimit"`
}

type CreateCustomerResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Customer ID
	DomainId string `json:"domainId"`

	//HUAWEI CLOUD username of the customer
	DomainName string `json:"domainName"`
}

func (e CreateCustomerResp) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return fmt.Sprintf(string(data))
}

type QueryCustomerResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Customer list
	CustomerInfoList []customerInfoList `json:"customerInfoList"`

	//Total number of records
	Count *int `json:"count,omitempty"`
}

type customerInfoList struct {
	//Name that has passed the real-name authentication.
	Name string `json:"name"`

	//Account name.
	DomainName string `json:"domainName"`

	//Customer ID.
	CustomerId string `json:"customerId"`

	//Time when a customer is associated with a partner.
	CooperationTime string `json:"cooperationTime"`

	//Association type
	CooperationType string `json:"cooperationType"`

	//Tag
	Label string `json:"label"`

	//Customer phone number
	Telephone string `json:"telephone"`

	//Real-name authentication status
	VerifiedStatus string `json:"verifiedStatus"`

	//Country code, which is the country code prefix of a phone number.
	CountryCode string `json:"countryCode"`

	//Customer type
	CustomerType *int `json:"customerType,omitempty"`

	//Whether to freeze the account.
	IsFrozen *int `json:"isFrozen,omitempty"`
}

type FrozenCustomerResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error message
	ErrorMsg string `json:"error_msg"`

	//Error cause description
	FailDetail []ErrorDetail `json:"failDetail"`

	//Number of failures
	FailNum *int `json:"failNum,omitempty"`

	//Number of successful operations
	SuccessNum *int `json:"successNum,omitempty"`
}

type ErrorDetail struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error message
	ErrorMsg string `json:"error_msg"`

	//The value corresponds to customerId
	Id string `json:"Id"`

}

type UnFrozenCustomerResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error message
	ErrorMsg string `json:"error_msg"`

	//Error cause description
	FailDetail []ErrorDetail `json:"failDetail"`

	//Number of failures
	FailNum *int `json:"failNum,omitempty"`

	//Number of successful operations
	SuccessNum *int `json:"successNum,omitempty"`
}

type CheckCustomerRegisterInfoResult struct {
	gophercloud.Result
}

func (r CheckCustomerRegisterInfoResult) Extract() (*CheckCustomerRegisterInfoResp, error) {
	var res *CheckCustomerRegisterInfoResp
	err := r.ExtractInto(&res)
	return res, err
}

type CreateCustomerResult struct {
	gophercloud.Result
}

func (r CreateCustomerResult) Extract() (*CreateCustomerResp, error) {
	var res *CreateCustomerResp
	err := r.ExtractInto(&res)
	return res, err
}

type QueryCustomerResult struct {
	gophercloud.Result
}

func (r QueryCustomerResult) Extract() (*QueryCustomerResp, error) {
	var res *QueryCustomerResp
	err := r.ExtractInto(&res)
	return res, err
}

type FrozenCustomerResult struct {
	gophercloud.Result
}

func (r FrozenCustomerResult) Extract() (*FrozenCustomerResp, error) {
	var res *FrozenCustomerResp
	err := r.ExtractInto(&res)
	return res, err
}

type UnFrozenCustomerResult struct {
	gophercloud.Result
}

func (r UnFrozenCustomerResult) Extract() (*UnFrozenCustomerResp, error) {
	var res *UnFrozenCustomerResp
	err := r.ExtractInto(&res)
	return res, err
}