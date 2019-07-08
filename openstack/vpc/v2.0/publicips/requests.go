package publicips

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOpts struct {
	PublicIP    PublicIP    `json:"publicip" required:"true"`
	Bandwidth   Bandwidth   `json:"bandwidth" required:"true"`
	ExtendParam ExtendParam `json:"extendParam,omitempty"`
}

type PublicIP struct {
	Type      string `json:"type" required:"true"`
	IPVersion int   `json:"ip_version,omitempty"`
}

type Bandwidth struct {
	Name       string `json:"name,omitempty"`
	Size       int   `json:"size,omitempty"`
	ID         string `json:"id,omitempty"`
	ShareType  string `json:"share_type" required:"true"`
	ChargeMode string `json:"charge_mode,omitempty"`
}

type ExtendParam struct {
	ChargeMode  string `json:"charge_mode,omitempty"`
	PeriodType  string `json:"period_type,omitempty"`
	PeriodNum   int   `json:"period_num,omitempty"`
	IsAutoRenew string `json:"is_auto_renew,omitempty"`
	IsAutoPay   string `json:"is_auto_pay,omitempty"`
}

func (opts CreateOpts) ToPublicIPCreateMap() (map[string]interface{}, error) {

	//Marshal opts as map[string]interface{}
	return gophercloud.BuildRequestBody(opts, "")
}

func Create(c *gophercloud.ServiceClient, opts CreateOpts) (interface{}, error) {
	var r CreateResult
	body, err := opts.ToPublicIPCreateMap()
	if err != nil {
		return nil, err
	}

	_, r.Err = c.Post(CreateURL(c), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})

	onDemandData, onDemandErr := r.extractOnDemand()
	orderData, orderErr := r.extract()

	if orderData.OrderID != "" {
		return orderData, orderErr
	}

	return onDemandData, onDemandErr

}
