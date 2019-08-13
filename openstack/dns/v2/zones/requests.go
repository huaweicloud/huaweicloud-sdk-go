/*

package zones

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToZoneListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	// Integer value for the limit of values to return.
	Limit int `q:"limit"`

	// UUID of the zone at which you want to set a marker.
	Marker string `q:"marker"`

	Description string `q:"description"`
	Email       string `q:"email"`
	Name        string `q:"name"`
	SortDir     string `q:"sort_dir"`
	SortKey     string `q:"sort_key"`
	Status      string `q:"status"`
	TTL         int    `q:"ttl"`
	Type        string `q:"type"`
}

// ToZoneListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToZoneListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List implements a zone List request.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		query, err := opts.ToZoneListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ZonePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get returns information about a zone, given its ID.
func Get(client *gophercloud.ServiceClient, zoneID string) (r GetResult) {
	_, r.Err = client.Get(zoneURL(client, zoneID), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToZoneCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies the attributes used to create a zone.
type CreateOpts struct {
	// Attributes are settings that supply hints and filters for the zone.
	Attributes map[string]string `json:"attributes,omitempty"`

	// Email contact of the zone.
	Email string `json:"email,omitempty"`

	// Description of the zone.
	Description string `json:"description,omitempty"`

	// Name of the zone.
	Name string `json:"name" required:"true"`

	// Masters specifies zone masters if this is a secondary zone.
	Masters []string `json:"masters,omitempty"`

	// TTL is the time to live of the zone.
	TTL int `json:"-"`

	// Type specifies if this is a primary or secondary zone.
	Type string `json:"type,omitempty"`
}

// ToZoneCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToZoneCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.TTL > 0 {
		b["ttl"] = opts.TTL
	}

	return b, nil
}

// Create implements a zone create request.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToZoneCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToZoneUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the attributes to update a zone.
type UpdateOpts struct {
	// Email contact of the zone.
	Email string `json:"email,omitempty"`

	// TTL is the time to live of the zone.
	TTL int `json:"-"`

	// Masters specifies zone masters if this is a secondary zone.
	Masters []string `json:"masters,omitempty"`

	// Description of the zone.
	Description string `json:"description,omitempty"`
}

// ToZoneUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToZoneUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.TTL > 0 {
		b["ttl"] = opts.TTL
	}

	return b, nil
}

// Update implements a zone update request.
func Update(client *gophercloud.ServiceClient, zoneID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToZoneUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(zoneURL(client, zoneID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete implements a zone delete request.
func Delete(client *gophercloud.ServiceClient, zoneID string) (r DeleteResult) {
	_, r.Err = client.Delete(zoneURL(client, zoneID), &gophercloud.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: &r.Body,
	})
	return
}


*/

package zones

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

type ListOpts struct {
	// Number of resources returned on each page.Value range:
	// 0–500.Commonly used values are 10, 20, and 50.
	Limit int `q:"limit"`

	// Start resource ID of pagination query.If the parameter is left
	// blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Zone type, which can be public or private.public: Public zones
	// are queried.private: Private zones are queried.If the value is left blank, public
	// zones are queried by default.
	Type string `q:"type"`
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Query zones in list.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ZonePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(client *gophercloud.ServiceClient, zoneId string) (r GetResult) {
	url := GetURL(client, zoneId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}

type CreateOptsBuilder interface {
	ToZonesCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// Mail address of the administrator managing the zone
	Email string `json:"email,omitempty"`

	// Description of the domain name, which cannot exceed 255
	// characters
	Description string `json:"description,omitempty"`

	// Name of the zone to be created,If the domain name is ended with
	// a dot (.), it cannot exceed 254 characters.Otherwise, the domain name cannot exceed
	// 253 characters.
	Name string `json:"name" required:"true"`

	// Zone type. The value must be private, indicating that private
	// network domain names accessible only to hosts in specified VPCs will be queried.
	ZoneType string `json:"zone_type" required:"true"`

	// Caching period of the SOA record set (in seconds).The default
	// value is 300s.The value range is 300–2147483647.
	TTL int `json:"ttl,omitempty"`

	// Router information (VPC associated with the private zone)
	Router RouterCreateOpts `json:"router"`
}

func (opts CreateOpts) ToZonesCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToZonesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{})
	return
}

type UpdateOptsBuilder interface {
	ToZonesUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	// Mail address of the administrator managing the zone
	Email string `json:"email,omitempty"`

	// Caching period of the SOA record set (in seconds).The default
	// value is 300s.The value range is 300–2147483647.
	TTL int `json:"ttl,omitempty"`

	// Description of the domain name, which cannot exceed 255
	// characters
	Description string `json:"description,omitempty"`
}

func (opts UpdateOpts) ToZonesUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, opts UpdateOptsBuilder, ZoneID string) (r UpdateResult) {
	b, err := opts.ToZonesUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Patch(UpdateURL(client, ZoneID), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{202}})
	return
}

func Delete(client *gophercloud.ServiceClient, zoneId string) (r DeleteResult) {
	url := DeleteURL(client, zoneId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: &r.Body,
	})
	return
}

//RouterCreateOpts,the parameters of the zone create, Router (VPC) information associated with the private zone.
type RouterCreateOpts struct {
	// Router ID (VPC ID)
	RouterId string `json:"router_id" required:"true"`

	// Region of the router (VPC).If it is left blank, the region of
	// the project in the token takes effect by default.
	RouterRegion string `json:"router_region,omitempty"`
}

//Router, Router (VPC) information associated with the private zone.
type Router struct {
	// Router ID (VPC ID)
	RouterId string `json:"router_id" required:"true"`

	// Region of the router (VPC).If it is left blank, the region of
	// the project in the token takes effect by default.
	RouterRegion string `json:"router_region,omitempty"`
}

type AssociateRouterOpts struct {
	// Router information (VPC associated with the zone)
	Router Router `json:"router"`
}

type AssociateRouterOptsBuilder interface {
	ToZonesAssociateRouterMap() (map[string]interface{}, error)
}

func (opts AssociateRouterOpts) ToZonesAssociateRouterMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func AssociateRouter(client *gophercloud.ServiceClient, zoneId string, opts AssociateRouterOptsBuilder) (r AssociateRouterResult) {
	b, err := opts.ToZonesAssociateRouterMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(AssociateRouterURL(client, zoneId), b, &r.Body, &gophercloud.RequestOpts{})
	return
}

type DisassociateRouterOpts struct {
	// Router information (VPC associated with the zone)
	Router Router `json:"router"`
}

type DisassociateRouterOptsBuilder interface {
	ToZonesDisassociateRouterMap() (map[string]interface{}, error)
}

func (opts DisassociateRouterOpts) ToZonesDisassociateRouterMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DisassociateRouter(client *gophercloud.ServiceClient, zoneId string, opts DisassociateRouterOptsBuilder) (r DisassociateRouterResult) {
	b, err := opts.ToZonesDisassociateRouterMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(DisassociateRouterURL(client, zoneId), b, &r.Body, &gophercloud.RequestOpts{})
	return
}

func ListNameServers(client *gophercloud.ServiceClient, zoneId string) (r ListNameServersResult) {
	url := ListNameServersURL(client, zoneId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}
