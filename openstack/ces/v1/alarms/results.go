package alarms

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateAlarm struct {
	AlarmId string `json:"alarm_id,omitempty"`
}

type Alarms struct {
	MetricAlarms []MetricAlarms `json:"metric_alarms"`
	MetaData     MetaData       `json:"meta_data"`
}

type Alarm struct {
	MetricAlarms []MetricAlarms `json:"metric_alarms"`
}

// 查询结果元数据信息，包括分页信息等。
type MetaData struct {
	// 当前返回结果条数。
	Count int `json:"count"`
	// 总条数。
	Total int `json:"total"`
	// 下一个开始的标记，用于分页。
	Marker string `json:"marker"`
}

type MetricAlarms struct {
	// 告警名称。
	AlarmName string `json:"alarm_name"`
	// 告警描述。
	AlarmDescription string `json:"alarm_description,omitempty"`
	// 告警指标。
	Metric MetricInfo `json:"metric"`
	// 告警触发条件。
	Condition Condition `json:"condition"`
	// 是否启用该条告警。
	AlarmEnabled bool `json:"alarm_enabled,omitempty"`
	// 告警级别，默认为2，级别为1、2、3、4。分别对应紧急、重要、次要、提示。
	AlarmLevel int `json:"alarm_level,omitempty"`
	// 告警类型。仅针对事件告警的参数，枚举类型：EVENT.SYS或者EVENT.CUSTOM
	AlarmType string `json:"alarm_type,omitempty"`
	// 是否启用该条告警触发的动作。
	AlarmActionEnabled bool `json:"alarm_action_enabled,omitempty"`
	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notificationList：告警状态发生变化时，被通知对象的列表。
	AlarmActions []Actions `json:"alarm_actions,omitempty"`
	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值：  notification：通知。  notificationList：告警状态发生变化时，被通知对象的列表。
	OkActions []Actions `json:"ok_actions,omitempty"`
	// 数据不足触发的动作（该参数已废弃，建议无需配置）。
	InsufficientdataActions []Actions `json:"insufficientdata_actions,omitempty"`
	// 告警规则的ID。
	AlarmId string `json:"alarm_id"`
	// 告警状态变更的时间，UNIX时间戳，单位毫秒。
	UpdateTime int64 `json:"update_time"`
	// 告警状态，取值说明：  ok，正常 alarm，告警 insufficient_data，数据不足
	AlarmState string `json:"alarm_state"`
}

type CreateResult struct {
	gophercloud.Result
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	gophercloud.Result
}

type UpdateResult struct {
	gophercloud.ErrResult
}

func (r CreateResult) Extract() (*CreateAlarm, error) {
	var s *CreateAlarm
	err := r.ExtractInto(&s)
	return s, err
}

func ExtractAlarms(r pagination.Page) (Alarms, error) {
	var s Alarms
	s.MetricAlarms = make([]MetricAlarms, 0)
	err := (r.(AlarmsPage)).ExtractInto(&s)
	if s.MetaData.Count == 0 {
		s.MetaData.Count = len(s.MetricAlarms)
		s.MetaData.Total = len(s.MetricAlarms)
	}
	return s, err
}

type AlarmsPage struct {
	pagination.LinkedPageBase
}

func (r GetResult) Extract() (*Alarm, error) {
	var s *Alarm
	err := r.ExtractInto(&s)
	return s, err
}

func (r AlarmsPage) NextPageURL() (string, error) {
	alarms, err := ExtractAlarms(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(alarms.MetricAlarms[len(alarms.MetricAlarms)-1].AlarmId)
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r AlarmsPage) IsEmpty() (bool, error) {
	s, err := ExtractAlarms(r)
	return len(s.MetricAlarms) == 0, err
}

/*
ExtractNextURL is an internal function useful for packages of collection
resources that are paginated in a certain way.

It attempts to extract the "start" URL from slice of Link structs, or
"" if no such URL is present.
*/
func (r AlarmsPage) WrapNextPageURL(start string) (string, error) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		return "", nil
	}
	uq := r.URL.Query()
	uq.Set("start", start)
	r.URL.RawQuery = uq.Encode()
	return r.URL.String(), nil
}
