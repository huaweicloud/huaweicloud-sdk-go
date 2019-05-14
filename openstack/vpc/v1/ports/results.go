package ports

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

type Port struct {
	// Specifies the port ID, which uniquely identifies the port.
	ID string `json:"id"`

	// Specifies the port name. The value can contain no more than 255
	// characters. This parameter is left blank by default.
	Name string `json:"name"`

	// Specifies the ID of the network to which the port belongs. The
	// network ID must be a real one in the network environment.
	NetworkId string `json:"network_id"`

	// Specifies the administrative state of the port. The value can
	// only be?true, and the default value is?true.
	AdminStateUp bool `json:"admin_state_up"`

	// Specifies the port MAC address. The system automatically sets
	// this parameter, and you are not allowed to configure the parameter value.
	MacAddress string `json:"mac_address"`

	// Specifies the port IP address. A port supports only one fixed
	// IP address that cannot be changed.
	FixedIps []FixedIp `json:"fixed_ips"`

	// Specifies the ID of the device to which the port belongs. The
	// system automatically sets this parameter, and you are not allowed to configure or
	// change the parameter value.
	DeviceId string `json:"device_id"`

	// Specifies the belonged device, which can be the DHCP server,
	// router, load balancers, or Nova. The system automatically sets this parameter, and
	// you are not allowed to configure or change the parameter value.
	DeviceOwner string `json:"device_owner"`

	// Specifies the ID of the tenant. Only the administrator can
	// specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id"`

	// Specifies the status of the port. The value can
	// be?ACTIVE,?BUILD, or?DOWN.
	Status string `json:"status"`

	// Specifies the UUID of the security group. This attribute is
	// extended.
	SecurityGroups []string `json:"security_groups"`

	// 1. Specifies a set of zero or more allowed address pairs. An
	// address pair consists of an IP address and MAC address. This attribute is extended.
	// For details, see parameter?allow_address_pair. 2. The IP address cannot be?0.0.0.0.
	// 3. Configure an independent security group for the port if a large CIDR block (subnet
	// mask less than 24) is configured for parameter?allowed_address_pairs.
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`

	// Specifies a set of zero or more extra DHCP option pairs. An
	// option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts"`

	// Specifies the interface type of the port. The value can
	// be?ovs,?hw_veb, or others. This attribute is extended. This parameter is visible only
	// to administrators.
	//BindingvifType string `json:"binding:vif_type"`

	// Specifies the host ID. This parameter is visible only to
	// administrators.
	//BindinghostId string `json:"binding:host_id"`

	// Specifies the type of the bound vNIC. The value can
	// be?normal?or?direct. Parameter?normal?indicates software switching.
	// Parameter?direct?indicates SR-IOV PCIe passthrough, which is not supported.
	BindingvnicType string `json:"binding:vnic_type"`

	//port security enabled default as false
	//PortSecurityEnabled bool `json:"port_security_enabled"`

	DnsAssignment []DnsAssignment `json:"dns_assignment"`

	DnsName string `json:"dns_name"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*Port, error) {
	var entity Port
	err := r.ExtractIntoStructPtr(&entity, "port")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*Port, error) {
	var entity Port
	err := r.ExtractIntoStructPtr(&entity, "port")
	return &entity, err
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*[]Port, error) {
	var list []Port
	err := r.ExtractIntoSlicePtr(&list, "ports")
	return &list, err
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*Port, error) {
	var entity Port
	err := r.ExtractIntoStructPtr(&entity, "port")
	return &entity, err
}

func (r PortPage) IsEmpty() (bool, error) {
	list, err := ExtractPorts(r)
	return len(list) == 0, err
}

type PortPage struct {
	pagination.LinkedPageBase
}

func ExtractPorts(r pagination.Page) ([]Port, error) {
	var s struct {
		Ports []Port `json:"Ports"`
	}
	err := r.(PortPage).ExtractInto(&s)
	return s.Ports, err
}

func (r PortPage) NextPageURL() (string, error) {
	s, err := ExtractPorts(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
}
