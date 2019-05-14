package peerings

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOpts struct {
	Name           string  `json:"name" required:"true"`
	RequestVpcInfo VPCInfo `json:"request_vpc_info" required:"true"`
	AcceptVpcInfo  VPCInfo `json:"accept_vpc_info" required:"true"`
}

type CreateOptsBuilder interface {
	ToPeeringCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToPeeringCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "peering")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPeeringCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{201}})
	return
}

func Delete(client *gophercloud.ServiceClient, peeringID string) (r DeleteResult) {
	url := DeleteURL(client, peeringID)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, peeringID string) (r GetResult) {
	url := GetURL(client, peeringID)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}

type ListOpts struct {
	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	ID string `q:"id"`

	Name string `q:"name"`

	Status string `q:"status"`

	VpcID string `q:"vpc_id"`

	TenantID string `q:"tenant_id"`
}

type ListOptsBuilder interface {
	ToPeeringListQuery() (string, error)
}

func (opts ListOpts) ToPeeringListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToPeeringListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PeeringPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOpts struct {
	Name string `json:"name,omitempty"`
}

type UpdateOptsBuilder interface {
	ToPeeringUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToPeeringUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "peering")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, peeringID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPeeringUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, peeringID), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

func Accept(client *gophercloud.ServiceClient, peeringID string) (r ActionResult) {
	url := AcceptURL(client, peeringID)
	_, r.Err = client.Put(url, nil, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

func Reject(client *gophercloud.ServiceClient, peeringID string) (r ActionResult) {
	url := RejectURL(client, peeringID)
	_, r.Err = client.Put(url, nil, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}
