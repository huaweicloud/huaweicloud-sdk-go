package routes

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOpts struct {
	Type        string `json:"type" required:"true"`
	Nexthop     string `json:"nexthop" required:"true"`
	Destination string `json:"destination" required:"true"`
	VpcID       string `json:"vpc_id" required:"true"`
}

type CreateOptsBuilder interface {
	ToRouteCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToRouteCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "route")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRouteCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, routeId string) (r DeleteResult) {
	url := DeleteURL(client, routeId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, routeId string) (r GetResult) {
	url := GetURL(client, routeId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
	// Specifies that the port ID is used as the filter.
	ID string `q:"id"`

	Type string `q:"type"`

	Destination string `q:"destination"`

	VpcID string `q:"vpc_id"`

	TenantID string `q:"tenant_id"`

	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url,
		func(r pagination.PageResult) pagination.Page {
			return RoutePage{pagination.LinkedPageBase{PageResult: r}}

		})
}
