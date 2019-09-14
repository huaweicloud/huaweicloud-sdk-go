package trigger

import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/pagination"
)

type TriggerPage struct {
    pagination.SinglePageBase
}

type commonResult struct {
    gophercloud.Result
}

type CreateResult struct {
    commonResult
}

type DeleteResult struct {
    gophercloud.ErrResult
}

type GetResult struct {
    commonResult
}

type Trigger struct {
    TriggerId       string                 `json:"trigger_id"`
    TriggerTypeCode string                 `json:"trigger_type_code"`
    EventData       map[string]interface{} `json:"event_data"`
    EventTypeCode   string                 `json:"event_type_code,omitempty"`
    Status          string                 `json:"trigger_status,omitempty"`
    LastUpdatedTime string                 `json:"last_updated_time,omitempty"`
    CreatedTime     string                 `json:"created_time,omitempty"`
    LastError       string                 `json:"last_error,omitempty"`
}

func (r commonResult) Extract() (*Trigger, error) {
    var s Trigger
    err := r.ExtractInto(&s)
    return &s, err
}

func ExtractList(r pagination.Page) ([]Trigger, error) {
    var s []Trigger
    err := (r.(TriggerPage)).ExtractInto(&s)
    return s, err
}
