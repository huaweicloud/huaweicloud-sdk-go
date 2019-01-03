package ports

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a port resource.
func (r commonResult) Extract() (*Port, error) {
	var s Port
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "port")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Port.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Port.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Port.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// IP is a sub-struct that represents an individual IP.
type IP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address,omitempty"`
}

// AddressPair contains the IP Address and the MAC address.
type AddressPair struct {
	IPAddress  string `json:"ip_address,omitempty"`
	MACAddress string `json:"mac_address,omitempty"`
}

// Port represents a Neutron port. See package documentation for a top-level
// description of what this is.
type Port struct {
	// UUID for the port.
	ID string `json:"id"`

	// Network that this port is associated with.
	NetworkID string `json:"network_id"`

	// Human-readable name for the port. Might not be unique.
	Name string `json:"name"`

	// Administrative state of port. If false (down), port does not forward
	// packets.
	AdminStateUp bool `json:"admin_state_up"`

	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// Mac address to use on this port.
	MACAddress string `json:"mac_address"`

	// Specifies IP addresses for the port thus associating the port itself with
	// the subnets where the IP addresses are picked from
	FixedIPs []IP `json:"fixed_ips"`

	// Owner of network.
	TenantID string `json:"tenant_id"`

	// Identifies the entity (e.g.: dhcp agent) using this port.
	DeviceOwner string `json:"device_owner"`

	// Specifies the IDs of any security groups associated with a port.
	SecurityGroups []string `json:"security_groups"`

	// Identifies the device (e.g., virtual server) using this port.
	DeviceID string `json:"device_id"`

	// Identifies the list of IP addresses the port will recognize/accept
	AllowedAddressPairs []AddressPair `json:"allowed_address_pairs"`

	// Specifies the extended DHCP option.
	// This is an extended attribute.
	ExtraDhcpOpts []ExtraDhcpOpt `json:"extra_dhcp_opts"`

	// Specifies the port virtual interface (VIF) type.
	// The value can be ovs or hw_veb.
	// This is an extended attribute.
	// This parameter is visible only to administrators.
	BindingVifType string `json:"binding:vif_type,omitempty"`

	// Specifies the VIF details.
	// Parameter ovs_hybrid_plug specifies
	// whether the OVS/bridge hybrid mode is used.
	BindingVifDetails map[string]interface{} `json:"binding:vif_details"`

	// Specifies the host ID.
	// This is an extended attribute.
	// This parameter is visible only to administrators.
	BindingHostId string `json:"binding:host_id,omitempty"`

	// Specifies the user-defined settings.
	// This is an extended attribute.
	BindingProfile map[string]interface{} `json:"binding:profile"`

	// Specifies the type of the bound vNIC.
	// normal: Softswitch
	BindingVnicType string `json:"binding:vnic_type"`

	// Specifies whether the security option is enabled for the port.
	// If the option is not enabled, the security group
	// and DHCP snooping do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled,omitempty"`

	// Specifies the default private network domain name information of the active NIC.
	// This is an extended attribute.
	DnsAssignment []map[string]string `json:"dns_assignment"`

	// Specifies the default private network DNS name of the active NIC.
	// This is an extended attribute.
	DnsName string `json:"dns_name,omitempty"`
}

//Specifies the extended DHCP option.
type ExtraDhcpOpt struct {
	//Specifies the option name.
	OptName string `json:"opt_name"`
	//Specifies the option value.
	OptValue string `json:"opt_value"`
}

// PortPage is the page returned by a pager when traversing over a collection
// of network ports.
type PortPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of ports has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r PortPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"ports_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a PortPage struct is empty.
func (r PortPage) IsEmpty() (bool, error) {
	is, err := ExtractPorts(r)
	return len(is) == 0, err
}

// ExtractPorts accepts a Page struct, specifically a PortPage struct,
// and extracts the elements into a slice of Port structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPorts(r pagination.Page) ([]Port, error) {
	var s []Port
	err := ExtractPortsInto(r, &s)
	return s, err
}

func ExtractPortsInto(r pagination.Page, v interface{}) error {
	return r.(PortPage).Result.ExtractIntoSlicePtr(v, "ports")
}
