package subnets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Subnet represents a subnet. See package documentation for a top-level
// description of what this is.
type Subnet struct {
	// Specifies a resource ID in UUID format.
	ID string `json:"id"`

	// Specifies the subnet name. The value is a string of 1 to 64
	// characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`

	// Specifies the network segment on which the subnet resides. The
	// value must be in CIDR format. The value must be within the CIDR block of the VPC. The
	// subnet mask cannot be greater than 28.
	Cidr string `json:"cidr"`

	// Specifies the gateway of the subnet. The value must be a valid
	// IP address. The value must be an IP address in the subnet segment.
	GatewayIP string `json:"gateway_ip"`

	// Specifies whether the DHCP function is enabled for the subnet.
	// The value can be true or false. If this parameter is left blank, it is set to true by
	// default.
	DhcpEnable bool `json:"dhcp_enable,omitempty"`

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
	VpcID string `json:"vpc_id"`

	// Specifies the status of the subnet. The value can be ACTIVE,
	// DOWN, UNKNOWN, or ERROR.
	Status string `json:"status"`

	// Specifies the network (Native OpenStack API) ID.
	NeutronNetworkID string `json:"neutron_network_id"`

	// Specifies the subnet (Native OpenStack API) ID.
	NeutronSubnetID string `json:"neutron_subnet_id"`
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
	s,err := ExtractSubnets(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
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

type commonResult struct {
	gophercloud.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*Subnet, error) {
	var entity Subnet
	err := r.ExtractIntoStructPtr(&entity, "subnet")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*Subnet, error) {
	var entity Subnet
	err := r.ExtractIntoStructPtr(&entity, "subnet")
	return &entity, err
}

type UpdateResult struct {
	commonResult
}

type UpdateResp struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (r UpdateResult) Extract() (*UpdateResp, error) {
	var entity UpdateResp
	err := r.ExtractIntoStructPtr(&entity, "subnet")
	return &entity, err
}
