package resource

import (
	"github.com/gophercloud/gophercloud"
)

type Resources struct {
	//订单列表。
	Data []ResourceInstance `json:"data"`

	//总记录数。
	TotalCount int `json:"total_count"`

	//状态码。
	ErrorCode string `json:"error_code"`

	//错误描述。
	ErrorMsg string `json:"error_msg"`
}

type ResourceInstance struct {
	//标示要开通资源的临时ID，资源开通以后生成的ID为resourceID。对应订购关系ID。
	Id string `json:"id"`

	//资源实例ID
	ResourceId string `json:"resource_id"`

	//资源实例名
	ResourceName string `json:"resource_name"`

	//云服务资源池区域编码
	RegionCode string `json:"region_code"`

	//用户购买云服务产品的云服务类型。
	CloudServiceTypeCode string `json:"cloud_service_type_code"`

	//用户购买云服务产品的资源类型。
	ResourceTypeCode string `json:"resource_type_code"`

	//用户购买云服务产品的资源规格。
	ResourceSpecCode string `json:"resource_spec_code"`

	//资源项目ID。
	ProjectCode string `json:"project_code"`

	//产品ID
	ProductId string `json:"product_id"`

	//主资源ID
	MainResourceId string `json:"main_resource_id"`

	//是否是主资源。
	IsMainResource int `json:"is_main_resource"`

	//资源状态。
	Status int `json:"status"`

	//资源生效时间。
	ValidTime string `json:"valid_time"`

	//资源过期时间。
	ExpireTime string `json:"expire_time"`
}

type commonResult struct {
	gophercloud.Result
}

type ListResult struct {
	commonResult
}

func (r commonResult) Extract() (Resources, error) {
	var res Resources
	err := r.ExtractInto(&res)
	return res, err
}
