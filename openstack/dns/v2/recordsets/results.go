package recordsets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a RecordSet.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*RecordSet, error) {
	var s *RecordSet
	err := r.ExtractInto(&s)
	return s, err
}

// CreateResult is the result of a Create operation. Call its Extract method to
// interpret the result as a RecordSet.
type CreateResult struct {
	commonResult
}

// GetResult is the result of a Get operation. Call its Extract method to
// interpret the result as a RecordSet.
type GetResult struct {
	commonResult
}

// RecordSetPage is a single page of RecordSet results.
type RecordSetPage struct {
	pagination.LinkedPageBase
}

// UpdateResult is result of an Update operation. Call its Extract method to
// interpret the result as a RecordSet.
type UpdateResult struct {
	commonResult
}

// DeleteResult is result of a Delete operation. Call its ExtractErr method to
// determine if the operation succeeded or failed.
type DeleteResult struct {
	commonResult
}

// IsEmpty returns true if the page contains no results.
func (r RecordSetPage) IsEmpty() (bool, error) {
	response, err := ExtractRecordSets(r)
	return len(response.Recordsets) == 0, err
}

// ExtractRecordSets extracts a slice of RecordSets from a List result.
func ExtractRecordSets(r pagination.Page) (*ListRecordsetResponse, error) {
	var list ListRecordsetResponse
	err := (r.(RecordSetPage)).ExtractInto(&list)
	return &list, err
}

// RecordSet represents a DNS Record Set.
type RecordSet struct {
	// ID is the unique ID of the recordset
	ID string `json:"id"`

	// ZoneID is the ID of the zone the recordset belongs to.
	ZoneID string `json:"zone_id"`

	// ProjectID is the ID of the project that owns the recordset.
	ProjectID string `json:"project_id"`

	// Name is the name of the recordset.
	Name string `json:"name"`

	// ZoneName is the name of the zone the recordset belongs to.
	ZoneName string `json:"zone_name"`

	// Type is the RRTYPE of the recordset.
	Type string `json:"type"`

	// Records are the DNS records of the recordset.
	Records []string `json:"records"`

	// TTL is the time to live of the recordset.
	TTL int `json:"ttl"`

	// Status is the status of the recordset.
	Status string `json:"status"`

	// Description is the description of the recordset.
	Description string `json:"description"`

	// CreatedAt is the date when the recordset was created.
	CreatedAt string `json:"create_at"`

	// UpdatedAt is the date when the recordset was updated.
	UpdatedAt string `json:"update_at"`

	// default means recordset is default
	Default bool `json:"default"`

	// Links includes HTTP references to the itself,
	// useful for passing along to other APIs that might want a recordset
	// reference.
	Links Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Self string `json:"self"`
	Next string `json:"next"`
}

type ListRecordsetResponse struct {
	// Link of the current resource or other related resources.When a
	// response is broken into pages, a next link is provided to retrieve all results.
	Links Link `json:"links"`
	// Zone list object
	Recordsets []RecordSet `json:"recordsets"`

	// Number of resources that meet the filter condition
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	// Total number of resources
	TotalCount int `json:"total_count"`
}
