package logs

import (
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type ScalingActivityLog struct {
	// Specifies the status of the scaling action.SUCCESS: indicates
	// the scaling action is successfully performed.FAIL: indicates the action failed to be
	// performed.DOING: indicates the scaling action is being performed.
	Status string `json:"status"`

	// Specifies the start time of the scaling action. The time format
	// must comply with UTC.
	StartTime string `json:"start_time"`

	// Specifies the end time of the scaling action. The time format
	// must comply with UTC.
	EndTime string `json:"end_time"`

	// Specifies the scaling action log ID.
	ID string `json:"id"`

	// Specifies the name list of the instances removed from the AS
	// group after the scaling action is complete. The instance names are separated by
	// commas (,).
	InstanceRemovedList string `json:"instance_removed_list"`

	// Specifies the name list of the instances deleted from the AS
	// group and deleted after the scaling action is complete. The instance names are
	// separated by commas (,).
	InstanceDeletedList string `json:"instance_deleted_list"`

	// Specifies the name list of the instances added to the AS group
	// after the scaling action is complete. The instance names are separated by commas
	// (,).
	InstanceAddedList string `json:"instance_added_list"`

	// Specifies the number of added or deleted instances during the
	// scaling.
	ScalingValue int `json:"scaling_value"`

	// Specifies the description of the scaling action.
	Description string `json:"description"`

	// Specifies the number of instances in the AS group.
	InstanceValue int `json:"instance_value"`

	// Specifies the expected number of instances in the scaling
	// action.
	DesireValue int `json:"desire_value"`
}

type ListResponse struct {
	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start line number.
	StartNumber int `json:"start_number"`

	// Specifies the number of query records.
	Limit int `json:"limit"`

	// Specifies the scaling action log list.
	ScalingActivityLog []ScalingActivityLog `json:"scaling_activity_log"`
}

type LogPage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r LogPage) IsEmpty() (bool, error) {
	data, err := ExtractLogs(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastStartNumber returns the last service in a ListResult.
func (r LogPage) LastStartNumber() (string, error) {
	data, err := ExtractLogs(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractLogs is a function that takes a ListResult and returns the information.
func ExtractLogs(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(LogPage)).ExtractInto(&s)
	return s, err
}
