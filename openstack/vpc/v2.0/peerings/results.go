package peerings

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Peering struct {
	// Specifies a resource ID in UUID format.
	ID string `json:"id"`

	Name string `json:"name,omitempty"`

	Status string `json:"status"`

	Description string `json:"description"`
	RequestVpcInfo VPCInfo `json:"request_vpc_info"`
	AcceptVpcInfo  VPCInfo `json:"accept_vpc_info"`
}

type VPCInfo struct {
	VpcID string `json:"vpc_id"`
	TenantID string `json:"tenant_id,omitempty"`
}


type ActionResult struct {
	commonResult
}

func (r ActionResult) Extract() (*Peering, error) {
	var entity Peering
	err := r.ExtractInto(&entity)
	return &entity, err
}

type commonResult struct {
	gophercloud.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*Peering, error) {
	var entity Peering
	err := r.ExtractIntoStructPtr(&entity, "peering")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*Peering, error) {
	var entity Peering
	err := r.ExtractIntoStructPtr(&entity, "peering")
	return &entity, err
}

type PeeringPage struct {
	pagination.LinkedPageBase
}

func ExtractPeerings(r pagination.Page) ([]Peering, error) {
	var s struct {
		Peering []Peering `json:"peerings"`
	}
	err := r.(PeeringPage).ExtractInto(&s)
	return s.Peering, err
}

func (r PeeringPage) NextPageURL() (string, error) {
	publicIps, err := ExtractPeerings(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(publicIps[len(publicIps)-1].ID)
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r PeeringPage) IsEmpty() (bool, error) {
	s, err := ExtractPeerings(r)
	return len(s) == 0, err
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*Peering, error) {
	var entity Peering
	err := r.ExtractIntoStructPtr(&entity, "peering")
	return &entity, err
}
