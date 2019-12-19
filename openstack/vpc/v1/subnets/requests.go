package subnets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSubnetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the subnet attributes you want to see returned.
type ListOpts struct {
	//Specifies that the VPC ID is used as the filtering condition.
	VpcID string `q:"vpc_id"`

	//Specifies the number of records returned on each page.
	//The value ranges from 0 to intmax.
	Limit int `q:"limit"`

	//Specifies the resource ID of pagination query.
	//If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`
}

// ToSubnetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSubnetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// subnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those subnets that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSubnetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type CreateOpts struct {
	// Specifies the subnet name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`

	// Specifies the subnet description. The value is a string of 0 to 255
	// characters and cannot contain "<" or ">".
	Description string `json:"description,omitempty"`

	// Specifies the network segment on which the subnet resides. The
	// value must be in CIDR format. The value must be within the CIDR block of the VPC. The
	// subnet mask cannot be greater than 28.
	Cidr string `json:"cidr" required:"true"`

	// Specifies the gateway of the subnet. The value must be a valid
	// IP address. The value must be an IP address in the subnet segment.
	GatewayIP string `json:"gateway_ip" required:"true"`

	// Specifies whether the DHCP function is enabled for the subnet.
	// The value can be true or false. If this parameter is left blank, it is set to true by
	// default.
	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The
	// value must be a valid IP address.
	PrimaryDNS string `json:"primary_dns,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The
	// value must be a valid IP address.
	SecondaryDNS string `json:"secondary_dns,omitempty"`

	// Specifies the DNS server address list of a subnet. This field
	// is required if you need to use more than two DNS servers. This parameter value is the
	// superset of both DNS server address 1 and DNS server address 2.
	DNSList []string `json:"dnsList,omitempty"`

	// Identifies the availability zone (AZ) to which the subnet
	// belongs. The value must be an existing AZ in the system.
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs.
	VpcID string `json:"vpc_id" required:"true"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "subnet")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

func Delete(client *gophercloud.ServiceClient, vpcId string, subnetId string) (r DeleteResult) {
	url := DeleteURL(client, vpcId, subnetId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, subnetId string) (r GetResult) {
	url := GetURL(client, subnetId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}

type UpdateOpts struct {
	// Specifies the subnet name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`

	// Specifies the subnet description. The value is a string of 0 to 255
	// characters and cannot contain "<" or ">".
	Description string `json:"description,omitempty"`

	// Specifies whether the DHCP function is enabled for the subnet.
	// The value can be true or false. If this parameter is left blank, it is set to true by
	// default.
	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The
	// value must be a valid IP address.
	PrimaryDNS string `json:"primary_dns,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The
	// value must be a valid IP address.
	SecondaryDNS string `json:"secondary_dns,omitempty"`

	// Specifies the DNS server address list of a subnet. This field
	// is required if you need to use more than two DNS servers. This parameter value is the
	// superset of both DNS server address 1 and DNS server address 2.
	DNSList []string `json:"dnsList,omitempty"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "subnet")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, vpcId string, subnetId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, vpcId, subnetId), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}
