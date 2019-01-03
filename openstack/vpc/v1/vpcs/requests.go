package vpcs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOpts struct {
	// Specifies the name of the VPC. The name must be unique for a
	// tenant. The value is a string of no more than 64 characters and can contain digits,
	// letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Specifies the range of available subnets in the VPC. The value
	// must be in CIDR format, for example, 192.168.0.0/16. The value ranges from 10.0.0.0/8
	// to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to
	// 192.168.255.0/24.
	Cidr string `json:"cidr,omitempty"`

	// Specifies the enterprise project ID. This field can be used to
	// filter out the VPCs associated with a specified enterprise project.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type CreateOptsBuilder interface {
	ToVPCCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToVPCCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpc")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVPCCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{OkCodes:[]int{200}})
	return
}

func Delete(client *gophercloud.ServiceClient, vpcId string) (r DeleteResult) {
	url := DeleteURL(client, vpcId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, vpcId string) (r GetResult) {
	url := GetURL(client, vpcId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}

type ListOpts struct {
	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	//The value can contain a maximum of 36 characters.
	//It is string "0"&nbsp;or in UUID format with hyphens (-). Value&nbsp;
	//"0" indicates the default enterprise project.Specifies the enterprise project ID.
	//This field can be used to filter the security groups of an enterprise project.
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

type ListOptsBuilder interface {
	ToVPCListQuery() (string, error)
}

func (opts ListOpts) ToVPCListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToVPCListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VpcPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOpts struct {
	// Specifies the name of the VPC. The name must be unique for a
	// tenant. The value is a string of no more than 64 characters and can contain digits,
	// letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Specifies the range of available subnets in the VPC. The value
	// must be in CIDR format, for example, 192.168.0.0/16. The value ranges from 10.0.0.0/8
	// to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to
	// 192.168.255.0/24.
	Cidr string `json:"cidr,omitempty"`
}

type UpdateOptsBuilder interface {
	ToVPCUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToVPCUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpc")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, vpcId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVPCUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, vpcId), b, &r.Body, &gophercloud.RequestOpts{OkCodes:[]int{200}})
	return
}
