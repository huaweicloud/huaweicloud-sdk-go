package alarms

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOpts struct {
	// 取值范围(0,100]，默认值为100  用于限制结果数据条数。
	Limit int `q:"limit"`
	// 用于标识结果排序方法。  取值说明，默认值为desc。  asc：升序 desc：降序
	Order string `q:"order"`
	// 分页起始值，内容为alarm_id。
	Start string `q:"start"`
}

type Actions struct {
	Type             string   `json:"type"  required:"true"`
	NotificationList []string `json:"notificationList"  required:"true"`
}

type Condition struct {
	ComparisonOperator string `json:"comparison_operator"  required:"true"`
	Count              int    `json:"count"  required:"true"`
	Filter             string `json:"filter"  required:"true"`
	Period             int    `json:"period"  required:"true"`
	Unit               string `json:"unit,omitempty"`
	Value              int    `json:"value"  required:"true"`
	SuppressDuration   int    `json:"suppress_duration,omitempty"`
}

type CreateOpts struct {
	// 告警名称，只能包含0-9/a-z/A-Z/_/-或汉字。
	AlarmName        string `json:"alarm_name" required:"true"`
	AlarmDescription string `json:"alarm_description,omitempty"`
	// 告警指标。
	Metric MetricInfo `json:"metric" required:"true"`
	// 告警触发条件。
	Condition Condition `json:"condition"  required:"true"`
	// 是否启用该条告警，默认为true。
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`
	// 是否启用该条告警触发的动作，默认为true。注：若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notificationList值保持一致。
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// 告警级别，默认为2，级别为1、2、3、4。分别对应紧急、重要、次要、提示。
	AlarmLevel int `json:"alarm_level,omitempty"`
	// 告警类型。仅针对事件告警的参数，枚举类型：EVENT.SYS或者EVENT.CUSTOM
	AlarmType string `json:"alarm_type,omitempty"`
	// 告警触发的动作。 结构样例如下： { \"type\": \"notification\",\"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"] } type取值：notification：通知。 autoscaling：弹性伸缩。
	AlarmActions []Actions `json:"alarm_actions,omitempty"`
	// 数据不足触发的动作（该参数已废弃，建议无需配置）。
	InsufficientdataActions []Actions `json:"insufficientdata_actions,omitempty"`
	//告警恢复触发的动作。
	OkActions []Actions `json:"ok_actions,omitempty"`
}

// 指标信息
type MetricInfo struct {
	// 指标维度
	Dimensions []MetricsDimension `json:"dimensions" required:"true"`
	// 指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。  具体指标名请参见查询指标列表中查询出的指标名。
	MetricName string `json:"metric_name" required:"true"`
	// 指标命名空间，，例如弹性云服务器命名空间。格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空。
	Namespace string `json:"namespace" required:"true"`
}

// 指标维度
type MetricsDimension struct {
	// 维度名
	Name string `json:"name,omitempty"`
	// 维度值
	Value string `json:"value,omitempty"`
}

type UpdateOpts struct {
	// 告警是否启用。true：启动。false：停止
	AlarmEnabled *bool `json:"alarm_enabled" required:"true"`
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return
}

func Delete(client *gophercloud.ServiceClient, alarmId string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, alarmId), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

func Get(client *gophercloud.ServiceClient, alarmId string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, alarmId), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

func List(client *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	url := listURL(client) + q.String()
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return AlarmsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Update(client *gophercloud.ServiceClient, alarmId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updateURL(client, alarmId), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
