package bandwidths

import (
	"github.com/gophercloud/gophercloud"
)

type PrePaid struct {
	OrderID string `json:"order_id"`
}

type PostPaid struct {
	Name                string         `json:"name"`
	Size                int           `json:"size"`
	ID                  string         `json:"id"`
	ShareType           string         `json:"share_type"`
	ChargeMode          string         `json:"charge_mode"`
	BandwidthType       string         `json:"bandwidth_type"`
	TenantID            string         `json:"tenant_id"`
	PublicipInfo        []PublicipInfo `json:"publicip_info"`
	//EnterpriseProjectID string         `json:"enterprise_project_id"`
	//BillingInfo         string         `json:"billing_info"`
}

type PublicipInfo struct {
	PublicipID        string `json:"publicip_id"`
	PublicIPAddress   string `json:"publicip_address"`
	Publicipv6Address string `json:"publicipv_6_address"`
	IpVersion         int    `json:"ip_version"`
	PublicipType      string `json:"publicip_type"`
}

type UpdateResult struct {
	gophercloud.Result
}

func (r UpdateResult) ExtractOrderID() (PrePaid, error) {
	var s PrePaid
	err := r.ExtractInto(&s)
	return s, err
}

func (r UpdateResult) Extract() (PostPaid, error) {
	var s struct {
		Bandwidth PostPaid `json:"bandwidth"`
	}
	err := r.ExtractInto(&s)
	return s.Bandwidth, err
}
