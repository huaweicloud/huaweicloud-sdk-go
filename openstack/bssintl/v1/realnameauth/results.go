package realnameauth

import "github.com/gophercloud/gophercloud"

type IndividualRealNameAuthResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Whether to transfer to manual review
	IsReview int `json:"isReview,omitempty"`

	//Error list.
	ErrorItems []string `json:"errorItems"`
}

type EnterpriseRealNameAuthResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Whether to transfer to manual review
	IsReview int `json:"isReview,omitempty"`

	//Error list.
	ErrorItems []string `json:"errorItems"`
}

type ChangeEnterpriseRealNameAuthResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Whether to transfer to manual review
	IsReview int `json:"isReview,omitempty"`

	//Error list.
	ErrorItems []string `json:"errorItems"`
}

type QueryRealNameAuthResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Real-name authentication review result.
	ReviewResult int `json:"reviewResult,omitempty"`

	//Review comment.
	Opinion string `json:"opinion"`
}

type IndividualRealNameAuthResult struct {
	gophercloud.Result
}

type EnterpriseRealNameAuthResult struct {
	gophercloud.Result
}

type ChangeEnterpriseRealNameAuthResult struct {
	gophercloud.Result
}

type QueryRealNameAuthResult struct {
	gophercloud.Result
}

func (r IndividualRealNameAuthResult) Extract() (*IndividualRealNameAuthResp, error) {
	var res *IndividualRealNameAuthResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r EnterpriseRealNameAuthResult) Extract() (*EnterpriseRealNameAuthResp, error) {
	var res *EnterpriseRealNameAuthResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r ChangeEnterpriseRealNameAuthResult) Extract() (*ChangeEnterpriseRealNameAuthResp, error) {
	var res *ChangeEnterpriseRealNameAuthResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r QueryRealNameAuthResult) Extract() (*QueryRealNameAuthResp, error) {
	var res *QueryRealNameAuthResp
	err := r.ExtractInto(&res)
	return res, err
}
