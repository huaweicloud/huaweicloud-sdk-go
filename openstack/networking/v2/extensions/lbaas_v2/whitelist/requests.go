package whitelist

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToWhihieListsListMap() (string, error)
}

//ListOpts allows the filtering and sorting of paginated collections through the API.
//Filtering is achieved by passing in struct field values
//that map the floating whitelist attributes you want to see returned.
//Marker and Limit are used for pagination.
type ListOpts struct {
	// Specifies the ID of the last whitelist on the previous page.
	Marker          string `q:"marker"`

	// Specifies the number of records on each page.
	Limit           int    `q:"limit"`

	// Specifies the pagination direction.
	PageReverse     bool   `q:"page_reverse"`

	// Specifies the whitelist ID.
	ID              string `q:"id"`

	// Specifies the project ID.
	TenantId        string `q:"tenant_id"`

	//Specifies the listener ID.
	ListenerId      string `q:"listener_id"`

	// Specifies whether to enable access control.
	EnableWhitelist *bool   `q:"enable_whitelist"`

	// Lists the IP addresses in the whitelist.
	Whitelist       string `q:"whitelist"`
}

// ToWhihieListsListMap formats a ListOpts into a query string.
func (opts ListOpts) ToWhihieListsListMap() (string, error) {
	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return s.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// whitelists. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those whitelists that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		queryString, err := opts.ToWhihieListsListMap()
		if err != nil {

			return pagination.Pager{Err: err}
		}
		url += queryString

	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return WhiteListPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToWhiteListCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a whitelist.
type CreateOpts struct {
	//Specifies the tenant ID.
	TenantId        string `json:"tenant_id,omitempty"`

	//Specifies the listener ID.
	ListenerId      string `json:"listener_id"  required:"true"`

	//Specifies whether to enable the access control.
	EnableWhitelist *bool   `json:"enable_whitelist,omitempty"`

	//Lists the IP addresses in the whitelist.
	Whitelist       string `json:"whitelist,omitempty"`
}

// ToWhiteListCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToWhiteListCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "whitelist")
}

// Create is an operation which provisions a new Whitelist based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Whitelist on behalf of other tenants by
// specifying a TenantID attribute different than their own.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToWhiteListCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToWhiteListUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options for updating a WhiteList.
type UpdateOpts struct {
	// Specifies whether to enable access control.
	EnableWhitelist *bool   `json:"enable_whitelist,omitempty"`

	// Lists the IP addresses in the whitelist.
	Whitelist       *string `json:"whitelist,omitempty"`
}

// ToWhiteListUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToWhiteListUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "whitelist")
}

// Update is an operation which modifies the attributes of the specified whitelist.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToWhiteListUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves a particular whitelist based on its whitelist ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// Delete will permanently delete a particular whitelist based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
