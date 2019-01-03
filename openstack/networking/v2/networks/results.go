package networks

import (
	"encoding/json"
	"io"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a network resource.
func (r commonResult) Extract() (*Network, error) {
	var s Network
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "network")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Network.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Network.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Network.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Network represents, well, a network.
type Network struct {
	// UUID for the network
	ID string `json:"id"`

	// Human-readable name for the network. Might not be unique.
	Name string `json:"name"`

	// The administrative state of network. If false (down), the network does not
	// forward packets.
	AdminStateUp bool `json:"admin_state_up"`

	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// Subnets associated with this network.
	Subnets []string `json:"subnets"`

	// Owner of network.
	TenantID string `json:"tenant_id"`

	// Specifies whether the network resource can be accessed by any tenant.
	Shared bool `json:"shared"`

	// Availability zone hints groups network nodes that run services like DHCP, L3, FW, and others.
	// Used to make network resources highly available.
	AvailabilityZoneHints []string `json:"availability_zone_hints"`

	// Specifies whether the network is an external network.
	// This is an extended attribute.This attribute is for administrators only.
	// Tenants cannot configure or update this attribute and can only query it.
	RouterExternal bool `json:"router:external,omitempty"`

	// Specifies the physical network used by this network.
	// This is an extended attribute.This attribute is available only to administrators.
	ProviderPhysicalNetwork  string `json:"provider:physical_network,omitempty"`

	// Specifies the network type. Only the VXLAN and GENEVE networks are supported.
	ProviderNetworkType string `json:"provider:network_type,omitempty"`

	// Specifies the network segment ID.
	// The value is a VLAN ID for a VLAN network and is a VNI for a VXLAN network.
	// This is an extended attribute.This attribute is available only to administrators.
	ProviderSegmentationId int64 `json:"provider:segmentation_id,omitempty"`

	// Specifies the availability zone of this network.
	// An empty list is returned.
	AvailabilityZones []string `json:"availability_zones,omitempty"`

	// Specifies whether the security option is enabled for the port.
	// If the option is not enabled, the security group and DHCP
	// snooping settings of all VMs in the network do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled,omitempty"`

	// Specifies the default private network DNS domain address.
	// The system automatically sets this parameter,
	// and you are not allowed to configure or change the parameter value.
	DnsDomain string `json:"dns_domain,omitempty"`
}

// NetworkPage is the page returned by a pager when traversing over a
// collection of networks.
type NetworkPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r NetworkPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"networks_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r NetworkPage) IsEmpty() (bool, error) {
	is, err := ExtractNetworks(r)
	return len(is) == 0, err
}

// ExtractNetworks accepts a Page struct, specifically a NetworkPage struct,
// and extracts the elements into a slice of Network structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractNetworks(r pagination.Page) ([]Network, error) {
	var s []Network
	err := ExtractNetworksInto(r, &s)
	return s, err
}

func ExtractNetworksInto(r pagination.Page, v interface{}) error {
	return r.(NetworkPage).Result.ExtractIntoSlicePtr(v, "networks")
}

/************ 自研 *************/
type IpUsed struct {
	NetworkIpAvail NetworkIpAvailability `json:"network_ip_availability"`
}

type NetworkIpAvailability struct {
	//网络ID
	NetworkId string `json:"network_id"`

	//网络名称
	NetworkName string `json:"network_name"`

	//租户ID
	TenantId string `json:"tenant_id"`

	//网络中已经使用的IP数目（不包含系统预留地址）
	UsedIps int `json:"used_ips"`

	//网络中IP总数（不包含系统预留地址）
	TotalIps int `json:"total_ips"`

	//子网IP使用情况的对象
	SubnetIpAvail SubnetIpAvailabilitiy `json:"subnet_ip_availabilitiy"`
}

type SubnetIpAvailabilitiy struct {
	SubnetId   string `json:"subnet_id"`
	SubnetName string `json:"subnet_name"`
	Cidr       string `json:"cidr"`
	UsedIps    int    `json:"used_ips"`
	TotalIps   int    `json:"total_ips"`
	IpVersion  int    `json:"ip_version"`
}

func (r commonResult) ExtractIpUsed() (*IpUsed, error) {
	var iu IpUsed
	err := r.ExtractIntoIpUsed(&iu)
	return &iu, err
}

func (r commonResult) ExtractIntoIpUsed(to interface{}) error {
	if r.Err != nil {
		return r.Err
	}

	if reader, ok := r.Body.(io.Reader); ok {
		if readCloser, ok := reader.(io.Closer); ok {
			defer readCloser.Close()
		}
		return json.NewDecoder(reader).Decode(to)
	}

	b, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, to)

	return err
}
