package notifications

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

type Topic struct {

	// Specifies a unified topic in SMN.
	TopicUrn string `json:"topic_urn"`

	// Specifies a notification scenario, which can be one of the
	// following:
	TopicScene []string `json:"topic_scene"`

	// Specifies the topic name in SMN.
	TopicName string `json:"topic_name"`
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type EnableResult struct {
	commonResult
}

func (r EnableResult) Extract() (*EnableResponse, error) {
	var response EnableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EnableResponse struct {

	// Specifies a unified topic in SMN.
	TopicUrn string `json:"topic_urn"`

	// Specifies a notification scenario, which can be one of the
	// following:
	TopicScene []string `json:"topic_scene"`

	// Specifies the topic name in SMN.
	TopicName string `json:"topic_name"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Specifies the AS group notification list.
	Topics []Topic `json:"topics"`
}
