package vpcs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type VPC struct {
	// Specifies a resource ID in UUID format.
	ID string `json:"id"`

	// Specifies the name of the VPC. The name must be unique for a
	// tenant. The value is a string of no more than 64 characters and can contain digits,
	// letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Specifies the range of available subnets in the VPC. The value
	// must be in CIDR format, for example, 192.168.0.0/16. The value ranges from 10.0.0.0/8
	// to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to
	// 192.168.255.0/24.
	Cidr string `json:"cidr,omitempty"`

	// Specifies the status of the VPC. The value can be CREATING, OK,
	// DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.
	Status string `json:"status"`

	// Specifies the route information.
	Routes []Route `json:"routes"`

	// Specifies the enterprise project ID. This field can be used to
	// filter out the VPCs associated with a specified enterprise project.
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

type Route struct {
	// Specifies the destination network segment of a route.
	Destination string `json:"destination"`

	// Specifies the next hop of a route.
	Nexthop string `json:"nexthop"`
}

type commonResult struct {
	gophercloud.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*VPC, error) {
	var entity VPC
	err := r.ExtractIntoStructPtr(&entity, "vpc")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*VPC, error) {
	var entity VPC
	err := r.ExtractIntoStructPtr(&entity, "vpc")
	return &entity, err
}

type VpcPage struct {
	pagination.LinkedPageBase
}

func ExtractVpcs(r pagination.Page) ([]VPC, error) {
	var s struct{
		Vpcs []VPC `json:"vpcs"`
	}
	err:=r.(VpcPage).ExtractInto(&s)
	return s.Vpcs, err
}

func (r VpcPage) NextPageURL() (string, error) {
	publicIps,err:= ExtractVpcs(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(publicIps[len(publicIps)-1].ID)
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r VpcPage) IsEmpty() (bool, error) {
	s,err:= ExtractVpcs(r)
	return len(s)==0, err
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*VPC, error) {
	var entity VPC
	err := r.ExtractIntoStructPtr(&entity, "vpc")
	return &entity, err
}
