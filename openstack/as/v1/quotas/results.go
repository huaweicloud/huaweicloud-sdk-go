package quotas

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

type Quota struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	// Specifies the quota type.
	Type string `json:"type"`

	// Specifies the used amount of the quota.When type is set to
	// scaling_Policy or scaling_Instance, this parameter is reserved, and the system
	// returns -1 as the parameter value. You can query the used quota amount of AS policies
	// and AS instances in a specified AS group. For details, see Querying Quotas for AS
	// Policies and AS Instances.
	Used int `json:"used"`

	// Specifies the total amount of the quota.
	Quota int `json:"quota"`

	// Specifies the quota upper limit.
	Max int `json:"max"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*Quota, error) {

	var s struct {
		Quotas Quota `json:"quotas"`
	}
	err := r.ExtractInto(&s)
	return &s.Quotas, err
}

type ListWithInstancesResult struct {
	commonResult
}

func (r ListWithInstancesResult) Extract() (*Quota, error) {
	var s struct {
		Quotas Quota `json:"quotas"`
	}
	err := r.ExtractInto(&s)
	return &s.Quotas, err
}
