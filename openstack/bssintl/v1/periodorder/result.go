package periodorder

import (
	"github.com/gophercloud/gophercloud"
)

type QueryOrderListResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Order list.
	Data Result `json:"data"`
}

type Result struct {
	//Number of records per page.
	PageSize *int `json:"pageSize,omitempty"`

	//Current page number.
	PageIndex *int `json:"pageIndex,omitempty"`

	//Number of records that match the query conditions.
	TotalSize *int `json:"totalSize,omitempty"`

	//Order details
	OrderInfos []CustomerOrder `json:"orderInfos"`
}

type CustomerOrder struct {
	//Order ID
	OrderId string `json:"orderId"`

	//Parent order ID (order ID before splitting)
	BaseOrderId string `json:"baseOrderId"`

	//operation entity ID
	BeId string `json:"beId"`

	//Customer ID.
	CustomerId string `json:"customerId"`

	//Customer order source type
	SourceType *int `json:"sourceType,omitempty"`

	//Order status
	Status *int `json:"status,omitempty"`

	//Order type
	OrderType *int `json:"orderType,omitempty"`

	//Order amount
	Currency *float64 `json:"currency,omitempty"`

	//Order amount after the discount
	CurrencyAfterDiscount *float64 `json:"currencyAfterDiscount,omitempty"`

	//Order amount unit
	MeasureId *int `json:"measureId,omitempty"`

	//Amount unit name
	MeasureName string `json:"measureName"`

	//Creation time.
	CreateTime string `json:"createTime"`

	//Payment time.
	PaymentTime string `json:"paymentTime"`

	//Last status update time
	LastUpdateTime string `json:"lastUpdateTime"`

	//Requiring approval or not.
	NeedAudit *bool `json:"needAudit,omitempty"`

	//Currency code.
	CurrencyType string `json:"currencyType"`

	//Contract ID.
	ContractId string `json:"contractId"`

	//Order amount (list price).
	CurrencyOfficial *float64 `json:"currencyOfficial,omitempty"`

	//Order details
	AmountInfo AmountInfo `json:"amountInfo"`

	//Cloud service type code.
	ServiceType string `json:"serviceType"`
}

type AmountInfo struct {
	//Item
	DiscountList []DiscountItem `json:"discountList"`

	//Flexi-purchase coupon amount
	CashcouponAmount *float64 `json:"cashcouponAmount,omitempty"`

	//Cash coupon amount.
	CouponAmount *float64 `json:"couponAmount,omitempty"`

	//Stored-value card amount
	CardAmount *float64 `json:"cardAmount,omitempty"`

	//Handling fee (only for unsubscription orders).
	CommissionAmount *float64 `json:"commissionAmount,omitempty"`

	//Consumptions (only for unsubscription orders).
	ConsumedAmount *float64 `json:"consumedAmount,omitempty"`
}

type DiscountItem struct {
	//Discount type
	DiscountType string `json:"discountType"`

	//Discounted amount.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
}

type QueryOrderDetailResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Order details
	OrderInfo CustomerOrderEntity `json:"orderInfo"`

	//Order item ID array.
	Count *int `json:"count,omitempty"`

	//ID of the primary order item mapping the order item.
	OrderlineItems []OrderLineItemEntity `json:"orderlineItems"`
}

type CustomerOrderEntity struct {
	//Order ID.
	OrderId string `json:"orderId"`

	//Order ID.
	BaseOrderId string `json:"baseOrderId"`

	//operation entity ID
	BeId string `json:"beId"`

	//Customer ID.
	CustomerId string `json:"customerId"`

	//Customer order source type.
	SourceType *int `json:"sourceType,omitempty"`

	//Order status
	Status *int `json:"status,omitempty"`

	//Order type
	OrderType *int `json:"orderType,omitempty"`

	//Order amount
	Currency *float64 `json:"currency,omitempty"`

	//Order amount after the discount
	CurrencyAfterDiscount *float64 `json:"currencyAfterDiscount,omitempty"`

	//Order amount unit.
	MeasureId *int `json:"measureId,omitempty"`

	//Amount unit name.
	MeasureName string `json:"measureName"`

	//Creation time.
	CreateTime string `json:"createTime"`

	//Payment time.
	PaymentTime string `json:"paymentTime"`

	//Last status update time.
	LastUpdateTime string `json:"lastUpdateTime"`

	//Requiring approval or not.
	NeedAudit *bool `json:"needAudit,omitempty"`

	//Order amount (list price).
	CurrencyOfficial *float64 `json:"currencyOfficial,omitempty"`

	//Order details
	AmountInfo AmountInfo `json:"amountInfo"`

	//Currency code.
	CurrencyType string `json:"currencyType"`

	//Contract ID.
	ContractId string `json:"contractId"`

	//Cloud service type code
	ServiceType string `json:"serviceType"`
}

type OrderLineItemEntity struct {
	//Order ID.
	OrderLineItemId string `json:"orderLineItemId"`

	//Cloud service type code
	CloudServiceType string `json:"cloudServiceType"`

	//Product ID.
	ProductId string `json:"productId"`

	//Product specification description
	ProductSpecDesc string `json:"productSpecDesc"`

	//Period type
	PeriodType *int `json:"periodType,omitempty"`

	//Number of periods.
	PeriodNum *int `json:"periodNum,omitempty"`

	//Effective time
	ValidTime string `json:"validTime"`

	//Expiration time.
	ExpireTime string `json:"expireTime"`

	//Number of subscriptions
	SubscriptionNum *int `json:"subscriptionNum,omitempty"`

	//Order amount (original price).
	Currency *float64 `json:"currency,omitempty"`

	//Order amount after the discount (excluding the vouchers or cards).）
	CurrencyAfterDiscount *float64 `json:"currencyAfterDiscount,omitempty"`

	//Order amount (list price).
	CurrencyOfficial *float64 `json:"currencyOfficial,omitempty"`

	//Order details
	AmountInfo AmountInfo `json:"amountInfo"`

	//Currency code.
	CurrencyType string `json:"currencyType"`

	//Product catalog code.
	CategoryCode string `json:"categoryCode"`
}

type PayPeriodOrderResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description
	ErrorMsg string `json:"error_msg"`

	//Payment sequence number corresponding to the order.
	TradeNo string `json:"tradeNo"`

	//Information about the resources whose quota or capacity is insufficient.
	QuotaInfos []QuotaInfo `json:"quotaInfos"`

	//Information about the enterprise project whose fund is insufficient.
	EnterpriseProjectAuthResult []EnterpriseProject `json:"enterpriseProjectAuthResult"`
}

type QuotaInfo struct {
	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//Cloud service type code
	CloudServiceType string `json:"cloudServiceType"`

	//Resource type code
	ResourceType string `json:"resourceType"`

	//Verification result of the change of the cloud service quota, capacity, or specifications.
	ResourceSpecCode string `json:"resourceSpecCode"`

	//Verification result of the change of the cloud service quota, capacity, or specifications.
	AuthResult *int `json:"authResult,omitempty"`

	//AZ ID.
	AvailableZoneId string `json:"availableZoneId"`
}

type EnterpriseProject struct {
	//ID of the enterprise project to which the order belongs.
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	//Enterprise project name.
	EnterpriseProjectName string `json:"enterpriseProjectName"`

	//Verification result of the enterprise project's fund quota.
	AuthStatus *int `json:"authStatus,omitempty"`
}

type UnsubscribePeriodOrderResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//New order ID generated for the unsubscription
	UnsubOrderIds []string `json:"unsub_order_ids"`
}

type CancelOrderResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`
}

type QueryResourceStatusByOrderIdResp struct {
	//	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Total resources
	TotalSize *int `json:"totalSize,omitempty"`

	//Resource list.
	Resources []Resource `json:"resources"`
}

type Resource struct {
	//Resource instance ID.
	ResourceId string `json:"resourceId"`

	//Cloud service type code.
	CloudServiceType string `json:"cloudServiceType"`

	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//Resource type code
	ResourceType string `json:"resourceType"`

	//resourceSpecCode
	ResourceSpecCode string `json:"resourceSpecCode"`

	//Resource capacity.
	ResourceSize float64 `json:"resourceSize,omitempty"`

	//Resource capacity measurement ID
	ResouceSizeMeasureId *int `json:"resouceSizeMeasureId,omitempty"`

	//Resource provisioning status
	Status *int `json:"status,omitempty"`
}

type QueryRefundOrderAmountResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Total queries
	TotalCount *int `json:"total_count,omitempty"`

	//Resource list.
	ResourceInfoList []ResourceInfo `json:"resource_info_list"`
}

type ResourceInfo struct {
	//Record ID.
	Id string `json:"id"`

	//Resource instance ID.
	ResourceId string `json:"resource_id"`

	//Amount.
	Amount string `json:"amount"`

	//Measurement unit.
	MeasureId string `json:"measure_id"`

	//Customer ID.
	CustomerId string `json:"customer_id"`

	//Resource type code.
	ResourceType string `json:"resourceType"`

	//Cloud service type code
	CloudServiceType string `json:"cloudServiceType"`

	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//ID of the original order corresponding to the unsubscription amount, consumption amount, or unsubscription handling fee.
	PreOrderId string `json:"preOrderId"`
}

type QueryOrderListResult struct { 
	gophercloud.Result
}

type QueryOrderDetailResult struct { 
	gophercloud.Result
}

type PayPeriodOrderResult struct { 
	gophercloud.Result
}

type UnsubscribePeriodOrderResult struct { 
	gophercloud.Result
}

type CancelOrderResult struct { 
	gophercloud.Result
}

type QueryRefundOrderAmountResult struct { 
	gophercloud.Result
}

type QueryResourceStatusByOrderIdResult struct { 
	gophercloud.Result
}

func (r QueryOrderListResult) Extract() (*QueryOrderListResp, error) {
	var res *QueryOrderListResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r QueryOrderDetailResult) Extract() (*QueryOrderDetailResp, error) {
	var res *QueryOrderDetailResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r PayPeriodOrderResult) Extract() (*PayPeriodOrderResp, error) {
	var res *PayPeriodOrderResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r UnsubscribePeriodOrderResult) Extract() (*UnsubscribePeriodOrderResp, error) {
	var res *UnsubscribePeriodOrderResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r CancelOrderResult) Extract() (*CancelOrderResp, error) {
	var res *CancelOrderResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r QueryRefundOrderAmountResult) Extract() (*QueryRefundOrderAmountResp, error) {
	var res *QueryRefundOrderAmountResp
	err := r.ExtractInto(&res)
	return res, err
}

func (r QueryResourceStatusByOrderIdResult) ExtractQueryResourceStatusByOrderId() (*QueryResourceStatusByOrderIdResp, error) {
	var res *QueryResourceStatusByOrderIdResp
	err := r.ExtractInto(&res)
	return res, err
}
