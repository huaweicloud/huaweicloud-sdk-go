package payperuseresource

import "github.com/gophercloud/gophercloud"



type QueryCustomerResourceResp struct {
	//Status code
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Customer resources.
	CustomerResource []CustomerResource `json:"customerResources"`

	//Total number of query records.
	TotalCount *int `json:"totalCount"`

}

type CustomerResource struct {
	//Customer resource ID.。
	CustomerResourceId string `json:"customerResourceId"`

	//Customer ID.
	CustomerId string `json:"customerId"`

	//Cloud service region code
	RegionCode string `json:"regionCode"`

	//AZ code.
	azCode string `json:"azCode"`

	//Cloud service type code
	CloudServiceTypeCode string `json:"cloudServiceTypeCode"`

	//Resource type code
	ResourceTypeCode string `json:"resourceTypeCode"`

	//Resource ID.
	ResourceId string `json:"resourceId"`

	//Resource instance name.
	ResourceName string `json:"resourceName"`

	//Effective time.
	StartTime string `json:"startTime"`

	//Expiration time.
	EndTime string `json:"endTime"`

	//Resource status.
	Status *int `json:"status,omitempty"`

	//Specifications code of the pay-per-use resource.
	ResourceSpecCode string `json:"resourceSpecCode"`

	//Resource capacity
	ResourceInfo string `json:"resourceInfo"`

	//Whether the billing mode can be changed from pay-per-use to yearly/monthly
	ChargingModeChangeFlag string `json:"chargingModeChangeFlag"`

	//Account type of the last deduction for the resource
	LastDeductType string `json:"lastDeductType"`
}

type QueryCustomerResourceResult struct {
	gophercloud.Result
}

func (r QueryCustomerResourceResult) Extract() (*QueryCustomerResourceResp, error) {
	var res *QueryCustomerResourceResp
	err := r.ExtractInto(&res)
	return res, err
}


