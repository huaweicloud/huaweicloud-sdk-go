package ptrs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

type Link struct {
	Self string `json:"self"`
}

type Metadata struct {
	// Total number of resources
	TotalCount int `json:"total_count"`
}

type FloatingIp struct {
	// PTR record ID, which is in {region}:{floatingip_id} format
	ID string `json:"id"`

	// Domain name of the PTR record
	Ptrdname string `json:"ptrdname"`

	// PTR record description
	Description string `json:"description"`

	// Caching period of a PTR record (in seconds).The default value
	// is 300s.The value range is 300–2147483647.
	TTL int `json:"ttl"`

	// EIP
	Address string `json:"address"`

	// Resource status.The value can be PENDING_CREATE, ACTIVE,
	// PENDING_DELETE, PENDING_UPDATE, or ERROR.
	Status string `json:"status"`

	// Requested operation on the resource.The value can be CREATE,
	// UPDATE, or DELETE.
	Action string `json:"action"`

	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*FloatingIp, error) {
	var response FloatingIp
	err := r.ExtractInto(&response)
	return &response, err
}

type ListPage struct {
	pagination.LinkedPageBase
}

func (r ListPage) IsEmpty() (bool, error) {
	response, err := ExtractList(r)
	return len(response.Floatingips) == 0, err
}

func ExtractList(r pagination.Page) (*ListResponse, error) {

	var list ListResponse
	err := (r.(ListPage)).ExtractInto(&list)
	return &list, err
}

type ListResponse struct {
	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`

	// Number of resources that meet the filter condition
	Metadata Metadata `json:"metadata"`

	// PTR record object list
	Floatingips []FloatingIp `json:"floatingips"`
}

type RestoreResult struct {
	gophercloud.ErrResult
}

func (r RestoreResult) Extract() (*FloatingIp, error) {
	var response FloatingIp
	err := r.ExtractInto(&response)
	return &response, err
}

type SetupResult struct {
	commonResult
}

func (r SetupResult) Extract() (*FloatingIp, error) {
	var response FloatingIp
	err := r.ExtractInto(&response)
	return &response, err
}
