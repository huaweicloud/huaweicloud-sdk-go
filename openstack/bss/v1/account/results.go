package account

import (
	"github.com/gophercloud/gophercloud"
)

type ResourceDaily struct {

	ErrorCode string `json:"error_code"`

	ErrorMsg string `json:"error_msg"`

	TotalRecord int `json:"totalRecord"`

	Currency string `json:"currency"`

	TotalAmount int `json:"totalAmount"`

	MeasureId int `json:"measureId"`

	DailyRecords []DailyRecord `json:"dailyRecords"`

}

type DailyRecord struct {

	//Account type 1: cloud account 2: partner funding account
	Type string `json:"type"`

	//BpID
	BpId string `json:"bpId"`

	//Name of partner company
	BpName string `json:"bpName"`

	//Cloud Service Type。
	CloudServiceType string `json:"cloudServiceType"`

	//Region Code
	RegionCode string `json:"regionCode"`

	//Resource Type
	ResourceType string `json:"resourceType"`

	//Resource Id
	ResourceId string `json:"resourceId"`

	//Resource Name。
	ResourceName string `json:"resourceName"`

	//Resource tag
	Resourcetag string `json:"resourcetag"`

	//Price Factor Name。
	PriceFactorName string `json:"priceFactorName"`

	//Consume Time
	ConsumeTime string `json:"consumeTime"`

	//Consume Amount
	ConsumeAmount int `json:"consumeAmount"`

	//Offical Amount
	OfficalAmount int `json:"officalAmount"`

	//Debt Amount
	DebtAmount int `json:"debtAmount"`

	//Discount amount
	DisCountAmount int `json:"disCountAmount"`

	//Total consumption of old coupons。
	OldCouponTotalAmount int `json:"oldCouponTotalAmount"`

	//The amount of the unit is 1: yuan 2: angle 3: points.
	MeasureId int `json:"measureId"`

	//EnterpriseProjectId
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	//DeductDetailInfos。
	DeductDetailInfos []DeductDetailInfo `json:"DeductDetailInfos"`

	//Example of bidding
	Spot string `json:"spot"`
}

type DeductDetailInfo struct {

	//Account type。
	BalanceTypeId string `json:"balanceTypeId"`

	//Amount
	Amount int `json:"amount"`

	//The amount of the unit is 1: yuan 2: angle 3: points.
	MeasureId int `json:"measureId"`

}

type commonResult struct {
	gophercloud.Result
}


func (r commonResult) Extract() (ResourceDaily, error) {
	var res ResourceDaily
	err := r.ExtractInto(&res)
	return res, err
}
