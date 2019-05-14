package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ScheduledPolicy struct {
	// Specifies the time when the scaling action is triggered. The
	// time format must comply with UTC.If scaling_policy_type is set to SCHEDULED, the time
	// format is YYYY-MM-DDThh:mmZ.If scaling_policy_type is set to RECURRENCE, the time
	// format is hh:mm.
	LaunchTime string `json:"launch_time" required:"true"`

	// Specifies the periodic triggering type. This parameter is
	// mandatory when scaling_policy_type is set to RECURRENCE.Daily: indicates that the
	// scaling action is triggered once a day.Weekly: indicates that the scaling action is
	// triggered once a week.Monthly indicates that the scaling action is triggered once a
	// month.
	RecurrenceType string `json:"recurrence_type,omitempty"`

	// Specifies the frequency at which scaling actions are
	// triggered.If recurrence_type is set to Daily, the value is null, indicating that the
	// scaling action is triggered once a day.If recurrence_type is set to Weekly, the value
	// ranges from 1 (Sunday) to 7 (Saturday). The digits refer to dates in each week and
	// separated by a comma. For example, 1,3,5.If recurrence_type is set to Monthly, the
	// value ranges from 1 to 31. The digits refer to the dates in each month and separated
	// by a comma, such as 1,10,13,28.
	RecurrenceValue string `json:"recurrence_value,omitempty"`

	// Specifies the start time of the scaling action triggered
	// periodically. The time format complies with UTC.The current time is used by
	// default.The time format is YYYY-MM-DDThh:mmZ.
	StartTime string `json:"start_time,omitempty"`

	// Specifies the end time of the scaling action triggered
	// periodically. The time format complies with UTC. This parameter is mandatory when
	// scaling_policy_type is set to RECURRENCE.When the scaling action is triggered
	// periodically, the end time cannot be earlier than the current and start time.The time
	// format is YYYY-MM-DDThh:mmZ
	EndTime string `json:"end_time,omitempty"`
}

type CreateScalingPolicyAction struct {
	// Specifies the operation to be performed. The default operation
	// is ADD.ADD: adds specified number of instances to the AS group.REMOVE: removes
	// specified number of instances from the AS group.SET: sets the number of instances in
	// the AS group.
	Operation string `json:"operation,omitempty"`

	// Specifies the number of instances to be operated. The default
	// number is 1.Either instance_number or instance_percentage is required.
	InstanceNumber *int `json:"instance_number,omitempty"`

	// Indicates the percentage of instances to be operated. You can
	// increase/decrease or set the number of instances in an AS group to the specified
	// percentage of the current number of instances.If neither instance_number nor
	// instance_percentage is specified, the number of instances to be operated is 1.Either
	// instance_number or instance_percentage is required.
	InstancePercentage *int `json:"instance_percentage,omitempty"`
}

type ScalingPolicyAction struct {
	// Specifies the operation to be performed. The default operation
	// is ADD.ADD: adds specified number of instances to the AS group.REMOVE: removes
	// specified number of instances from the AS group.SET: sets the number of instances in
	// the AS group.
	Operation string `json:"operation,omitempty"`

	// Specifies the number of instances to be operated. The default
	// number is 1.Either instance_number or instance_percentage is required.
	InstanceNumber *int `json:"instance_number,omitempty"`

	// Indicates the percentage of instances to be operated. You can
	// increase/decrease or set the number of instances in an AS group to the specified
	// percentage of the current number of instances.If neither instance_number nor
	// instance_percentage is specified, the number of instances to be operated is 1.Either
	// instance_number or instance_percentage is required.
	InstancePercentage *int `json:"instance_percentage,omitempty"`
}

type ActionOpts struct {
	// Specifies the operation flag for an AS group.execute: executes
	// the AS group.resume: enables the AS group.pause: disables the AS group.
	Action string `json:"action" required:"true"`
}

type ActionOptsBuilder interface {
	ToPoliciesActionMap() (map[string]interface{}, error)
}

