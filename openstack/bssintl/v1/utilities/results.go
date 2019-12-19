package utilities

import "github.com/gophercloud/gophercloud"



type SendVerificationCodeResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`
}


type SendVerificationCodeResult struct {
	gophercloud.Result
}

func (r SendVerificationCodeResult) Extract() (*SendVerificationCodeResp, error) {
	var res *SendVerificationCodeResp
	err := r.ExtractInto(&res)
	return res, err
}


