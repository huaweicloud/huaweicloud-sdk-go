package enquiry

import "github.com/gophercloud/gophercloud"

type QueryRatingOpts struct {
	//Project ID.
	TenantId string `json:"tenantId" required:"true"`

	//Region ID
	RegionId string `json:"regionId" required:"true"`

	//AZ ID.
	AvaliableZoneId string `json:"avaliableZoneId"`

	//Billing mode
	ChargingMode *int `json:"chargingMode" required:"true"`

	//Order period type
	PeriodType int `json:"periodType"`

	//Number of subscription periods.
	PeriodNum int `json:"periodNum"`

	//Expiration date.
	PeriodEndDate string `json:"periodEndDate"`

	//Associated resource ID
	RelativeResourceId string `json:"relativeResourceId"`

	//Period type of the associated resource
	RelativeResourcePeriodType int `json:"relativeResourcePeriodType"`

	//Number of subscriptions
	SubscriptionNum int `json:"subscriptionNum"`

	//Product information
	ProductInfo []ProductInfo `json:"productInfos"`

	//Inquiry date.
	InquiryTime string `json:"inquiryTime"`
}

type ProductInfo struct {
	//ID.
	Id string `json:"id" required:"true"`

	//Cloud service type code
	CloudServiceType string `json:"cloudServiceType" required:"true"`

	//Resource type code
	ResourceType string `json:"resourceType" required:"true"`

	//Resource Spec Code
	ResourceSpecCode string `json:"resourceSpecCode" required:"true"`

	//Resource capacity, which is used together with resouceSizeMeasureId.
	ResourceSize int `json:"resourceSize"`

	//Resource capacity measurement ID
	ResouceSizeMeasureId int `json:"resouceSizeMeasureId"`

	//Usage value
	UsageFactor string `json:"usageFactor"`

	//Usage value
	UsageValue float64 `json:"usageValue"`

	//Usage measurement ID
	UsageMeasureId int `json:"usageMeasureId"`

	//Extended parameter, optional
	ExtendParams string `json:"extendParams"`
}

type QueryRatingOptsBuilder interface {
	ToQueryRatingOptsMap() (map[string]interface{}, error)
}

func (opts QueryRatingOpts) ToQueryRatingOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
 * The customer platform obtains the product prices on the HUAWEI CLOUD official website based on the product catalog
 * This API can be invoked using the customer token, or the partner's AK/SK or token.
 */
func QueryRating(client *gophercloud.ServiceClient, opts QueryRatingOptsBuilder) (r QueryRatingResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToQueryRatingOptsMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getQueryRatingURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}
