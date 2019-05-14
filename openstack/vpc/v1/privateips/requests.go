package privateips

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOpts struct {
	// Specifies the private IP address list objects.
	Privateips []PrivateIpCreate `json:"privateips"`
}

type PrivateIpCreate struct {
	// 功能说明：分配IP的子网标识
	SubnetId string `json:"subnet_id,omitempty"`

	// Specifies the private IP address obtained.
	IpAddress string `json:"ip_address,omitempty"`
}

type CreateOptsBuilder interface {
	ToPrivateipsCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToPrivateipsCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPrivateipsCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, privateipId string) (r DeleteResult) {
	url := DeleteURL(client, privateipId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, privateipId string) (r GetResult) {
	url := GetURL(client, privateipId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
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

func List(client *gophercloud.ServiceClient, subnetId string, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client, subnetId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url,
		func(r pagination.PageResult) pagination.Page {
			return PrivateIpPage{pagination.LinkedPageBase{PageResult: r}}

		})
}
