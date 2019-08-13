package ptrs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
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

type PtrPage struct {
	pagination.LinkedPageBase
}

func (r PtrPage) IsEmpty() (bool, error) {
	response, err := ExtractPtrs(r)
	return len(response.Floatingips) == 0, err
}

func ExtractPtrs(r pagination.Page) (*ListPtrResponse, error) {

	var list ListPtrResponse
	err := (r.(PtrPage)).ExtractInto(&list)
	return &list, err
}

type ListPtrResponse struct {
	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`

	// Number of resources that meet the filter condition
	Metadata Metadata `json:"metadata"`

	// PTR record object list
	Floatingips []FloatingIp `json:"floatingips"`
}

type DeletePtrResponse struct {
	Ptrdname string `json:"ptrdname"`
}

type RestoreResult struct {
	gophercloud.ErrResult
}

type SetupResult struct {
	commonResult
}

func (r SetupResult) Extract() (*FloatingIp, error) {
	var response FloatingIp
	err := r.ExtractInto(&response)
	return &response, err
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Self string `json:"self"`
	Next string `json:"next"`
}
