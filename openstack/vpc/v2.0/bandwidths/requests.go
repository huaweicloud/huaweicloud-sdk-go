package bandwidths

import (
	"github.com/gophercloud/gophercloud"
)

type UpdateOpts struct {
	Bandwidth   Bandwidth    `json:"bandwidth" required:"true"`
	ExtendParam *ExtendParam `json:"extendParam,omitempty"`
}
type Bandwidth struct {
	Name string `json:"name,omitempty"`
	Size int   `json:"size,omitempty"`
}
type ExtendParam struct {
	IsAutoPay string `json:"is_auto_pay,omitempty"`
}

func (opts UpdateOpts) ToBandWidthUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func Update(c *gophercloud.ServiceClient, bandwidthID string, opts UpdateOpts) (interface{}, error) {
	// only update prepaid bandwidth
	var size int
	var r UpdateResult
	body, err := opts.ToBandWidthUpdateMap()
	if err != nil {
		return nil, err
	}

	_, r.Err = c.Put(UpdateURL(c, bandwidthID), body, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})

	if opts.Bandwidth.Size == size ||opts.Bandwidth.Size == 0{
		// extract bandwidth
		return r.Extract()
	}
	//extract order id
	return r.ExtractOrderID()

}
