package bill

import "github.com/gophercloud/gophercloud"

type QueryPartnerMonthlyBillsResp struct {
	//Error code.
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Billing cycle
	BillCycle string `json:"billCycle"`

	//Bill amount, which is calculated based on the special commercial discount of the partner.
	CreditDebtAmount *float64 `json:"creditDebtAmount,omitempty"`

	//Consumption amount, which is calculated based on the special commercial discount of the partner.
	ConsumeAmount *float64 `json:"consumeAmount,omitempty"`

	//Write-off amount (negative value), which is calculated based on the special commercial discount of the partner.
	Writeoffdebt *float64 `json:"writeoffdebt,omitempty"`

	//Unsubscription amount (negative value), which is calculated based on the special commercial discount of the partner.
	unsubscribeAmount *float64 `json:"unsubscribeAmount,omitempty"`

	//Unit
	measureId *int `json:"measureId,omitempty"`

	//This parameter is returned only when the query is successful.
	Currency string `json:"currency"`

	//Tax amount, which is the tax amount in the creditDebtAmount field.
	TaxAmount *float64 `json:"taxAmount,omitempty"`

	//Bill amount that is not settled, which is calculated based on the special commercial discount of the partner.
	UnclearedAmount *float64 `json:"unclearedAmount,omitempty"`

	//Due date for bills.
	DueDate string `json:"dueDate"`

	//Bill list.
	BillList []PostpaidBillInfo `json:"billList"`

}

type PostpaidBillInfo struct {
	//Bill type
	BillType string `json:"billType"`

	//Cloud service type code.
	CloudServiceTypeCode string `json:"cloudServiceTypeCode"`

	//Resource type code.
	ResourceTypeCode string `json:"resourceTypeCode"`

	//Billing mode.
	PayMethod string `json:"payMethod"`

	//Transaction amount/unsubscription amount/refund amount of the customer, including the vouchers, flexi-purchase coupons, reserved flexi-purchase coupons, and stored-value cards.
	CreditDebtAmount *float64 `json:"creditDebtAmount"`

	//Transaction amount/unsubscription amount/refund amount of the customer,not including the vouchers, flexi-purchase coupons, reserved flexi-purchase coupons, or stored-value cards.
	CustomerAmountDue *float64 `json:"customerAmountDue"`

	//Settlement product type.
	SettlementType *int `json:"settlementType,omitempty"`

	//Partner discount percentage
	PartnerRatio *float64 `json:"partnerRatio,omitempty"`

	//Amount that the partner needs to refund/Amount that the partner has refund
	PartnerAmount *float64 `json:"partnerAmount,omitempty"`

	//Yearly/monthly unit.
	PeriodType *int `json:"periodType,omitempty"`

	//Number of yearly/month periods.
	PeriodNum *int `json:"periodNum,omitempty"`

	//Product category code.
	CategoryCode string `json:"categoryCode"`
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

	//Cloud service region code
	ResourceTypeCode string `json:"resourceTypeCode"`

	//Resource type code
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

