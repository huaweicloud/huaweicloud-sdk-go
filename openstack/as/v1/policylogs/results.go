package policylogs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type commonResult struct {
	gophercloud.Result
}

type ScalingPolicyExecuteLog struct {
	// Specifies the AS policy execution status.SUCCESS: indicates
	// that the AS policy is successfully executed.FAIL: indicates that the AS policy failed
	// to be executed.EXECUTING: The task is in process.
	Status string `json:"status"`

	// Specifies the AS policy execution failure.
	FailedReason string `json:"failed_reason"`

	// Specifies the AS policy execution type.SCHEDULE: automatically
	// triggered at a specified time point,RECURRENCE: automatically triggered at a
	// specified time period,ALARM: alarm-triggered,MANUAL: manually triggered
	ExecuteType string `json:"execute_type"`

	// Specifies the time when an AS policy is executed. The time
	// format complies with UTC.
	ExecuteTime string `json:"execute_time"`

	// Specifies the ID of an AS policy execution log.
	ID string `json:"id"`

	// Specifies the tenant ID.
	TenantId string `json:"tenant_id"`

	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id"`

	// Specifies the scaling resource type.AS group:
	// SCALING_GROUP,Bandwidth: BANDWIDTH
	ScalingResourceType string `json:"scaling_resource_type"`

	// Specifies the scaling resource ID.
	ScalingResourceId string `json:"scaling_resource_id"`

	// Specifies the source value.
	OldValue string `json:"old_value"`

	//Specifies the limit value.
	LimitValue string `json:"limit_value"`

	// Specifies the target value.
	DesireValue string `json:"desire_value"`

	// Specifies the AS policy execution type.ADD: indicates adding
	// instances.REMOVE: indicates reducing instances.SET: indicates setting the number of
	// instances to a specified value.
	Type string `json:"type"`

	// Specifies the tasks contained in a scaling action based on an
	// AS policy.
	JobRecords []JobRecord `json:"job_records"`

	//key value metadata
	MetaData map[string]interface{} `json:"meta_data"`
}

type JobRecord struct {
	// Specifies the task name.
	JobName string `json:"job_name"`

	// Specifies the record type.API: API calling type,MEG: message
	// type
	RecordType string `json:"record_type"`

	// Specifies the record time.
	RecordTime string `json:"record_time"`

	// Specifies the request body. This parameter is valid only if
	// record_type is set to API.
	Request string `json:"request"`

	// Specifies the response body. This parameter is valid only if
	// record_type is set to API.
	Response string `json:"response"`

	// Specifies the returned code. This parameter is valid only if
	// record_type is set to API.
	Code string `json:"code"`

	// Specifies the message. This parameter is valid only if
	// record_type is set to MEG.
	Message string `json:"message"`

	// Specifies the execution status of the task.SUCCESS: The task is
	// successfully executed.FAIL: The task failed to be executed.
	JobStatus string `json:"job_status"`
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
	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start line number.
	StartNumber int `json:"start_number"`

	// Specifies the number of query records.
	Limit int `json:"limit"`

	// Specifies the logs of AS policy execution.
	ScalingPolicyExecuteLog []ScalingPolicyExecuteLog `json:"scaling_policy_execute_log"`
}

type PolicyLogPage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r PolicyLogPage) IsEmpty() (bool, error) {
	data, err := ExtractPolicyLogs(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastStartNumber returns the last service in a ListResult.
func (r PolicyLogPage) LastStartNumber() (string, error) {
	data, err := ExtractPolicyLogs(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractPolicyLogs is a function that takes a ListResult and returns the information.
func ExtractPolicyLogs(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(PolicyLogPage)).ExtractInto(&s)
	return s, err
}
