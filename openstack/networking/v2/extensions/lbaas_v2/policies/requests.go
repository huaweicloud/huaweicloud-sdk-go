package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"fmt"
)

type ListOptsBuilder interface {
	ToPolicyListMap() (string, error)
}

type ListOpts struct {
	RedirectPoolID     string `q:"redirect_pool_id"`
	Description        string `q:"description"`
	AdminStateUp       *bool   `q:"admin_state_up"`
	TenantID           string `q:"tenant_id"`
	ListenerID         string `q:"listener_id"`
	RedirectURL        string `q:"redirect_url"`
	Action             string `q:"action"`
	Position           int    `q:"position"`
	ID                 string `q:"id"`
	Name               string `q:"name"`
	Limit              int    `q:"limit"`
	Marker             string `q:"marker"`
	PageReverse        *bool   `q:"page_reverse"`
	ProvisioningStatus string `q:"provisioning_status"`
}

func (opts ListOpts) ToPolicyListMap() (string, error) {

	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}

	return s.String(), err

}

type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	RedirectPoolID string `json:"redirect_pool_id" required:"true"`
	Description    string `json:"description,omitempty"`
	AdminStateUp   *bool   `json:"admin_state_up,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	ListenerID     string `json:"listener_id" required:"true"`
	RedirectURL    string `json:"redirect_url,omitempty"`
	Action         string `json:"action" required:"true"`
	Position       int    `json:"position,omitempty"`
	Name           string `json:"name,omitempty"`
}

const RedirectToPool = "REDIRECT_TO_POOL"

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

type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}
type UpdateOpts struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	RedirectPoolID string `json:"redirect_pool_id,omitempty"`
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`
}

func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "l7policy")
}

func Create(sc *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {

	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = sc.Post(rootURL(sc), b, &r.Body, nil)
	return

}

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

func Get(sc *gophercloud.ServiceClient, id string) (r GetResult) {

	_, r.Err = sc.Get(resourceURL(sc, id), &r.Body, nil)
	return

}

func Update(sc *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {

	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = sc.Put(resourceURL(sc, id), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200, 201}})

	return

}

func Delete(sc *gophercloud.ServiceClient, id string) (r DeleteResult) {

	_, r.Err = sc.Delete(resourceURL(sc, id), nil)

	return
}
