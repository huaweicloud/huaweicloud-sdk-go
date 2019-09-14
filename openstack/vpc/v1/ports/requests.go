package ports

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type FixedIp struct {
	// Specifies the subnet ID. You cannot change the parameter
	// value.
	SubnetId string `json:"subnet_id,omitempty"`

	// Specifies the port IP address. You cannot change the parameter
	// value.
	IpAddress string `json:"ip_address,omitempty"`
}

type DnsAssignment struct {

	// 功能说明：fqdn
	Fqdn string `json:"fqdn,omitempty"`

	// 功能说明：hostname
	HostName string `json:"hostname,omitempty"`

	// 功能说明：ip_address
	IpAddress string `json:"ip_address,omitempty"`
}

type ExtraDHCPOpt struct {
	// 功能说明：Option名称
	OptName string `json:"opt_name,omitempty"`

	// 功能说明：Option值
	OptValue string `json:"opt_value,omitempty"`
}

type AllowedAddressPair struct {
	// Specifies the IP address. You cannot set it to 0.0.0.0.
	// Configure an independent security group for the port if a large CIDR block (subnet
	// mask less than 24) is configured for parameter allowed_address_pairs.
	IpAddress string `json:"ip_address,omitempty"`

	// Specifies the MAC address.
	MacAddress string `json:"mac_address,omitempty"`
}

type CreateOpts struct {
	// Specifies the port name. The value can contain no more than 255
	// characters. This parameter is left blank by default.
	Name string `json:"name,omitempty"`

	// Specifies the ID of the network to which the port belongs. The
	// network ID must be a real one in the network environment.
	NetworkId string `json:"network_id" required:"true"`

	// Specifies the administrative state of the port. The value can
	// only be?true, and the default value is?true.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Specifies the port IP address. A port supports only one fixed
	// IP address that cannot be changed.
	FixedIps []FixedIp `json:"fixed_ips,omitempty"`

	// Specifies the ID of the tenant. Only the administrator can
	// specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the status of the port. The value can
	// be?ACTIVE,?BUILD, or?DOWN.
	//Status string `json:"status,omitempty"`

	// Specifies the UUID of the security group. This attribute is
	// extended.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// 1. Specifies a set of zero or more allowed address pairs. An
	// address pair consists of an IP address and MAC address. This attribute is extended.
	// For details, see parameter?allow_address_pair. 2. The IP address cannot be?0.0.0.0.
	// 3. Configure an independent security group for the port if a large CIDR block (subnet
	// mask less than 24) is configured for parameter?allowed_address_pairs.
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// Specifies a set of zero or more extra DHCP option pairs. An
	// option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

type CreateOptsBuilder interface {
	ToPortsCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToPortsCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "port")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPortsCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, portId string) (r DeleteResult) {
	url := DeleteURL(client, portId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, portId string) (r GetResult) {
	url := GetURL(client, portId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
	// Specifies that the port ID is used as the filter.
	ID string `q:"id"`

	// Specifies that the port name is used as the filter.
	Name string `q:"name"`

	// Specifies that the administrative state is used as the
	// filter.
	AdminStateUp bool `q:"admin_state_up"`

	// Specifies that the network ID is used as the filter.
	NetworkId string `q:"network_id"`

	// Specifies that the MAC address is used as the filter.
	MacAddress string `q:"mac_address"`

	// Specifies that the device ID is used as the filter.
	DeviceId string `q:"device_id"`

	// Specifies that the device owner is used as the filter.
	DeviceOwner string `q:"device_owner"`

	// Specifies that the status is used as the filter.
	Status string `q:"status"`

	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	// Specifies that the EnterpriseProjectId is used as the filter.
	EnterpriseProjectId string `q:"enterprise_project_id"`
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
			return PortPage{pagination.LinkedPageBase{PageResult: r}}

		})
}

type UpdateOpts struct {
	// Specifies the port name. The value can contain no more than 255
	// characters. This parameter is left blank by default.
	Name string `json:"name,omitempty"`

	// Specifies the UUID of the security group. This attribute is
	// extended.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// 1. Specifies a set of zero or more allowed address pairs. An
	// address pair consists of an IP address and MAC address. This attribute is extended.
	// For details, see parameter?allow_address_pair. 2. The IP address cannot be?0.0.0.0.
	// 3. Configure an independent security group for the port if a large CIDR block (subnet
	// mask less than 24) is configured for parameter?allowed_address_pairs.
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// Specifies a set of zero or more extra DHCP option pairs. An
	// option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

type UpdateOptsBuilder interface {
	ToPortsUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToPortsUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "port")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, portId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortsUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, portId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
