package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type ScalingGroupInstance struct {
	// Specifies the instance ID.
	InstanceId string `json:"instance_id"`

	// Specifies the instance name.
	InstanceName string `json:"instance_name"`

	// Specifies the ID of the AS group to which the instance
	// belongs.
	ScalingGroupId string `json:"scaling_group_id"`

	// Specifies the name of the AS group to which the instance
	// belongs.
	ScalingGroupName string `json:"scaling_group_name"`

	// Specifies the instance lifecycle status in the AS
	// group.INSERVICE: The instance in the AS group is in use.PENDING: The instance is
	// being added to the AS group.PENDING_WAIT: The instance is waiting to be added to the
	// AS group.REMOVING: The instance is being removed from the AS group.REMOVING_WAIT: The
	// instance is waiting to be removed from the AS group.
	LifeCycleState string `json:"life_cycle_state"`

	// Specifies the instance health status.The status can be NORMAL
	// or ERROR.
	HealthStatus string `json:"health_status"`

	// Specifies the AS configuration name.If the AS configuration has
	// been deleted, no information is displayed.If the instance is manually added to the AS
	// group, MANNUAL_ADD is returned.
	ScalingConfigurationName string `json:"scaling_configuration_name"`

	// Specifies the AS configuration ID.
	ScalingConfigurationId string `json:"scaling_configuration_id"`

	// Specifies the time when the instance is added to the AS group.
	// The time format complies with UTC.
	CreateTime string `json:"create_time"`

	// Specifies the instance protection status.
	ProtectFromScalingDown bool `json:"protect_from_scaling_down"`
}

type ActionResult struct {
	gophercloud.ErrResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type ListResponse struct {
	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start line number.
	StartNumber int `json:"start_number"`

	// Specifies the number of query records.
	Limit int `json:"limit"`

	// Specifies details about the instances in the AS group.
	ScalingGroupInstances []ScalingGroupInstance `json:"scaling_group_instances"`
}

type InstancePage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r InstancePage) IsEmpty() (bool, error) {
	data, err := ExtractInstances(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastStartNumber returns the last service in a ListResult.
func (r InstancePage) LastStartNumber() (string, error) {
	data, err := ExtractInstances(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractInstances is a function that takes a ListResult and returns the services' information.
func ExtractInstances(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(InstancePage)).ExtractInto(&s)
	return s, err
}
