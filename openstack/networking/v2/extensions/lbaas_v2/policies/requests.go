package policies

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPolicyListMap() (string, error)
}

//ListOpts allows the filtering and sorting of paginated collections through the API.
//Filtering is achieved by passing in struct field values
//that map the floating policies attributes you want to see returned.
//Marker and Limit are used for pagination.
type ListOpts struct {
	// Specifies the ID of the backend server group to which the requests are forwarded.
	RedirectPoolID     string `q:"redirect_pool_id"`

	// Provides supplementary information about the forwarding policy.
	Description        string `q:"description"`

	// Specifies the administrative status.
	AdminStateUp       *bool  `q:"admin_state_up"`

	// Specifies the project ID.
	TenantID           string `q:"tenant_id"`

	// Specifies the ID of the listener for which the forwarding policy is added.
	ListenerID         string `q:"listener_id"`

	// This field is not in use yet.
	RedirectURL        string `q:"redirect_url"`

	// Specifies the URL matching rule.
	// The value can be REDIRECT_TO_POOL or REDIRECT_TO_LISTENER.
	Action             string `q:"action"`

	// Specifies the forwarding priority. The value ranges from 1 to 100.
	Position           int    `q:"position"`

	// Specifies the forwarding policy ID.
	ID                 string `q:"id"`

	// Specifies the forwarding policy name.
	Name               string `q:"name"`

	// Specifies the number of records on each page.
	Limit              int    `q:"limit"`

	// Specifies the ID of the last forwarding policy on the previous page.
	Marker             string `q:"marker"`

	// Specifies the pagination direction.
	PageReverse        *bool  `q:"page_reverse"`

	// Specifies the provisioning status.
	// The value can be ACTIVE, PENDING_CREATE, or ERROR.
	ProvisioningStatus string `q:"provisioning_status"`
}

// ToPolicyListMap formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListMap() (string, error) {
	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}

	return s.String(), err
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a policy.
type CreateOpts struct {
	// Specifies the ID of the backend server group to which the requests are forwarded.
	RedirectPoolID string `json:"redirect_pool_id" required:"true"`

	// Provides supplementary information about the forwarding policy.
	Description    string `json:"description,omitempty"`

	// Specifies the administrative status.
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`

	// Specifies the project ID.
	TenantID       string `json:"tenant_id,omitempty"`

	// Specifies the ID of the listener for which the forwarding policy is added.
	ListenerID     string `json:"listener_id" required:"true"`

	// This field is not in use yet.
	RedirectURL    string `json:"redirect_url,omitempty"`

	// Specifies the URL matching rule.
	// The value can be REDIRECT_TO_POOL or REDIRECT_TO_LISTENER.
	Action         string `json:"action" required:"true"`

	// Specifies the forwarding priority. The value ranges from 1 to 100.
	Position       int    `json:"position,omitempty"`

	// Specifies the forwarding policy name.
	Name           string `json:"name,omitempty"`
}

const RedirectToPool = "REDIRECT_TO_POOL"

// ToPolicyCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {

	b, err := gophercloud.BuildRequestBody(opts, "l7policy")

	if err != nil {
		return nil, err
	}

	if opts.Action != RedirectToPool {

		message := fmt.Sprintf(gophercloud.CE_InvalidInputMessage, "Action only support for REDIRECT_TO_POOL")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_InvalidInputCode, message)
		return nil, err
	}

	return b, nil

}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options for updating a policy.
type UpdateOpts struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	RedirectPoolID string `json:"redirect_pool_id,omitempty"`
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`
}

// ToPolicyUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "l7policy")
}

// Create is an operation which provisions a new policy based on the
// configuration defined in the CreateOpts struct.
func Create(sc *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {

	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = sc.Post(rootURL(sc), b, &r.Body, nil)
	return

}

// List returns a Pager which allows you to iterate over a collection of
// policys.
func List(sc *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {

	url := rootURL(sc)

	if opts != nil {
		queryString, err := opts.ToPolicyListMap()
		if err != nil {

			return pagination.Pager{Err: err}
		}

		url += queryString

	}

	return pagination.NewPager(sc, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.LinkedPageBase{PageResult: r}}
	})

}

// Get retrieves a particular policy based on its policy ID.
func Get(sc *gophercloud.ServiceClient, id string) (r GetResult) {

	_, r.Err = sc.Get(resourceURL(sc, id), &r.Body, nil)
	return

}

// Update is an operation which modifies the attributes of the specified policy.
func Update(sc *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {

	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = sc.Put(resourceURL(sc, id), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200, 201}})

	return

}

// Delete will permanently delete a particular policy based on its policy ID.
func Delete(sc *gophercloud.ServiceClient, id string) (r DeleteResult) {

	_, r.Err = sc.Delete(resourceURL(sc, id), nil)

	return
}

// RulesListOptsBuilder allows extensions to add additional parameters to the
// List request.
type RulesListOptsBuilder interface {
	ToPolicyRulesListMap() (string, error)
}

