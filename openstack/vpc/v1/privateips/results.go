package privateips

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

type PrivateIp struct {

	// Specifies the status of the private IP address. The value can
	// be?ACTIVE?or?DOWN.
	Status string `json:"status"`

	// Specifies the ID of the private IP address.
	ID string `json:"id"`

	// 功能说明：分配IP的子网标识
	SubnetId string `json:"subnet_id"`

	// Specifies the tenant ID of the operator.
	TenantId string `json:"tenant_id"`

	// Specifies the VM using the private IP address. The parameter is
	// left blank if it is not used. The value can
	// be?network:dhcp,?network:router_interface_distributed, or?compute:xxx?(xxx?specifies
	// the AZ name, for example,?compute:aa-bb-cc?indicates that the private IP address is
	// used by VM in the?aa-bb-ccAZ). The value range specifies only the type of private IP
	// addresses supported by the current service.
	DeviceOwner string `json:"device_owner"`

	// Specifies the private IP address obtained.
	IpAddress string `json:"ip_address"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*[]PrivateIp, error) {
	var list []PrivateIp
	err := r.ExtractIntoSlicePtr(&list, "privateips")
	return &list, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*PrivateIp, error) {
	var entity PrivateIp
	err := r.ExtractIntoStructPtr(&entity, "privateip")
	return &entity, err
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*[]PrivateIp, error) {
	var list []PrivateIp
	err := r.ExtractIntoSlicePtr(&list, "privateips")
	return &list, err
}

func (r PrivateIpPage) IsEmpty() (bool, error) {
	list, err := ExtractPrivateIps(r)
	return len(list) == 0, err
}

type PrivateIpPage struct {
	pagination.LinkedPageBase
}

func ExtractPrivateIps(r pagination.Page) ([]PrivateIp, error) {
	var s struct {
		PrivateIps []PrivateIp `json:"privateips"`
	}
	err := r.(PrivateIpPage).ExtractInto(&s)
	return s.PrivateIps, err
}

func (r PrivateIpPage) NextPageURL() (string, error) {
	s, err := ExtractPrivateIps(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
}
