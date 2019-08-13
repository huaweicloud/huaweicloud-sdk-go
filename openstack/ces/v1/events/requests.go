package events

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOpts []EventItem

type EventItem struct {
	// 事件名称。  必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。
	EventName string `json:"event_name" required:"true"`
	// 事件来源。  格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。
	EventSource string `json:"event_source,omitempty"`
	// 事件发生时间。UNIX时间戳，单位毫秒。  说明： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-1小时+20秒，当前时间+10分钟-20秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。
	Time int64 `json:"time" required:"true"`
	//  事件详情。
	Detail EventItemDetail `json:"detail" required:"true"`
}

type EventItemDetail struct {
	// 事件内容，最大长度4096。
	Content string `json:"content,omitempty"`
	// 所属分组。  资源分组对应的ID，必须传存在的分组ID。
	GroupId string `json:"group_id,omitempty"`
	// 资源ID，支持字母、数字_ -：，最大长度128。
	ResourceId string `json:"resource_id,omitempty"`
	// 资源名称，支持字母 中文 数字_ -. ，最大长度128。
	ResourceName string `json:"resource_name,omitempty"`
	// 事件状态。  枚举类型：normal\\warning\\incident
	EventState string `json:"event_state,omitempty"`
	// 事件级别。  枚举类型：Critical, Major, Minor, Info
	EventLevel string `json:"event_level,omitempty"`
	// 事件用户。  支持字母 数字_ -/空格 ，最大长度64。
	EventUser string `json:"event_user,omitempty"`
}

func (opts EventItem) ToMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func (opts CreateOpts) ToCreateMap() ([]map[string]interface{}, error) {
	newOpts := make([]map[string]interface{}, len(opts))
	for i, opt := range opts {
		opt, err := opt.ToMap()
		if err != nil {
			return nil, err
		}
		newOpts[i] = opt
	}
	return newOpts, nil
}

type CreateOptsBuilder interface {
	ToCreateMap() ([]map[string]interface{}, error)
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
