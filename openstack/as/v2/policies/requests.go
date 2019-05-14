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

	// Specifies the number of instances or the bandwidth. The default
	// value is 1.If scaling_resource_type is set to SCALING_GROUP, the value of this
	// parameter is the number of instances.If scaling_resource_type is set to BANDWIDTH,
	// the value of this parameter is the bandwidth (Mbit/s).Either size or percentage is
	// required.
	Size *int `json:"size,omitempty"`

	// Specifies the percentage of instances to be operated.If neither
	// instance_number nor instance_percentage is specified, the number of instances to be
	// operated is 1.Either size or percentage is required.
	Percentage *int `json:"percentage,omitempty"`

	// Specifies the operation restrictions.If scaling_resource_type
	// is set to BANDWIDTH, this parameter takes effect and the unit is Mbit/s.In this
	// case:If operation is set to ADD, this parameter indicates the maximum bandwidth.If
	// operation is set to REDUCE, this parameter indicates the minimum bandwidth.
	Limits *int `json:"limits,omitempty"`
}

type ScalingPolicyAction struct {
	// Specifies the operation to be performed. The default operation
	// is ADD.ADD: adds specified number of instances to the AS group.REMOVE: removes
	// specified number of instances from the AS group.SET: sets the number of instances in
	// the AS group.
	Operation string `json:"operation,omitempty"`

	// Specifies the number of instances or the bandwidth. The default
	// value is 1.If scaling_resource_type is set to SCALING_GROUP, the value of this
	// parameter is the number of instances.If scaling_resource_type is set to BANDWIDTH,
	// the value of this parameter is the bandwidth (Mbit/s).Either size or percentage is
	// required.
	Size int `json:"size,omitempty"`

	// Specifies the percentage of instances to be operated.If neither
	// instance_number nor instance_percentage is specified, the number of instances to be
	// operated is 1.Either size or percentage is required.
	Percentage int `json:"percentage,omitempty"`

	// Specifies the operation restrictions.If scaling_resource_type
	// is set to BANDWIDTH, this parameter takes effect and the unit is Mbit/s.In this
	// case:If operation is set to ADD, this parameter indicates the maximum bandwidth.If
	// operation is set to REDUCE, this parameter indicates the minimum bandwidth.
	Limits int `json:"limits,omitempty"`
}

type CreateOpts struct {
	// Specifies the AS policy name. The name can contain letters,
	// digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingPolicyName string `json:"scaling_policy_name" required:"true"`

	// Specifies the scaling resource ID, which is the unique ID of an
	// AS group or bandwidth.
	ScalingResourceId string `json:"scaling_resource_id" required:"true"`

	// Specifies the scaling resource type.AS group:
	// SCALING_GROUP.Bandwidth: BANDWIDTH
	ScalingResourceType string `json:"scaling_resource_type" required:"true"`

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

	// Specifies the scaling action of the AS policy.
	ScalingPolicyAction CreateScalingPolicyAction `json:"scaling_policy_action"`

	// Specifies the cooling duration (in seconds), which is 900 by
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

func Get(client *gophercloud.ServiceClient, scalingPolicyId string) (r GetResult) {
	url := GetURL(client, scalingPolicyId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ResourceListOpts struct {
	// Specifies the AS policy name.
	ScalingPolicyName string `q:"scaling_policy_name"`

	// Specifies the AS policy type.
	ScalingPolicyType string `q:"scaling_policy_type"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value
	// is 20 and the maximum value is 100.
	Limit int `q:"limit"`

	// Specifies the AS policy id.
	ScalingPolicyID string `q:"scaling_policy_id"`
}

type ResourceListOptsBuilder interface {
	ToResourceListQuery() (string, error)
}

func (opts ResourceListOpts) ToResourceListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func GetPolicyListByResourceID(client *gophercloud.ServiceClient, scalingResourceId string, opts ResourceListOptsBuilder) pagination.Pager {
	url := ListURL(client, scalingResourceId)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
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

type ListOpts struct {
	// Specifies the AS policy name.
	ScalingPolicyName string `q:"scaling_policy_name"`

	// Specifies the AS policy type.
	ScalingPolicyType string `q:"scaling_policy_type"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value
	// is 20 and the maximum value is 100.
	Limit int `q:"limit"`

	// Specifies the AS policy id.
	ScalingPolicyID string `q:"scaling_policy_id"`

	ScalingResourceID   string `q:"scaling_resource_id"`
	ScalingResourceType string `q:"scaling_resource_type"`
	SortBy              string `q:"sort_by"`
	Order               string `q:"order"`
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := CreateURL(client)
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

	// Specifies the scaling resource ID, which is the unique ID of an
	// AS group or bandwidth.
	ScalingResourceId string `json:"scaling_resource_id,omitempty"`

	// Specifies the scaling resource type.AS group:
	// SCALING_GROUP.Bandwidth: BANDWIDTH
	ScalingResourceType string `json:"scaling_resource_type,omitempty"`

	//
	AlarmId string `json:"alarm_id,omitempty"`

	// Specifies the alarm rule ID. This parameter is mandatory when
	// scaling_policy_type is set to ALARM. After this parameter is specified, the value of
	// scheduled_policy does not take effect.After you create an alarm policy, the system
	// automatically adds an alarm triggering activity of the autoscaling type to the
	// alarm_actions field in the alarm rule specified by the parameter value.You can obtain
	// the parameter value by querying the CES alarm rule list. For details, see section
	// Querying Alarms in the Cloud Eye API Reference.
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
