package resource

import (
	"github.com/gophercloud/gophercloud"
)

type ListOptsBuilder interface {
	ToResourcesListQuery() (string, error)
}

// ListOpts represents options used to get resources.
type ListOpts struct {
	//客户唯一标示。是domain ID
	CustomerId string `q:"customer_id,required"`

	//资源ID列表。是Server ID
	ResourceIds string `q:"resource_ids"`

	//订单号。
	OrderId string `q:"order_id"`

	//是否只查询主资源。
	OnlyMainResource int `q:"only_main_resource"`

	//资源状态。
	StatusList string `q:"status_list"`

	//页码
	PageNo int `q:"page_no"`

	//每页条数。
	PageSize int `q:"page_size"`
}

// ToResourcesListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourcesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

//该接口对应的api不支持类似pagination分页机制。
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	domainID := client.ProviderClient.DomainID
	url := listURL(client, domainID)
	if opts != nil {
		query, err := opts.ToResourcesListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
