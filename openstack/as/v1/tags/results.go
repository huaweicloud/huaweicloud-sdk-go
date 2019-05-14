package tags

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

type Tags struct {
	// Specifies the resource tag key.
	Key string `json:"key"`

	// Specifies the resource tag values.
	Values []string `json:"values"`
}

type ListResourceTagsResult struct {
	commonResult
}

func (r ListResourceTagsResult) Extract() (*ListResourceTagsResponse, error) {
	var response ListResourceTagsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResourceTagsResponse struct {
	// Specifies the resource tag.
	Tags    []Tag `json:"tags"`
	SysTags []Tag `json:"sys_tags"`
}

type ListTenantTagsResult struct {
	commonResult
}

func (r ListTenantTagsResult) Extract() (*ListTenantTagsResponse, error) {
	var response ListTenantTagsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListTenantTagsResponse struct {
	// Specifies the resource tag.
	Tags []Tags `json:"tags"`
}

type UpdateResult struct {
	gophercloud.ErrResult
}

type ListInstanceTagsResult struct {
	commonResult
}

func (r ListInstanceTagsResult) Extract() (*ListInstanceTagsResponse, error) {
	var response ListInstanceTagsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type Resource struct {
	ResourceID     string `json:"resource_id"`
	ResourceDetail string `json:"resource_detail"`
	Tags           []Tag    `json:"tags"`
	ResourceName   string `json:"resource_name"`
}

type ListInstanceTagsResponse struct {
	Resources  []Resource `json:"resources"`
	Marker     string     `json:"marker"`
	TotalCount int        `json:"total_count"`
}