func (opts ActionOpts) ToPoliciesActionMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Action(client *gophercloud.ServiceClient, scalingPolicyId string, opts ActionOptsBuilder) (r ActionResult) {
	b, err := opts.ToPoliciesActionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ActionURL(client, scalingPolicyId), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

type CreateOpts struct {
	// Specifies the AS policy name. The name can contain letters,
	// digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingPolicyName string `json:"scaling_policy_name" required:"true"`

	// Specifies the AS group ID. You can obtain its value from the
	// API used to query AS groups. For details, see Querying AS
	// Groups.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063030.html
	ScalingGroupId string `json:"scaling_group_id" required:"true"`

	// Specifies the AS policy type.ALARM (corresponding to alarm_id):
	// indicates that the scaling action is triggered by an alarm.SCHEDULED (corresponding
	// to scheduled_policy): indicates that the scaling action is triggered as
	// scheduled.RECURRENCE (corresponding to scheduled_policy): indicates that the scaling
	// action is triggered periodically.
	ScalingPolicyType string `json:"scaling_policy_type" required:"true"`

	// Specifies the alarm rule ID. This parameter is mandatory when
	// scaling_policy_type is set to ALARM. After this parameter is specified, the value of
	// scheduled_policy does not take effect.After you create an alarm policy, the system
	// automatically adds an alarm triggering activity of the autoscaling type to the
	// alarm_actions field in the alarm rule specified by the parameter value.You can obtain
	// the parameter value by querying the CES alarm rule list. For details, see section
	// Querying Alarms in the Cloud Eye API Reference.
	AlarmId string `json:"alarm_id,omitempty"`

	// Specifies the periodic or scheduled AS policy. This parameter
	// is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE. After this
	// parameter is specified, the value of alarm_id does not take effect.
	ScheduledPolicy ScheduledPolicy `json:"scheduled_policy"`

	// Specifies the action of the AS policy.
	ScalingPolicyAction CreateScalingPolicyAction `json:"scaling_policy_action"`

	// Specifies the cooling duration (in seconds), and is 900 by
	// default.
	CoolDownTime *int `json:"cool_down_time,omitempty"`
}

type CreateOptsBuilder interface {
	ToPoliciesCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToPoliciesCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPoliciesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, scalingPolicyId string) (r DeleteResult) {
	url := DeleteURL(client, scalingPolicyId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

func Get(client *gophercloud.ServiceClient, scalingPolicyId string) (r GetResult) {
	url := GetURL(client, scalingPolicyId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
	// Specifies the AS policy ID.
	ScalingPolicyID string `q:"scaling_policy_id"`
	// Specifies the AS policy name.
	ScalingPolicyName string `q:"scaling_policy_name"`

	// Specifies the AS policy type.
	ScalingPolicyType string `q:"scaling_policy_type"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value
	// is 20 and the maximum value is 100.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, scalingGroupId string, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client, scalingGroupId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := PolicyPage{pagination.NumberPageBase{PageResult: r}}
		p.NumberPageBase.Owner = p
		return p
	})
}

type UpdateOpts struct {
	// Specifies the AS policy name. The name can contain letters,
	// digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingPolicyName string `json:"scaling_policy_name,omitempty"`

	// Specifies the AS policy type.ALARM (corresponding to alarm_id):
	// indicates that the scaling action is triggered by an alarm.SCHEDULED (corresponding
	// to scheduled_policy): indicates that the scaling action is triggered as
	// scheduled.RECURRENCE (corresponding to scheduled_policy): indicates that the scaling
	// action is triggered periodically.
	ScalingPolicyType string `json:"scaling_policy_type,omitempty"`

	// Specifies the alarm rule ID. This parameter is mandatory when
	// scaling_policy_type is set to ALARM. After this parameter is specified, the value of
	// scheduled_policy does not take effect.After you create an alarm policy, the system
	// automatically adds an alarm triggering activity of the autoscaling type to the
	// alarm_actions field in the alarm rule specified by the parameter value.You can obtain
	// the parameter value by querying the CES alarm rule list. For details, see section
	// Querying Alarms in the Cloud Eye API Reference.
	AlarmId string `json:"alarm_id,omitempty"`

	// Specifies the periodic or scheduled AS policy. This parameter
	// is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE. After this
	// parameter is specified, the value of alarm_id does not take effect.
	ScheduledPolicy ScheduledPolicy `json:"scheduled_policy"`

	// Specifies the scaling action of the AS policy.
	ScalingPolicyAction ScalingPolicyAction `json:"scaling_policy_action"`

	// Specifies the cooling duration (in seconds), which is 900 by
	// default.
	CoolDownTime *int `json:"cool_down_time,omitempty"`
}

type UpdateOptsBuilder interface {
	ToPoliciesUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToPoliciesUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, scalingPolicyId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoliciesUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, scalingPolicyId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
