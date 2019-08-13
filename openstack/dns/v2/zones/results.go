/*

package zones

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a Zone.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Zone, error) {
	var s *Zone
	err := r.ExtractInto(&s)
	return s, err
}

// CreateResult is the result of a Create request. Call its Extract method
// to interpret the result as a Zone.
type CreateResult struct {
	commonResult
}

// GetResult is the result of a Get request. Call its Extract method
// to interpret the result as a Zone.
type GetResult struct {
	commonResult
}

// UpdateResult is the result of an Update request. Call its Extract method
// to interpret the result as a Zone.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method
// to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// ZonePage is a single page of Zone results.
type ZonePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (r ZonePage) IsEmpty() (bool, error) {
	s, err := ExtractZones(r)
	return len(s) == 0, err
}

// ExtractZones extracts a slice of Zones from a List result.
func ExtractZones(r pagination.Page) ([]Zone, error) {
	var s struct {
		Zones []Zone `json:"zones"`
	}
	err := (r.(ZonePage)).ExtractInto(&s)
	return s.Zones, err
}

// Zone represents a DNS zone.
type Zone struct {
	// ID uniquely identifies this zone amongst all other zones, including those
	// not accessible to the current tenant.
	ID string `json:"id"`

	// PoolID is the ID for the pool hosting this zone.
	PoolID string `json:"pool_id"`

	// ProjectID identifies the project/tenant owning this resource.
	ProjectID string `json:"project_id"`

	// Name is the DNS Name for the zone.
	Name string `json:"name"`

	// Email for the zone. Used in SOA records for the zone.
	Email string `json:"email"`

	// Description for this zone.
	Description string `json:"description"`

	// TTL is the Time to Live for the zone.
	TTL int `json:"ttl"`

	// Serial is the current serial number for the zone.
	Serial int `json:"-"`

	// Status is the status of the resource.
	Status string `json:"status"`

	// Action is the current action in progress on the resource.
	Action string `json:"action"`

	// Version of the resource.
	Version int `json:"version"`

	// Attributes for the zone.
	Attributes map[string]string `json:"attributes"`

	// Type of zone. Primary is controlled by Designate.
	// Secondary zones are slaved from another DNS Server.
	// Defaults to Primary.
	Type string `json:"type"`

	// Masters is the servers for slave servers to get DNS information from.
	Masters []string `json:"masters"`

	// CreatedAt is the date when the zone was created.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the date when the last change was made to the zone.
	UpdatedAt time.Time `json:"-"`

	// TransferredAt is the last time an update was retrieved from the
	// master servers.
	TransferredAt time.Time `json:"-"`

	// Links includes HTTP references to the itself, useful for passing along
	// to other APIs that might want a server reference.
	Links map[string]interface{} `json:"links"`
}

func (r *Zone) UnmarshalJSON(b []byte) error {
	type tmp Zone
	var s struct {
		tmp
		CreatedAt     gophercloud.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt     gophercloud.JSONRFC3339MilliNoZ `json:"updated_at"`
		TransferredAt gophercloud.JSONRFC3339MilliNoZ `json:"transferred_at"`
		Serial        interface{}                     `json:"serial"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Zone(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)
	r.TransferredAt = time.Time(s.TransferredAt)

	switch t := s.Serial.(type) {
	case float64:
		r.Serial = int(t)
	case string:
		switch t {
		case "":
			r.Serial = 0
		default:
			serial, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return err
			}
			r.Serial = int(serial)
		}
	}

	return err
}




*/

package zones

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the result of a Create request. Call its Extract method
// to interpret the result as a Zone.
type CreateResult struct {
	commonResult
}

// GetResult is the result of a Get request. Call its Extract method
// to interpret the result as a Zone.
type GetResult struct {
	commonResult
}

// UpdateResult is the result of an Update request. Call its Extract method
// to interpret the result as a Zone.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method
// to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

type ZonePage struct {
	pagination.LinkedPageBase
}

func (r ZonePage) IsEmpty() (bool, error) {
	response, err := ExtractZones(r)
	return len(response.Zones) == 0, err
}

func ExtractZones(r pagination.Page) (*ListZoneResponse, error) {
	var list ListZoneResponse
	err := (r.(ZonePage)).ExtractInto(&list)
	return &list, err
}

