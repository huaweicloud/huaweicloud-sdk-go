package bill

import "github.com/gophercloud/gophercloud"

type QueryPartnerMonthlyBillsResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//total count
	Count *int `json:"count,omitempty"`

	//Bill list.
	BillSums []BillSumInfo `json:"billSums"`

}

type BillSumInfo struct {
	//Customer Id。
	CustomerId string `json:"customerId"`

	//Bill type
	BillType string `json:"billType"`

	//Cloud service type code.
	CloudServiceTypeCode string `json:"cloudServiceTypeCode"`

	//Billing mode.
	PayMethod string `json:"payMethod"`

	//total amount
	Amount *float64 `json:"amount,omitempty"`

	//Outstanding amount
	DebtAmount *float64 `json:"debtAmount,omitempty"`

	//Write Off Debt Amount
	WriteOffDebtAmount *float64 `json:"writeoffDebtAmount,omitempty"`

	//Discount Amount
	DiscountAmount *float64 `json:"discountAmount,omitempty"`

	//Unit
	MeasureId *int `json:"measureId,omitempty"`

	//Currency
	Currency string `json:"currency"`

	//account detail list
	AccountDetails []BalanceTypeDeductSum `json:"accountDetails"`
}

type BalanceTypeDeductSum struct {
	//account type
	BalanceTypeId string `json:"balanceTypeId"`

	//total amount
	Amount *float64 `json:"amount,omitempty"`

	//bill type
	BillType string `json:"billType"`
}

type QueryMonthlyExpenditureSummaryResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Currency.
	Currency string `json:"currency"`

	//Number of the total records
	TotalCount *int `json:"total_count,omitempty"`

	//Record information
	BillSums []BillSumRecordInfo `json:"bill_sums"`
}

type BillSumRecordInfo struct {
	//Customer ID.
	CustomerId string `json:"customer_id"`

	//Resource type code
	ResourceTypeCode string `json:"resource_type_code"`

	//Cloud service region
	RegionCode string `json:"region_code"`

	//Cloud service type code
	CloudServiceTypeCode string `json:"cloud_service_type_code"`

	//Expenditure data collection period
	ConsumeTime string `json:"consume_time"`

	//Expenditure type
	PayMethod string `json:"pay_method"`

	//Consumption amount
	ConsumeAmount *float64 `json:"consume_amount,omitempty"`

	//Outstanding amount
	Debt *float64 `json:"debt,omitempty"`

	//Discounted amount
	Discount *float64 `json:"discount,omitempty"`

	//Unit
	MeasureId *int `json:"measure_id,omitempty"`

	//Expenditure type
	BillType *int `json:"bill_type,omitempty"`

	//Total payment amount distinguished by expenditure type and payment method of an account.
	AccountDetails []BalanceTypePay `json:"account_details"`

	//Discounted amount details
	DiscountDetailInfos []DiscountDetailInfo `json:"discount_detail_infos"`

	//Enterprise project ID
	EnterpriseProjectId string `json:"enterpriseProjectId"`
}

type BalanceTypePay struct {
	//Account type
	BalanceTypeId string `json:"balance_type_id,omitempty"`

	//Deducted amount
	DeductAmount float64 `json:"deduct_amount,omitempty"`
}

type DiscountDetailInfo struct {
	//Discount type
	PromotionType string `json:"promotion_type"`

	//Discounted amount
	DiscountAmount *float64 `json:"discount_amount,omitempty"`

	//Discount type ID
	PromotionId string `json:"promotion_id"`

	//Unit
	MeasureId *int `json:"measure_id,omitempty"`
}

type QueryResourceUsageDetailsResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Currency unit
	Currency string `json:"currency"`

	//Number of result sets
	TotalCount *int `json:"total_count,omitempty"`

	//Resource usage record
	MonthlyRecords []MonthlyRecord `json:"monthlyRecords"`
}

