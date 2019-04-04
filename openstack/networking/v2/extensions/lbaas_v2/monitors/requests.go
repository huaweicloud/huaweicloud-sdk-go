package monitors

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMonitorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Monitor attributes you want to see returned. SortKey allows you to
// sort by a particular Monitor attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID            string `q:"id"`
	Name          string `q:"name"`
	TenantID      string `q:"tenant_id"`
	Type          string `q:"type"`
	Delay         int    `q:"delay"`
	Timeout       int    `q:"timeout"`
	MaxRetries    int    `q:"max_retries"`
	MonitorPort   int    `q:"monitor_port"`
	DomainName    string `q:"domain_name"`
	HTTPMethod    string `q:"http_method"`
	URLPath       string `q:"url_path"`
	ExpectedCodes string `q:"expected_codes"`
	AdminStateUp  *bool  `q:"admin_state_up"`
	Limit         int    `q:"limit"`
	Marker        string `q:"marker"`
	PageReverse   *bool  `q:"page_reverse"`
}

// ToMonitorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMonitorListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// health monitors. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those health monitors that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToMonitorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return MonitorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Constants that represent approved monitoring types.
const (
	TypePING  = "PING"
	TypeTCP   = "TCP"
	TypeHTTP  = "HTTP"
	TypeHTTPS = "HTTPS"
	TypeUDP   = "UDP_CONNECT"
)

//var (
//	errDelayMustGETimeout = fmt.Errorf("Delay must be greater than or equal to timeout")
//)

// CreateOptsBuilder allows extensions to add additional parameters to the
// List request.
type CreateOptsBuilder interface {
	ToMonitorCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// The Pool to Monitor.
	PoolID string `json:"pool_id" required:"true"`

	// The type of probe, which is PING, TCP, HTTP, or HTTPS, that is
	// sent by the load balancer to verify the member state.('HTTP', 'HTTPS', 'PING', 'TCP', 'UDP_CONNECT')
	Type string `json:"type" required:"true"`

	// The time, in seconds, between sending probes to members.
	Delay int `json:"delay" required:"true"`

	// Maximum number of seconds for a Monitor to wait for a ping reply
	// before it times out. The value must be less than the delay value.
	Timeout int `json:"timeout" required:"true"`

	// Number of permissible ping failures before changing the member's
	// status to INACTIVE. Must be a number between 1 and 10.
	MaxRetries int `json:"max_retries" required:"true"`

	// URI path that will be accessed if Monitor type is HTTP or HTTPS.
	// Required for HTTP(S) types.
	URLPath string `json:"url_path,omitempty"`

	// The HTTP method used for requests by the Monitor. If this attribute
	// is not specified, it defaults to "GET". Required for HTTP(S) types.
	HTTPMethod string `json:"http_method,omitempty"`

	// Expected HTTP codes for a passing HTTP(S) Monitor. You can either specify
	// a single status like "200", or a range like "200-202". Required for HTTP(S)
	// types.
	ExpectedCodes string `json:"expected_codes,omitempty"`

	// The UUID of the tenant who owns the Monitor. Only administrative users
	// can specify a tenant UUID other than their own.
	TenantID string `json:"tenant_id,omitempty"`

	// The Name of the Monitor.
	Name string `json:"name,omitempty"`

	// The administrative state of the Monitor. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	//heath check port range in [1,65535]. default is null.
	MonitorPort int `json:"monitor_port,omitempty"`

	//domain name for http request
	DomainName string `json:"domain_name,omitempty"`
}

// ToMonitorCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToMonitorCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "healthmonitor")
	if err != nil {
		return nil, err
	}

	//switch opts.Type {
	//case TypeHTTP, TypeHTTPS:
	//	switch opts.URLPath {
	//	case "":
	//		return nil, fmt.Errorf("URLPath must be provided for HTTP and HTTPS")
	//	}
	//	switch opts.ExpectedCodes {
	//	case "":
	//		return nil, fmt.Errorf("ExpectedCodes must be provided for HTTP and HTTPS")
	//	}
	//}

	return b, nil
}

/*
 Create is an operation which provisions a new Health Monitor. There are
 different types of Monitor you can provision: PING, TCP or HTTP(S). Below
 are examples of how to create each one.

 Here is an example config struct to use when creating a PING or TCP Monitor:

 CreateOpts{Type: TypePING, Delay: 20, Timeout: 10, MaxRetries: 3}
 CreateOpts{Type: TypeTCP, Delay: 20, Timeout: 10, MaxRetries: 3}

 Here is an example config struct to use when creating a HTTP(S) Monitor:

 CreateOpts{Type: TypeHTTP, Delay: 20, Timeout: 10, MaxRetries: 3,
 HttpMethod: "HEAD", ExpectedCodes: "200", PoolID: "2c946bfc-1804-43ab-a2ff-58f6a762b505"}
*/
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMonitorCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	return
}

// Get retrieves a particular Health Monitor based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToMonitorUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// The time, in seconds, between sending probes to members.
	Delay int `json:"delay,omitempty"`

	// Maximum number of seconds for a Monitor to wait for a ping reply
	// before it times out. The value must be less than the delay value.
	Timeout int `json:"timeout,omitempty"`

	// Number of permissible ping failures before changing the member's
	// status to INACTIVE. Must be a number between 1 and 10.
	MaxRetries int `json:"max_retries,omitempty"`

	// URI path that will be accessed if Monitor type is HTTP or HTTPS.
	// Required for HTTP(S) types.
	URLPath string `json:"url_path,omitempty"`

	// The HTTP method used for requests by the Monitor. If this attribute
	// is not specified, it defaults to "GET". Required for HTTP(S) types.
	HTTPMethod string `json:"http_method,omitempty"`
	//domain name for http request
	DomainName *string `json:"domain_name,omitempty"`
	// Expected HTTP codes for a passing HTTP(S) Monitor. You can either specify
	// a single status like "200", or a range like "200-202". Required for HTTP(S)
	// types.
	ExpectedCodes string `json:"expected_codes,omitempty"`

	// The Name of the Monitor.
	Name string `json:"name,omitempty"`

	// The administrative state of the Monitor. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	//heath check port range in [1,65535]. default is null.
	MonitorPort *int `json:"monitor_port,omitempty"`
}

// ToMonitorUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToMonitorUpdateMap() (map[string]interface{}, error) {
	b,err := gophercloud.BuildRequestBody(opts, "healthmonitor")
	if err != nil {
		return nil,err
	}

	hm := b["healthmonitor"].(map[string]interface{})
	if hm["domain_name"] == "" {
		hm["domain_name"] = nil
	}
	if hm["monitor_port"] == float64(0){
		hm["monitor_port"] = nil
	}

	return b,err
}

// Update is an operation which modifies the attributes of the specified
// Monitor.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToMonitorUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular Monitor based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
