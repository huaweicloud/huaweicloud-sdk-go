package publicips

import (
	"github.com/gophercloud/gophercloud"
)

type PrePaid struct {
	OrderID    string `json:"order_id"`
	PublicipID string `json:"publicip_id"`
}

type PostPaid struct {
	ID                string `json:"id"`
	Status            string `json:"status"`
	Type              string `json:"type"`
	PublicIPAddress   string `json:"public_ip_address"`
	PublicIPv6Address string `json:"public_ipv6_address"`
	TenantID          string `json:"tenant_id"`
	CreateTime        string `json:"create_time"`
	BandwidthSize     int    `json:"bandwidth_size"`
	IPVersion         int    `json:"ip_version"`
}

type CreateResult struct {
	gophercloud.Result
}

func (r CreateResult) extractOnDemand() (PostPaid, error) {
	var s struct {
		Publicip PostPaid `json:"publicip"`
	}
	err := r.ExtractInto(&s)
	return s.Publicip, err

}

func (r CreateResult) extract() (PrePaid, error) {
	var s PrePaid
	err := r.ExtractInto(&s)
	return s, err
}
