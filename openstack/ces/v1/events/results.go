package events

import (
	"github.com/gophercloud/gophercloud"
)

type Event struct {
	// 事件ID。
	EventId string `json:"event_id"`
	// 事件名称。  必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。
	EventName string `json:"event_name"`
}

type Events []Event

type CreateResult struct {
	gophercloud.Result
}

 func (r CreateResult) Extract() (*Events, error) {
 	var s *Events
 	err := r.ExtractInto(&s)
 	return s, err
 }
