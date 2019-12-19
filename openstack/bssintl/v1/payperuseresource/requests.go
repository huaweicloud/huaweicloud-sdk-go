package payperuseresource

import "github.com/gophercloud/gophercloud"

type QueryCustomerResourceOpts struct {
	//Customer resource ID.
	CustomerResourceId string `json:"customerResourceId,omitempty"`

	//Customer ID.
	CustomerId string `json:"customerId" required:"true"`

	//Cloud service region code
	RegionCode string `json:"regionCode,omitempty"`

	//Cloud service type code
	CloudServiceTypeCode string `json:"cloudServiceTypeCode,omitempty"`

	//Resource type code
	ResourceTypeCode string `json:"resourceTypeCode,omitempty"`

	//Queries resource IDs in batches
	ResourceIds []string `json:"resourceIds,omitempty"`

	//Resource instance name
	ResourceName string `json:"resourceName,omitempty"`

	//Start time of the validity period
	StartTimeBegin string `json:"startTimeBegin,omitempty"`

	//End time of the validity period
	StartTimeEnd string `json:"startTimeEnd,omitempty"`

	//Current page
	PageNo int `json:"pageNo,omitempty"`

	//Number of records displayed on each page
	PageSize int `json:"pageSize,omitempty"`
}

type QueryCustomerResourceOptsBuilder interface {
	ToQueryCustomerResourceMap() (map[string]interface{}, error)
}

func (opts QueryCustomerResourceOpts) ToQueryCustomerResourceMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
 * A customer can query its pay-per-use resources on the partner sales platform. The on-demand resource data has a latency, and the latency for each cloud service data varies. The data obtained using this API is for reference only.
 * This API can be invoked using the partner AK/SK or token only.
 */
func QueryCustomerResource(client *gophercloud.ServiceClient, opts QueryCustomerResourceOptsBuilder) (r QueryCustomerResourceResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToQueryCustomerResourceMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getQueryCustomerResourceURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}