package groups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the group attributes you want to see returned.
type ListOpts struct {
	//Specifies that the VPC ID is used as the filtering condition.
	VpcID               string `q:"vpc_id"`

	//Specifies the tenant ID of the operator.
	TenantID            string `q:"tenant_id"`

	//Specifies the number of records returned on each page.
	//The value ranges from 0 to intmax.
	Limit               int    `q:"limit"`

	//Specifies the resource ID of pagination query.
	//If the parameter is left blank, only resources on the first page are queried.
	Marker              string `q:"marker"`

	//The value can contain a maximum of 36 characters.
	//It is string "0"&nbsp;or in UUID format with hyphens (-). Value&nbsp;
	//"0" indicates the default enterprise project.Specifies the enterprise project ID.
	//This field can be used to filter the security groups of an enterprise project.
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

// List returns a Pager which allows you to iterate over a collection of
// security groups. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	u := rootURL(c) + q.String()
	return pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return SecGroupPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
