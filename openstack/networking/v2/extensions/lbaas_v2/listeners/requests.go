package listeners

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Type Protocol represents a listener protocol.
type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolHTTP Protocol = "HTTP"
	//ProtocolHTTPS Protocol = "HTTPS"
	ProtocolTerminatedHTTPS Protocol = "TERMINATED_HTTPS"
)

// Attributes allow set to Json null
var	StringToJsonNull = []string{
	"default_tls_container_ref",
	"client_ca_tls_container_ref"}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular listener attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID                      string `q:"id"`
	Name                    string `q:"name"`
	Description             string `q:"description"`
	DefaultPoolID           string `q:"default_pool_id"`
	Protocol                string `q:"protocol"`
	ProtocolPort            int    `q:"protocol_port"`
	Limit                   int    `q:"limit"`
	Marker                  string `q:"marker"`
	PageReverse             *bool   `q:"page_reverse"`
	DefaultTlsContainerRef  string `q:"default_tls_container_ref"`
	ClientCaTlsContainerRef string `q:"client_ca_tls_container_ref"`
	MemberAddress           string `q:"member_address"`
	MemberDeviceID          string `q:"member_device_id"`
}

// ToListenerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// listeners. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those listeners that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a listener.
type CreateOpts struct {
	// Human-readable name for the Listener. Does not have to be unique.
	Name string `json:"name,omitempty"`

	// TenantID is only required if the caller has an admin role and wants
	// to create a pool for another tenant.
	TenantID string `json:"tenant_id,omitempty"`

	// Human-readable description for the Listener.
	Description string `json:"description,omitempty"`

	// The protocol - can either be TCP, UDP,HTTP or HTTPS.
	Protocol Protocol `json:"protocol" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`
	// The maximum number of connections allowed for the Listener.
	ConnLimit *int `json:"connection_limit,omitempty"`
	// The load balancer on which to provision this listener.
	LoadbalancerID string `json:"loadbalancer_id" required:"true"`
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Http2Enable *bool `json:"http_2_enable,omitempty"`

	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID string `json:"default_pool_id,omitempty"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`

	// A list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`
	//
	//
	// client ca cert id
	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref,omitempty"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "listener")
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other tenants by
// specifying a TenantID attribute different than their own.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	return
}

// Get retrieves a particular Listeners based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options for updating a Listener.
type UpdateOpts struct {
	// Human-readable name for the Listener. Does not have to be unique.
	Name string `json:"name,omitempty"`

	// Human-readable description for the Listener.
	Description string `json:"description,omitempty"`

	// The maximum number of connections allowed for the Listener.
	ConnLimit *int `json:"connection_limit,omitempty"`

	// related pool id
	DefaultPoolID *string `json:"default_pool_id,omitempty"`

	// default container cert id
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// client container cert id
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	// A list of references to TLS secrets.
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`
}

// ToListenerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	b,err := gophercloud.BuildRequestBody(opts, "listener")
	if err != nil{
		return nil,err
	}

	listener := b["listener"].(map[string]interface{})
	for _,v := range StringToJsonNull{
		if listener[v] == ""{
			listener[v] = nil
		}
	}

	return b,err
}

// Update is an operation which modifies the attributes of the specified
// Listener.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular Listeners based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
