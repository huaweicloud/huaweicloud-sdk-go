package groups

import (
	"github.com/gophercloud/gophercloud/pagination"
)

// SecGroup represents a container for security group rules.
type SecGroup struct {
	ID                  string              `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	VpcID               string              `json:"vpc_id"`
	EnterpriseProjectID string              `json:"enterprise_project_id"`
	SecurityGroupRules  []SecurityGroupRule `json:"security_group_rules"`
}

type SecurityGroupRule struct {
	Direction       string `json:"direction"`
	Ethertype       string `json:"ethertype"`
	ID              string `json:"id"`
	Description     string `json:"description"`
	SecurityGroupID string `json:"security_group_id"`
	RemoteGroupID   string `json:"remote_group_id,omitempty"`
	RemoteIPPrefix string `json:"remote_ip_prefix,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	PortRangeMin   int    `json:"port_range_min,omitempty"`
	PortRangeMax   int    `json:"port_range_max,omitempty"`
}

// SecGroupPage is the page returned by a pager when traversing over a
// collection of security groups.
type SecGroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a SecGroupPage struct is empty.
func (r SecGroupPage) IsEmpty() (bool, error) {
	is, err := ExtractGroups(r)
	return len(is) == 0, err
}

// ExtractGroups accepts a Page struct, specifically a SecGroupPage struct,
// and extracts the elements into a slice of SecGroup structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractGroups(r pagination.Page) ([]SecGroup, error) {
	var s struct {
		SecGroups []SecGroup `json:"security_groups"`
	}
	err := (r.(SecGroupPage)).ExtractInto(&s)
	return s.SecGroups, err
}


func (r SecGroupPage) NextPageURL() (string, error) {
	s,err := ExtractGroups(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
}