type MonthlyRecord struct {
	//Cloud service type code
	CloudServiceTypeCode string `json:"cloudServiceTypeCode"`

	//Resource type code
	ResourceTypeCode string `json:"resourceTypeCode"`

	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//Resource instance ID
	ResInstanceId string `json:"resInstanceId"`

	//Resource name
	ResourceName string `json:"resourceName"`

	//Resource tag
	ResourceTag string `json:"resourceTag"`

	//Consumption amount of a cloud service, including the amount of cash coupons.
	ConsumeAmount *float64 `json:"consumeAmount,omitempty"`

	//Expenditure month
	Cycle string `json:"cycle"`

	//Unit
	MeasureId *int `json:"measureId,omitempty"`

	//Enterprise project ID
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	//Billing mode
	PayMethod string `json:"payMethod"`
}

type QueryResourceUsageRecordResp struct {
	//Error code
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Currency unit.
	Currency string `json:"currency"`

	//Number of result sets
	TotalCount *int `json:"totalCount,omitempty"`

	//Resource usage record
	FeeRecords []ResFeeRecord `json:"feeRecords"`
}

type ResFeeRecord struct {
	//Fee generation time
	CreateTime string `json:"createTime"`

	//Start time of using the resource corresponding to the fee.
	EffectiveTime string `json:"effectiveTime"`

	//End time of using the resource corresponding to the fee
	ExpireTime string `json:"expireTime"`

	//Fee record serial number
	FeeId string `json:"feeId"`

	//Product ID
	ProductId string `json:"productId"`

	//Product name
	ProductName string `json:"productName"`

	//Order ID
	OrderId string `json:"orderId"`

	//Consumption amount, including the amount of cash coupons.
	Amount *float64 `json:"amount,omitempty"`

	//Unit
	MeasureId *int `json:"measureId,omitempty"`

	//Usage
	UsageAmount *float64 `json:"usageAmount,omitempty"`

	//Usage unit
	UsageMeasureId *int `json:"usageMeasureId,omitempty"`

	//Package usage.
	FreeResourceAmount *float64 `json:"freeResourceAmount,omitempty"`

	//Unit (package usage)
	FreeResourceMeasureId *int `json:"freeResourceMeasureId,omitempty"`

	//Cloud service type code
	CloudServiceTypeCode string `json:"cloudServiceTypeCode"`

	//Resource type code
	ResourceTypeCode string `json:"resourceTypeCode"`

	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//Payment method
	PayMethod string `json:"payMethod"`

	//Project ID.
	ProjectID string `json:"projectID"`

	//Project name.
	ProjectName string `json:"projectName"`

	//Resource tag
	ResourceTag string `json:"resourceTag"`

	//Resource name
	ResourceName string `json:"resourceName"`

	//Resource ID.
	ResourceId string `json:"resourceId"`

	//Expenditure type
	FeeSourceOperation *int `json:"feeSourceOperation,omitempty"`

	//Enterprise project ID.
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	//Period type
	PeriodType string `json:"periodType"`

	//Spot instance ID
	Spot string `json:"spot"`

	//Reserved instance usage
	RIAmount *float64 `json:"rIAmount,omitempty"`

	//Unit (reserved instance usage)
	RIMeasureId *int `json:"rIMeasureId,omitempty"`
}

type QueryPartnerMonthlyBillsResult struct {
	gophercloud.Result
}

func (r QueryPartnerMonthlyBillsResult) Extract() (*QueryPartnerMonthlyBillsResp, error) {
	var res *QueryPartnerMonthlyBillsResp
	err := r.ExtractInto(&res)
	return res, err
}

type QueryMonthlyExpenditureSummaryResult struct {
	gophercloud.Result
}

func (r QueryMonthlyExpenditureSummaryResult) Extract() (*QueryMonthlyExpenditureSummaryResp, error) {
	var res *QueryMonthlyExpenditureSummaryResp
	err := r.ExtractInto(&res)
	return res, err
}

type QueryResourceUsageDetailsResult struct {
	gophercloud.Result
}

func (r QueryResourceUsageDetailsResult) Extract() (*QueryResourceUsageDetailsResp, error) {
	var res *QueryResourceUsageDetailsResp
	err := r.ExtractInto(&res)
	return res, err
}

type QueryResourceUsageRecordResult struct {
	gophercloud.Result
}

func (r QueryResourceUsageRecordResult) Extract() (*QueryResourceUsageRecordResp, error) {
	var res *QueryResourceUsageRecordResp
	err := r.ExtractInto(&res)
	return res, err
}

