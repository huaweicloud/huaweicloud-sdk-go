package periodresource

import "github.com/gophercloud/gophercloud"



type QueryCustomerPeriodResourcesListResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Resource list.
	Data []ResourceInstance `json:"data"`

	//Total number of records
	TotalCount *int `json:"total_count,omitempty"`
}

type ResourceInstance struct {
	//Internal ID of the resource to be provisioned
	Id string `json:"id"`

	//Resource instance ID.
	ResourceId string `json:"resource_id"`

	//Resource instance name.
	ResourceName string `json:"resource_name"`

	//Resource pool region ID of cloud services.
	RegionCode string `json:"region_code"`

	//Cloud service type code.
	cloudServiceTypeCode string `json:"cloud_service_type_code"`

	//Resource type code
	ResourceTypeCode string `json:"resource_type_code"`

	//resource_spec_code
	ResourceSpecCode string `json:"resource_spec_code"`

	//Resource project ID.
	ProjectCode string `json:"project_code"`

	//Product ID.
	ProductId string `json:"product_id"`

	//Primary resource ID.
	MainResourceId string `json:"main_resource_id"`

	//Whether a primary resource.
	IsMainResource *int `json:"is_main_resource,omitempty"`

	//Resource status.
	Status *int `json:"status,omitempty"`

	//Effective time of a resource.
	ValidTime string `json:"valid_time"`

	//Expiration time of a resource.
	ExpireTime string `json:"expire_time"`

	//Next billing policy.
	NextOperationPolicy *int `json:"next_operation_policy,omitempty"`
}

type RenewSubscriptionByResourceIdResp struct {
	//Status code。
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//List of order IDs generated when resource subscription is renewed.
	OrderIds []string `json:"order_ids"`
}

type UnsubscribeByResourceIdResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Unsubscription order IDs.
	OrderIds []string `json:"order_ids"`
}

type EnableAutoRenewResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`
}

type DisableAutoRenewResp struct {
	//Status code.
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`
}

type QueryCustomerPeriodResourcesListResult struct {
	gophercloud.Result
}

func (r QueryCustomerPeriodResourcesListResult) Extract() (*QueryCustomerPeriodResourcesListResp, error) {
	var res *QueryCustomerPeriodResourcesListResp
	err := r.ExtractInto(&res)
	return res, err
}

type RenewSubscriptionByResourceIdResult struct {
	gophercloud.Result
}

func (r RenewSubscriptionByResourceIdResult) Extract() (*RenewSubscriptionByResourceIdResp, error) {
	var res *RenewSubscriptionByResourceIdResp
	err := r.ExtractInto(&res)
	return res, err
}

type UnsubscribeByResourceIdResult struct {
	gophercloud.Result
}

func (r UnsubscribeByResourceIdResult) Extract() (*UnsubscribeByResourceIdResp, error) {
	var res *UnsubscribeByResourceIdResp
	err := r.ExtractInto(&res)
	return res, err
}


type EnableAutoRenewResult struct {
	gophercloud.Result
}

func (r EnableAutoRenewResult) Extract() (*EnableAutoRenewResp, error) {
	var res *EnableAutoRenewResp
	err := r.ExtractInto(&res)
	return res, err
}

type DisableAutoRenewResult struct {
	gophercloud.Result
}

func (r DisableAutoRenewResult) Extract() (*DisableAutoRenewResp, error) {
	var res *DisableAutoRenewResp
	err := r.ExtractInto(&res)
	return res, err
}