type ListZoneResponse struct {
	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`
	// Zone list object
	Zones []Zone `json:"zones"`

	// Number of resources that meet the filter condition
	Metadata Metadata `json:"metadata"`
}

type ListNameServersResult struct {
	commonResult
}

func (r ListNameServersResult) Extract() (*ListNameServersResponse, error) {
	var response ListNameServersResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListNameServersResponse struct {
	// Name server list object
	Nameservers []NameServer `json:"nameservers"`
}

type Metadata struct {
	// Total number of resources
	TotalCount int `json:"total_count"`
}

// ZoneCreateResponse,The response of the zone creation.
type ZoneCreateResponse struct {
	// Zone ID, which is a UUID used to identify the zone
	ID string `json:"id"`

	// Zone name
	Name string `json:"name"`

	// Zone description
	Description string `json:"description"`

	// Mail address of the administrator managing the zone
	Email string `json:"email"`

	// Zone type, which can be  or
	ZoneType string `json:"zone_type"`

	// TTL value of the SOA record set in the zone
	TTL int `json:"ttl"`

	// Serial number in the SOA record set in the zone, which
	// identifies the change on the primary DNS server
	Serial int `json:"serial"`

	// Resource status.The value can be PENDING_CREATE, ACTIVE,
	// PENDING_DELETE, or ERROR.
	Status string `json:"status"`

	// Number of record sets in the zone
	RecordNum int `json:"record_num"`

	// Pool ID of the zone, which is assigned by the system
	PoolId string `json:"pool_id"`

	// Project ID of the zone
	ProjectId string `json:"project_id"`

	// Time when the zone was created
	CreatedAt string `json:"created_at"`

	// Time when the zone was updated
	UpdatedAt string `json:"updated_at"`

	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`

	// Master DNS servers, from which the slave servers get DNS
	// information
	Masters []string `json:"masters"`

	// Routers (VPCs associated with the zone)
	Router AssociateRouterResponse `json:"router"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Self string `json:"self"`
	Next string `json:"next"`
}

type Zone struct {
	// Zone ID, which is a UUID used to identify the zone
	ID string `json:"id"`

	// Zone name
	Name string `json:"name"`

	// Zone description
	Description string `json:"description"`

	// Mail address of the administrator managing the zone
	Email string `json:"email"`

	// Zone type, which can be  or
	ZoneType string `json:"zone_type"`

	// TTL value of the SOA record set in the zone
	TTL int `json:"ttl"`

	// Serial number in the SOA record set in the zone, which
	// identifies the change on the primary DNS server
	Serial int `json:"serial"`

	// Resource status.The value can be PENDING_CREATE, ACTIVE,
	// PENDING_DELETE, or ERROR.
	Status string `json:"status"`

	// Number of record sets in the zone
	RecordNum int `json:"record_num"`

	// Pool ID of the zone, which is assigned by the system
	PoolId string `json:"pool_id"`

	// Project ID of the zone
	ProjectId string `json:"project_id"`

	// Time when the zone was created
	CreatedAt string `json:"created_at"`

	// Time when the zone was updated
	UpdatedAt string `json:"updated_at"`

	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`

	// Master DNS servers, from which the slave servers get DNS
	// information
	Masters []string `json:"masters"`

	// Routers (VPCs associated with the zone)
	Routers []AssociateRouterResponse `json:"routers"`
}

type NameServer struct {
	// IP address of a name server
	Address string `json:"address"`

	// Priority of a name server.For example, if the priority of a
	// name server is 1, it is used to resolve domain names in first priority.
	Priority int `json:"priority"`
}

type AssociateRouterResult struct {
	commonResult
}

func (r AssociateRouterResult) Extract() (*AssociateRouterResponse, error) {
	var response AssociateRouterResponse
	err := r.ExtractInto(&response)
	return &response, err
}

//Router, Router (VPC) information associated with the private zone.
type AssociateRouterResponse struct {
	// Router ID (VPC ID)
	RouterId string `json:"router_id"`

	// Region of the router (VPC).If it is left blank, the region of
	// the project in the token takes effect by default.
	RouterRegion string `json:"router_region"`

	// Task status.The value can be PENDING_CREATE, PENDING_DELETE,
	// ACTIVE, or ERROR.
	Status string `json:"status"`
}

//Router, Router (VPC) information associated with the private zone.
type DisassociateRouterResponse struct {
	// Router ID (VPC ID)
	RouterId string `json:"router_id"`

	// Region of the router (VPC).If it is left blank, the region of
	// the project in the token takes effect by default.
	RouterRegion string `json:"router_region"`

	// Task status.The value can be PENDING_CREATE, PENDING_DELETE,
	// ACTIVE, or ERROR.
	Status string `json:"status"`
}

func (r CreateResult) Extract() (*ZoneCreateResponse, error) {
	var response ZoneCreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

func (r UpdateResult) Extract() (*Zone, error) {
	var response Zone
	err := r.ExtractInto(&response)
	return &response, err
}

func (r DeleteResult) Extract() (*Zone, error) {
	var response Zone
	err := r.ExtractInto(&response)
	return &response, err
}

type DisassociateRouterResult struct {
	commonResult
}

func (r DisassociateRouterResult) Extract() (*DisassociateRouterResponse, error) {
	var response DisassociateRouterResponse
	err := r.ExtractInto(&response)
	return &response, err
}

func (r GetResult) Extract() (*Zone, error) {
	var response Zone
	err := r.ExtractInto(&response)
	return &response, err
}
