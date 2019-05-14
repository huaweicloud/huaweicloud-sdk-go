package securitygrouprules

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

type SecurityGroupRule struct {

	// Default Value:None. Specifies the security group rule ID.
	ID string `json:"id"`

	// Default Value:None. Provides supplementary information about
	// the security group rule.
	Description string `json:"description"`

	// Default Value:None. Specifies the ID of the belonged security
	// group.
	SecurityGroupId string `json:"security_group_id"`

	// Default Value:None. Specifies the peer ID of the belonged
	// security group.
	RemoteGroupId string `json:"remote_group_id"`

	// Default Value:None. Specifies the direction of the traffic for
	// which the security group rule takes effect.
	Direction string `json:"direction"`

	// Default Value:None. Specifies the peer IP address segment.
	RemoteIpPrefix string `json:"remote_ip_prefix"`

	// Default Value:None. Specifies the protocol type or the IP
	// protocol number.
	Protocol string `json:"protocol"`

	// Default Value:None. Specifies the maximum port number. When
	// ICMP is used, the value is the ICMP code.
	PortRangeMax *int `json:"port_range_max"`

	// Default Value:None. Specifies the minimum port number. If the
	// ICMP protocol is used, this parameter indicates the ICMP type. When the TCP or UDP
	// protocol is used, both?port_range_max?and?port_range_min?must be specified, and
	// the?port_range_max?value must be greater than the?port_range_minvalue. When the ICMP
	// protocol is used, if you specify the ICMP code (port_range_max), you must also
	// specify the ICMP type (port_range_min).
	PortRangeMin *int `json:"port_range_min"`

	// Default Value:IPv4. Specifies the network type. Only IPv4 is
	// supported.
	Ethertype string `json:"ethertype"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*SecurityGroupRule, error) {
	var entity SecurityGroupRule
	err := r.ExtractIntoStructPtr(&entity, "security_group_rule")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*SecurityGroupRule, error) {
	var entity SecurityGroupRule
	err := r.ExtractIntoStructPtr(&entity, "security_group_rule")
	return &entity, err
}

type ListPage struct {
	pagination.LinkedPageBase
}

func (r ListPage) IsEmpty() (bool, error) {
	list, err := ExtractSecurityGroupRules(r)
	return len(list) == 0, err
}

func ExtractSecurityGroupRules(r pagination.Page) ([]SecurityGroupRule, error) {
	var s struct {
		SecurityGroupRule []SecurityGroupRule `json:"security_group_rules"`
	}
	err := (r.(ListPage)).ExtractInto(&s)
	return s.SecurityGroupRule, err
}

func (r ListPage) NextPageURL() (string, error) {
	s, err := ExtractSecurityGroupRules(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
}
