package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type commonResult struct {
	gophercloud.Result
}

type ScalingPolicy struct {
	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id"`

	// Specifies the AS policy name.
	ScalingPolicyName string `json:"scaling_policy_name"`

	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id"`

	// Specifies the AS policy status.INSERVICE: indicates that the AS
	// policy is in use.PAUSED: indicates that the AS policy is disabled.
	PolicyStatus string `json:"policy_status"`

	// Specifies the AS policy type.ALARM: indicates that the scaling
	// action is triggered by an alarm. A value is returned for alarm_id, and no value is
	// returned for scheduled_policy.SCHEDULED: indicates that the scaling action is
	// triggered as scheduled. A value is returned for scheduled_policy, and no value is
	// returned for alarm_id, recurrence_type, recurrence_value, start_time, or
	// end_time.RECURRENCE: indicates that the scaling action is triggered periodically.
	// Values are returned for scheduled_policy, recurrence_type, recurrence_value,
	// start_time, and end_time, and no value is returned for alarm_id.
	ScalingPolicyType string `json:"scaling_policy_type"`

	// Specifies the alarm ID.
	AlarmId string `json:"alarm_id"`

	// Specifies the periodic or scheduled AS policy.
	ScheduledPolicy ScheduledPolicy `json:"scheduled_policy"`

	// Specifies the scaling action of the AS policy.
	ScalingPolicyAction ScalingPolicyAction `json:"scaling_policy_action"`

	// Specifies the cooling duration (s).
	CoolDownTime int `json:"cool_down_time"`

	// Specifies the time when an AS policy is created. The time
	// format complies with UTC.
	CreateTime string `json:"create_time"`
}

type ActionResult struct {
	gophercloud.ErrResult
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {
	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id"`
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*ScalingPolicy, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response.ScalingPolicy, err
}

type GetResponse struct {
	// Specifies details about the AS policy.
	ScalingPolicy ScalingPolicy `json:"scaling_policy"`
}

type ListResponse struct {
	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start line number.
	StartNumber int `json:"start_number"`

	// Specifies the total number of query records.
	Limit int `json:"limit"`

	// Specifies the AS policy list.
	ScalingPolicies []ScalingPolicy `json:"scaling_policies"`
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {
	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id"`
}

type PolicyPage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r PolicyPage) IsEmpty() (bool, error) {
	data, err := ExtractPolicies(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastStartNumber returns the last service in a ListResult.
func (r PolicyPage) LastStartNumber() (string, error) {
	data, err := ExtractPolicies(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractPolicies is a function that takes a ListResult and returns the information.
func ExtractPolicies(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(PolicyPage)).ExtractInto(&s)
	return s, err
}
