package account

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"math"
)

type ResourceDailyOptsBuilder interface {
	ToResourcesDailyMap() (map[string]interface{}, error)
}

// ResourceDailyOpts represents options used to get resources.
type ResourceDailyOpts struct {
	//Starting point of consumption time
	StartTime string `json:"startTime" required:"true"`

	//End of consumption time
	EndTime string `json:"endTime" required:"true"`

	//Payment method. 0: package cycle 1: on demand
	PayMethod string `json:"payMethod" required:"true"`

	//Cloud Service Type。
	CloudServiceType string `json:"cloudServiceType,omitempty"`

	//Region Code。
	RegionCode string `json:"regionCode,omitempty"`

	//ResourceId
	ResourceId string `json:"resourceId,omitempty"`

	//EnterpriseProjectId
	EnterpriseProjectId string `json:"enterpriseProjectId,omitempty"`

	// Page No
	PageNo int `json:"pageNo,omitempty"`

	//Page Size。
	PageSize int `json:"pageSize,omitempty"`
}

/// ToResourcesDailyMap assembles a request body based on the contents of a
// ResourceDailyOpts.
func (opts ResourceDailyOpts) ToResourcesDailyMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}


func getResourceDaily(client *gophercloud.ServiceClient, opts ResourceDailyOptsBuilder) (r commonResult) {
	domainID := client.ProviderClient.DomainID
	url := getURL(client, domainID)
	b, err := opts.ToResourcesDailyMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(url, b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListResourceDaily(client *gophercloud.ServiceClient, opts ResourceDailyOpts) (*ResourceDaily, error) {
	var allRes ResourceDaily

	var  reqTmp ResourceDailyOpts
	reqTmp =  opts

	reqTmp.PageSize = 1
	reqTmp.PageNo = 1

	rspTmp,err := getResourceDaily(client, reqTmp).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return &allRes, err
	}
	allRes.TotalRecord = rspTmp.TotalRecord
	allRes.Currency = rspTmp.Currency
	allRes.TotalAmount = rspTmp.TotalAmount
	allRes.MeasureId = rspTmp.MeasureId
	allRes.ErrorCode = rspTmp.ErrorCode
	allRes.ErrorMsg = rspTmp.ErrorMsg
	//once query 2000
	onceCnt := 2000
	totalCnt := rspTmp.TotalRecord
	queryTimes := int(math.Ceil(float64(totalCnt) / float64(onceCnt)))

	reqTmp.PageSize = onceCnt

	for i := 1; i <= queryTimes; i++ {
		reqTmp.PageNo = i
		res, err := getResourceDaily(client,reqTmp).Extract()
		if err != nil {
			fmt.Println("err:", err)
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return &allRes, gophercloud.NewSystemCommonError("CBC.0999", "other error")
		}

		allRes.DailyRecords = append(allRes.DailyRecords, res.DailyRecords...)
	}
	return &allRes, nil

}