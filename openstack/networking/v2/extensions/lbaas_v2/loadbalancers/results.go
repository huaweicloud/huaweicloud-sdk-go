package loadbalancers

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// LoadBalancer is the primary load balancing configuration object that
// specifies the virtual IP address on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type LoadBalancer struct {
	// Human-readable description for the Loadbalancer.
	Description string `json:"description"`

	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// Owner of the LoadBalancer.
	TenantID string `json:"tenant_id"`

	// The provisioning status of the LoadBalancer.
	// This value is ACTIVE, PENDING_CREATE or ERROR.
	ProvisioningStatus string `json:"provisioning_status"`

	// The IP address of the Loadbalancer.
	VipAddress string `json:"vip_address"`

	// The UUID of the port associated with the IP address.
	VipPortID string `json:"vip_port_id"`

	// The UUID of the subnet on which to allocate the virtual IP for the
	// Loadbalancer address.
	VipSubnetID string `json:"vip_subnet_id"`

	// The unique ID for the LoadBalancer.
	ID string `json:"id"`

	// The operating status of the LoadBalancer. This value is ONLINE or OFFLINE.
	OperatingStatus string `json:"operating_status"`

	// Human-readable name for the LoadBalancer. Does not have to be unique.
	Name string `json:"name"`
	// The name of the provider.
	Provider string `json:"provider"`
	// tags for the Loadbalancer.
	Tags [] string `json:"tags"`
	// Listeners are the listeners related to this Loadbalancer.
	//Listeners []listeners.Listener `json:"listeners"`
	Listeners []Listeners `json:"listeners"`
	//related pool id
	Pools []Pools `json:"pools"`
}

type Listeners struct {
	ID string `json:"id"`
}
type Pools struct {
	ID string `json:"id"`
}

type LoadbalancerStatus struct {
	Name               string     `json:"name"`
	ProvisioningStatus string     `json:"provisioning_status"`
	Listeners          []Listener `json:"listeners"`
	Pools              []Pool     `json:"pools"`
	ID                 string     `json:"id"`
	OperatingStatus    string     `json:"operating_status"`
}

type Listener struct {
	Name               string        `json:"name"`
	ProvisioningStatus string        `json:"provisioning_status"`
	Pools              []Pool        `json:"pools"`
	L7Policies         []interface{} `json:"l7policies"`
	ID                 string        `json:"id"`
	OperatingStatus    string        `json:"operating_status"`
}

type Pool struct {
	Name               string        `json:"name"`
	ProvisioningStatus string        `json:"provisioning_status"`
	HealthMonitor      HealthMonitor `json:"healthmonitor"`
	Members            []Member      `json:"members"`
	ID                 string        `json:"id"`
	OperatingStatus    string        `json:"operating_status"`
}

type Member struct {
	ProtocolPort       int    `json:"protocol_port"`
	Address            string `json:"address"`
	ID                 string `json:"id"`
	OperatingStatus    string `json:"operating_status"`
	ProvisioningStatus string `json:"provisioning_status"`
}

type HealthMonitor struct {
	Type               string `json:"type"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	ProvisioningStatus string `json:"provisioning_status"`
}

// StatusTree represents the status of a loadbalancer.
type StatusTree struct {
	Loadbalancer *LoadbalancerStatus `json:"loadbalancer"`
}

//type StatsCount struct {
//	    BytesIn           int `json:"bytes_in"`
//		TotalConnections  int `json:"total_connections"`
//		ActiveConnections int `json:"active_connections"`
//		BytesOut          int `json:"bytes_out"`
//}
// LoadBalancerPage is the page returned by a pager when traversing over a
// collection of load balancers.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of load balancers has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r LoadBalancerPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"loadbalancers_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (p LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(p)
	return len(is) == 0, err
}

// ExtractLoadBalancers accepts a Page struct, specifically a LoadbalancerPage
// struct, and extracts the elements into a slice of LoadBalancer structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	var s struct {
		LoadBalancers []LoadBalancer `json:"loadbalancers"`
	}
	err := (r.(LoadBalancerPage)).ExtractInto(&s)
	return s.LoadBalancers, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*LoadBalancer, error) {
	var s struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancer, err
}

// GetStatusesResult represents the result of a GetStatuses operation.
// Call its Extract method to interpret it as a StatusTree.
type GetStatusesResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r GetStatusesResult) Extract() (*StatusTree, error) {
	var s struct {
		Statuses *StatusTree `json:"statuses"`
	}

	err := r.ExtractInto(&s)
	return s.Statuses, err
}

//type GetstatsResult struct {
//	gophercloud.Result
//}
//
//func (r GetstatsResult)Extract()(*StatsCount,error)  {
//
//	var s struct{
//		StatsCount *StatsCount `json:"stats"`
//	}
//
//	err:= r.ExtractInto(&s)
//	return s.StatsCount ,err
//
//}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a LoadBalancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