//RulesListOpts allows the filtering through the API.
//Filtering is achieved by passing in struct field values
//that map the policyrules attributes you want to see returned.
type RulesListOpts struct {
	// Specifies the forwarding rule ID.
	ID           string `q:"id"`

	// Specifies the project ID.
	TenantID     string `q:"tenant_id"`

	// Specifies the administrative status.
	AdminStateUp *bool  `q:"admin_state_up"`

	// Specifies the matching content.
	// The value can be HOST_NAME or PATH.
	Type         string `q:"type"`

	// Specifies the matching mode.
	CompareType  string `q:"compare_type"`

	// Specifies whether reverse match is supported.
	Invert       *bool  `q:"invert"`

	// Specifies the Key of the matching content.
	Key          string `q:"key"`

	// Specifies the value of the matching content.
	Values       string `q:"values"`
}

// ToPolicyRulesListMap formats a RulesListOpts into a query string.
func (opts RulesListOpts) ToPolicyRulesListMap() (string, error) {
	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}

	return s.String(), err
}

// ListRules returns a Pager which allows you to iterate over a collection of policyrules.
func ListRules(sc *gophercloud.ServiceClient, opts RulesListOptsBuilder, policyId string) pagination.Pager {

	url := rulesrootURL(sc, policyId)

	if opts != nil {
		queryString, err := opts.ToPolicyRulesListMap()
		if err != nil {

			return pagination.Pager{Err: err}
		}

		url += queryString

	}

	return pagination.NewPager(sc, url, func(r pagination.PageResult) pagination.Page {
		return PolicyRulesPage{pagination.LinkedPageBase{PageResult: r}}
	})

}

// Get retrieves a particular policyrule based on its policyrule ID.
func GetRule(sc *gophercloud.ServiceClient, policyId string, policyruleId string) (r RuleGetResult) {

	_, r.Err = sc.Get(rulesresourceURL(sc, policyId, policyruleId), &r.Body, nil)
	return

}

// CreateRuleOptsBuilder allows extensions to add additional parameters to the
// CreateRule request.
type CreateRuleOptsBuilder interface {
	ToPolicyRuleCreateMap() (map[string]interface{}, error)
}

// CreateRuleOpts represents options for creating a policyrule.
type CreateRuleOpts struct {
	// Specifies the project ID.
	TenantId     string `json:"tenant_id,omitempty"`

	// Specifies the administrative status.
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`

	// Specifies the matching content.
	// The value can be HOST_NAME or PATH.
	Type         string `json:"type" required:"true"`

	// Specifies the matching mode.
	CompareType  string `json:"compare_type" required:"true"`

	// Specifies the Key of the matching content.
	Key          string `json:"key,omitempty"`

	// Specifies the value of the matching content.
	Value        string `json:"value" required:"true"`
}

// ToPolicyRuleCreateMap builds a request body from CreateRuleOpts.
func (opts CreateRuleOpts) ToPolicyRuleCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "rule")
}

// Create is an operation which provisions a new policyrule based on
// the configuration defined in the CreateRuleOpts struct.
func CreateRule(sc *gophercloud.ServiceClient, opts CreateRuleOptsBuilder, policyid string) (r RuleCreateResult) {
	b, err := opts.ToPolicyRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = sc.Post(rulesrootURL(sc, policyid), b, &r.Body, nil)
	return
}

// Delete will permanently delete a particular policyrule based on its policyrule ID.
func DeleteRule(sc *gophercloud.ServiceClient, policyId string, policyruleId string) (r RuleDeleteResult) {

	_, r.Err = sc.Delete(rulesresourceURL(sc, policyId, policyruleId), nil)

	return
}

// UpdateRuleOptsBuilder allows extensions to add additional parameters to the
// UpdateRule request.
type UpdateRuleOptsBuilder interface {
	ToPolicyRuleUpdateMap() (map[string]interface{}, error)
}

// RuleUpdateOpts represents options for updating a policyrule.
type RuleUpdateOpts struct {
	// Specifies the administrative status.
	AdminStateUp *bool  `json:"admin_state_up,omitempty"`

	// Specifies the matching mode.
	CompareType  string `json:"compare_type,omitempty"`

	// Specifies whether reverse match is supported.
	Invert       string `json:"invert,omitempty"`

	// Specifies the Key of the matching content.
	Key          string `json:"key,omitempty"`

	// Specifies the value of the matching content.
	Value        string `json:"value,omitempty"`
}

// ToPolicyRuleUpdateMap builds a request body from RuleUpdateOpts.
func (opts RuleUpdateOpts) ToPolicyRuleUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "rule")
}

// UpdateRule is an operation which modifies the attributes of the specified policyrule.
func UpdateRule(sc *gophercloud.ServiceClient, opts UpdateRuleOptsBuilder, policyId string, policyruleId string) (r RuleUpdateResult) {
	b, err := opts.ToPolicyRuleUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = sc.Put(rulesresourceURL(sc, policyId, policyruleId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
