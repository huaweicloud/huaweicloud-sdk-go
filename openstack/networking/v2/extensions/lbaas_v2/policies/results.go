package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type RuleId struct {
	ID string `json:"id"`
}
type Policies struct {
	// Specifies the ID of the backend server group to which the requests are forwarded.
	RedirectPoolID     string  `json:"redirect_pool_id"`

	// Provides supplementary information about the forwarding policy.
	Description        string  `json:"description"`

	// Specifies the administrative status.
	AdminStateUp       bool    `json:"admin_state_up"`

	// Lists the forwarding rules under a forwarding policy.
	Rules              []RuleId `json:"rules"`

	// Specifies the project ID.
	TenantID           string  `json:"tenant_id"`

	// Specifies the ID of the listener for which the forwarding policy is added.
	ListenerID         string  `json:"listener_id"`

	// This field is not in use yet.
	RedirectURL        string  `json:"redirect_url"`

	// Specifies the URL matching rule.
	// The value can be REDIRECT_TO_POOL or REDIRECT_TO_LISTENER.
	Action             string  `json:"action"`

	// Specifies the forwarding priority. The value ranges from 1 to 100.
	Position           int     `json:"position"`

	// Specifies the forwarding policy ID.
	ID                 string  `json:"id"`

	// Specifies the forwarding policy name.
	Name               string  `json:"name"`

	// Specifies the provisioning status.
	// The value can be ACTIVE, PENDING_CREATE, or ERROR.
	ProvisioningStatus string  `json:"provisioning_status"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a policy.
func (r commonResult) Extract() (*Policies, error) {
	var s struct {
		Policy *Policies `json:"l7policy"`
	}
	err := r.ExtractInto(&s)
	return s.Policy, err
}

// PolicyPage is the page returned by a pager when traversing over a
// collection of policy.
type PolicyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a PolicyPage struct is empty.
func (p PolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractPolcies(p)
	return len(is) == 0, err

}

// NextPageURL will retrieve the next page URL.
func (page PolicyPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"l7policies_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractPolcies accepts a Page struct, specifically a PolicyPage struct,
// and extracts the elements into a slice of Policies structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPolcies(r pagination.Page) ([]Policies, error) {
	var s struct {
		Policies []Policies `json:"l7policies"`
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

// PolicyRulesPage is the page returned by a pager when traversing over a
// collection of policyrules.
type PolicyRulesPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a PolicyRulesPage struct is empty.
func (p PolicyRulesPage) IsEmpty() (bool, error) {
	is, err := ExtractPolicyRules(p)
	return len(is.Rules) == 0, err
}

// NextPageURL will retrieve the next page URL.
func (page PolicyRulesPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"rules_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractPolicyRules accepts a Page struct, specifically a PolicyRulesPage struct,
// and extracts the elements into a slice of policyrules structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPolicyRules(r pagination.Page) (Rules, error) {
	var s Rules
	err := (r.(PolicyRulesPage)).ExtractInto(&s)
	return s, err
}

type Rule struct {
	Rule PolicyRule `json:"rule"`
}

type Rules struct {
	Rules []PolicyRule `json:"rules"`
}

type PolicyRule struct {
	// Specifies the forwarding rule ID.
	ID           string `json:"id"`

	// Specifies the project ID.
	TenantId     string `json:"tenant_id"`

	// Specifies the administrative status.
	AdminStateUp bool   `json:"admin_state_up"`

	// Specifies the matching content.
	// The value can be HOST_NAME or PATH.
	Type         string `json:"type"`

	// Specifies the matching mode.
	CompareType  string `json:"compare_type"`

	// Specifies whether reverse match is supported.
	Invert       bool   `json:"invert"`

	// Specifies the Key of the matching content.
	Key          string `json:"key"`

	// Specifies the Value of the matching content.
	Value        string `json:"value"`
}

type rulecommonResult struct {
	gophercloud.Result
}

// RuleGetResult represents the result of a get operation. Call its Extract
// method to interpret it as a policyrule.
type RuleGetResult struct {
	rulecommonResult
}

// RuleCreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a policyrule.
type RuleCreateResult struct {
	rulecommonResult
}

// RuleDeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type RuleDeleteResult struct {
	gophercloud.ErrResult
}

// RuleUpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a policyrule.
type RuleUpdateResult struct {
	rulecommonResult
}

// Extract is a function that accepts a result and extracts a policyrule.
func (r rulecommonResult) Extract() (Rule, error) {
	var s Rule
	err := r.ExtractInto(&s)
	return s, err
}