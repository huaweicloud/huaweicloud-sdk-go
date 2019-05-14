package groups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type commonResult struct {
	gophercloud.Result
}

type LbListener struct {
	ListenerID   string `json:"listener_id"`
	PoolID       string `json:"pool_id"`
	ProtocolPort int    `json:"protocol_port"`
	Weight       int    `json:"weight"`
}

type ScalingGroup struct {
	// Specifies the name of the AS group.
	ScalingGroupName string `json:"scaling_group_name"`

	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id"`

	// Specifies the status of the AS group.
	ScalingGroupStatus string `json:"scaling_group_status"`

	// Specifies the AS configuration ID.
	ScalingConfigurationId string `json:"scaling_configuration_id"`

	// Specifies the AS configuration name.
	ScalingConfigurationName string `json:"scaling_configuration_name"`

	// Specifies the number of current instances in the AS group.
	CurrentInstanceNumber int `json:"current_instance_number"`

	// Specifies the expected number of instances in the AS group.
	DesireInstanceNumber int `json:"desire_instance_number"`

	// Specifies the minimum number of instances in the AS group.
	MinInstanceNumber int `json:"min_instance_number"`

	// Specifies the maximum number of instances in the AS group.
	MaxInstanceNumber int `json:"max_instance_number"`

	// Specifies the cooling duration (s).
	CoolDownTime int `json:"cool_down_time"`

	// Specifies the ID of a typical ELB listener. ELB listener IDs
	// are separated using a comma (,).
	LbListenerId string `json:"lb_listener_id"`

	// This field is reserved.
	LbaasListeners []LbListener `json:"lbaas_listeners"`

	// Specifies the AZ information.
	AvailableZones []string `json:"available_zones"`

	// Specifies network information.
	Networks []Network `json:"networks"`

	// Specifies security group information.
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// Specifies the time when an AS group was created. The time
	// format complies with UTC.
	CreateTime string `json:"create_time"`

	// Specifies the ID of the VPC to which the AS group belongs.
	VpcId string `json:"vpc_id"`

	// Specifies details about the AS group.
	Detail string `json:"detail"`

	// Specifies the scaling flag of the AS group.
	IsScaling bool `json:"is_scaling"`

	// Specifies the health check method.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method"`

	// Specifies the health check interval.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time"`

	// Specifies the health check interval.
	HealthPeriodicAuditTimeGracePeriod int `json:"health_periodic_audit_grace_period"`

	// Specifies the instance removal policy.
	InstanceTerminatePolicy string `json:"instance_terminate_policy"`

	// Specifies the notification mode.
	Notifications []string `json:"notifications"`

	// Specifies whether to delete the EIP bound to the ECS when
	// deleting the ECS.
	DeletePublicip bool `json:"delete_publicip"`

	// This field is reserved.
	CloudLocationId string `json:"cloud_location_id"`

	EnterpriseProjectID string `json:"enterprise_project_id"`

	ActivityType string `json:"activity_type"`
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
	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id"`
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
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*ScalingGroup, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response.ScalingGroup, err
}

type GetResponse struct {

	// Specifies details about the AS group.
	ScalingGroup ScalingGroup `json:"scaling_group"`
}


type ListResponse struct {

	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start number of query records.
	StartNumber int `json:"start_number"`

	// Specifies the number of query records.
	Limit int `json:"limit"`

	// Specifies the scaling group list.
	ScalingGroups []ScalingGroup `json:"scaling_groups"`
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
	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id"`
}

type GroupPage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r GroupPage) IsEmpty() (bool, error) {
	data, err := ExtractGroups(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastStartNumber returns the last service in a ListResult.
func (r GroupPage) LastStartNumber() (string, error) {
	data, err := ExtractGroups(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractGroups is a function that takes a ListResult and returns the information.
func ExtractGroups(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(GroupPage)).ExtractInto(&s)
	return s, err
}
