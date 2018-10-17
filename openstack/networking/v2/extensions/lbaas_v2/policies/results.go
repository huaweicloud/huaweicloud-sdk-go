package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Rules struct {
	ID string `json:"id"`
}
type Policies struct {
	RedirectPoolID string  `json:"redirect_pool_id"`
	Description    string  `json:"description"`
	AdminStateUp   bool    `json:"admin_state_up"`
	Rules          []Rules `json:"rules"`
	TenantID       string  `json:"tenant_id"`
	ListenerID     string  `json:"listener_id"`
	RedirectURL    string  `json:"redirect_url"`
	Action         string  `json:"action"`
	Position       int     `json:"position"`
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	ProvisioningStatus string `json:"provisioning_status"`
}

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Policies, error) {
	var s struct {
		Policy *Policies `json:"l7policy"`
	}
	err := r.ExtractInto(&s)
	return s.Policy, err
}

type PolicyPage struct {
	pagination.LinkedPageBase
}

func (p PolicyPage) IsEmpty() (bool, error) {

	is, err := ExtractPolcies(p)
	return len(is) == 0, err

}
func ExtractPolcies(r pagination.Page) ([] Policies, error) {
	var s struct {
		Policies [] Policies `json:"l7policies"`
	}

	err := (r.(PolicyPage)).ExtractInto(&s)
	return s.Policies, err

}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret the result as a policy.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret the result as a policy.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation. Call its Extract
// method to interpret the result as a policy.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
