package securitygrouprules

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOpts struct {

	// Default Value:None. Provides supplementary information about
	// the security group rule.
	Description string `json:"description,omitempty"`

	// Default Value:None. Specifies the ID of the belonged security
	// group.
	SecurityGroupId string `json:"security_group_id" required:"true"`

	// Default Value:None. Specifies the peer ID of the belonged
	// security group.
	RemoteGroupId string `json:"remote_group_id,omitempty"`

	// Default Value:None. Specifies the direction of the traffic for
	// which the security group rule takes effect.
	Direction string `json:"direction" required:"true"`

	// Default Value:None. Specifies the peer IP address segment.
	RemoteIpPrefix string `json:"remote_ip_prefix,omitempty"`

	// Default Value:None. Specifies the protocol type or the IP
	// protocol number.
	Protocol string `json:"protocol,omitempty"`

	// Default Value:None. Specifies the maximum port number. When
	// ICMP is used, the value is the ICMP code.
	PortRangeMax *int `json:"port_range_max,omitempty"`

	// Default Value:None. Specifies the minimum port number. If the
	// ICMP protocol is used, this parameter indicates the ICMP type. When the TCP or UDP
	// protocol is used, both?port_range_max?and?port_range_min?must be specified, and
	// the?port_range_max?value must be greater than the?port_range_minvalue. When the ICMP
	// protocol is used, if you specify the ICMP code (port_range_max), you must also
	// specify the ICMP type (port_range_min).
	PortRangeMin *int `json:"port_range_min,omitempty"`

	// Default Value:IPv4. Specifies the network type. Only IPv4 is
	// supported.
	Ethertype string `json:"ethertype,omitempty"`
}

type CreateOptsBuilder interface {
	ToSecuritygrouprulesCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToSecuritygrouprulesCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "security_group_rule")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSecuritygrouprulesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, securityGroupsRulesId string) (r DeleteResult) {
	url := DeleteURL(client, securityGroupsRulesId)
	_, r.Err = client.Delete(url, nil)
	return
}

func Get(client *gophercloud.ServiceClient, securityGroupsRulesId string) (r GetResult) {
	url := GetURL(client, securityGroupsRulesId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {

	// Specifies the resource ID of pagination query. If the parameter
	// is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page. The
	// value ranges from 0 to intmax.
	Limit int `q:"limit"`

	// Default Value:None. Specifies the ID of the belonged security
	// group.
	SecurityGroupId string `q:"security_group_id,omitempty"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// This interface is used to query security group rules using search
// criteria and to display the security group rules in a list.
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
			return ListPage{
				pagination.LinkedPageBase{PageResult: r},
			}

		})
}
