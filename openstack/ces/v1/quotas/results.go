package quotas

import (
	"github.com/gophercloud/gophercloud"
)

// This is a auto create Response Object
type Quota struct {
	Quotas Quotas `json:"quotas"`
}

type Quotas struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	// 配额总数。
	Quota int `json:"quota"`
	// 配额类型。  枚举值说明：  alarm，告警规则
	Type string `json:"type"`
	// 单位。
	Unit string `json:"unit"`
	// 已使用配额数。
	Used int `json:"used"`
}

type GetResult struct {
	gophercloud.Result
}

func (r GetResult) Extract() (*Quota, error) {
	var s *Quota
	err := r.ExtractInto(&s)
	return s, err
}
