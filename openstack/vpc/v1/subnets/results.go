package subnets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Subnet represents a subnet. See package documentation for a top-level
// description of what this is.
type Subnet struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Cidr             string   `json:"cidr"`
	DNSList          []string `json:"dnsList"`
	Status           string   `json:"status"`
	VpcID            string   `json:"vpc_id"`
	GatewayIP        string   `json:"gateway_ip"`
	DhcpEnable       *bool    `json:"dhcp_enable"`
	PrimaryDNS       string   `json:"primary_dns"`
	SecondaryDNS     string   `json:"secondary_dns"`
	AvailabilityZone string   `json:"availability_zone"`
	NeutronNetworkID string   `json:"neutron_network_id"`
	NeutronSubnetID  string   `json:"neutron_subnet_id"`
}

// SubnetPage is the page returned by a pager when traversing over a collection
// of subnets.
type SubnetPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of subnets has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r SubnetPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"subnets_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SubnetPage struct is empty.
func (r SubnetPage) IsEmpty() (bool, error) {
	is, err := ExtractSubnets(r)
	return len(is) == 0, err
}

// ExtractSubnets accepts a Page struct, specifically a SubnetPage struct,
// and extracts the elements into a slice of Subnet structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractSubnets(r pagination.Page) ([]Subnet, error) {
	var s struct {
		Subnets []Subnet `json:"subnets"`
	}
	err := (r.(SubnetPage)).ExtractInto(&s)
	return s.Subnets, err
}
