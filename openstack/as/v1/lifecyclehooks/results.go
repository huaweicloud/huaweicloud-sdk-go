package lifecyclehooks

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

type LifecycleHook struct {

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`

	// Specifies the lifecycle hook
	// type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type"`

	// Specifies the default lifecycle hook callback
	// operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result"`

	// Specifies the lifecycle hook timeout duration in the unit of
	// second.
	DefaultTimeout int `json:"default_timeout"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata"`

	// Specifies the time when the lifecycle hook is created. The time
	// is UTC-compliant.
	CreateTime string `json:"create_time"`
}

type InstanceHangingInfo struct {

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`

	// Specifies the lifecycle action key, which determines the
	// lifecycle callback object.
	LifecycleActionKey string `json:"lifecycle_action_key"`

	// Specifies the AS instance ID.
	InstanceId string `json:"instance_id"`

	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id"`

	// Specifies the lifecycle hook status.HANGING: suspends the
	// instance.CONTINUE: continues the instance.ABANDON: terminates the instance.
	LifecycleHookStatus string `json:"lifecycle_hook_status"`

	// Specifies the timeout duration in the format of
	// "YYYY-MM-DDThh:mm:ssZ". The time is UTC-compliant.
	Timeout string `json:"timeout"`

	// Specifies the default lifecycle hook callback operation.
	DefaultResult string `json:"default_result"`
}

type CallBackResult struct {
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

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`

	// Specifies the lifecycle hook
	// type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type"`

	// Specifies the default lifecycle hook callback
	// operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result"`

	// Specifies the lifecycle hook timeout duration in the unit of
	// second.
	DefaultTimeout int `json:"default_timeout"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata"`

	// Specifies the time when the lifecycle hook is created. The time
	// is UTC-compliant.
	CreateTime string `json:"create_time"`
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`

	// Specifies the lifecycle hook
	// type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type"`

	// Specifies the default lifecycle hook callback
	// operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result"`

	// Specifies the lifecycle hook timeout duration in the unit of
	// second.
	DefaultTimeout int `json:"default_timeout"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata"`

	// Specifies the time when the lifecycle hook is created. The time
	// is UTC-compliant.
	CreateTime string `json:"create_time"`
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

	// Specifies lifecycle hooks.
	LifecycleHooks []LifecycleHook `json:"lifecycle_hooks"`
}

type ListWithSuspensionResult struct {
	commonResult
}

func (r ListWithSuspensionResult) Extract() (*ListWithSuspensionResponse, error) {
	var response ListWithSuspensionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListWithSuspensionResponse struct {

	// Specifies lifecycle hook information about an AS instance.
	InstanceHangingInfo []InstanceHangingInfo `json:"instance_hanging_info"`
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

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`

	// Specifies the lifecycle hook
	// type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type"`

	// Specifies the default lifecycle hook callback
	// operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result"`

	// Specifies the lifecycle hook timeout duration in the unit of
	// second.
	DefaultTimeout int `json:"default_timeout"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata"`

	// Specifies the time when the lifecycle hook is created. The time
	// is UTC-compliant.
	CreateTime string `json:"create_time"`
}
