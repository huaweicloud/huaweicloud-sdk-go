package bill

import "github.com/gophercloud/gophercloud"

type QueryPartnerMonthlyBillsOpts struct {
	//Customer Id
	CustomerId string `q:"customerId,omitempty"`

	//Cloud service type code
	CloudServiceTypeCode string `q:"cloud_service_type_code,omitempty"`

	//period
	Period string `q:"period,omitempty"`

	//Payment method
	PayMethod string `q:"payMethod,omitempty"`

	//Page to be queried
	Offset *int `q:"offset,omitempty"`

	//Number of records on each page
	Limit *int `q:"limit,omitempty"`

	//bill type
	BillType string `q:"bill_type,omitempty"`

}

type QueryMonthlyExpenditureSummaryOpts struct {
	//Query cycle
	Cycle string `q:"cycle" required:"true"`

	//Cloud service type code
	CloudServiceTypeCode string `q:"cloud_service_type_code,omitempty"`

	//Account type
	Type string `q:"type,omitempty"`

	//Enterprise project ID.
	EnterpriseProjectId string `q:"enterpriseProjectId,omitempty"`
}

type QueryResourceUsageDetailsOpts struct {
	//Expenditure month
	Cycle string `q:"cycle" required:"true"`

	//Cloud service type code
	CloudServiceTypeCode string `q:"cloudServiceTypeCode,omitempty"`

	//Resource type code
	ResourceTypeCode string `q:"resourceTypeCode,omitempty"`

	//Cloud service region code
	RegionCode string `q:"regionCode,omitempty"`

	//Resource instance ID
	ResInstanceId string `q:"resInstanceId,omitempty"`

	//Payment method
	PayMethod string `q:"payMethod" required:"true"`

	//Enterprise project ID
	EnterpriseProjectId string `q:"enterpriseProjectId,omitempty"`

	//Page number.
	Offset *int `q:"offset" required:"true"`

	//Indicates the page limit
	Limit *int `q:"limit" required:"true"`
}

type QueryResourceUsageRecordOpts struct {
	//Start time.
	StartTime string `q:"startTime" required:"true"`

	//End time
	EndTime string `q:"endTime" required:"true"`

	//Cloud service type code
	CloudServiceTypeCode string `q:"cloudServiceTypeCode,omitempty"`

	//Cloud service region code
	RegionCode string `q:"regionCode,omitempty"`

	//Order ID.
	OrderId string `q:"orderId,omitempty"`

	//Payment method
	PayMethod string `q:"payMethod" required:"true"`

	//Page number
	Offset int `q:"offset,omitempty"`

	//Number of records per page
	Limit int `q:"limit,omitempty"`

	//Resource ID.
	ResourceId string `q:"resourceId,omitempty"`

	//Enterprise project ID
	EnterpriseProjectId string `q:"enterpriseProjectId,omitempty"`
}

type QueryPartnerMonthlyBillsOptsBuilder interface {
	ToQueryPartnerMonthlyBillsMap() (string, error)
}

func (opts QueryPartnerMonthlyBillsOpts) ToQueryPartnerMonthlyBillsMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type QueryMonthlyExpenditureSummaryOptsBuilder interface {
	ToQueryMonthlyExpenditureSummaryMap() (string, error)
}

func (opts QueryMonthlyExpenditureSummaryOpts) ToQueryMonthlyExpenditureSummaryMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type QueryResourceUsageDetailsOptsBuilder interface {
	ToQueryResourceUsageDetailsMap() (string, error)
}

func (opts QueryResourceUsageDetailsOpts) ToQueryResourceUsageDetailsMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type QueryResourceUsageRecordOptsBuilder interface {
	ToQueryResourceUsageRecordMap() (string, error)
}

func (opts QueryResourceUsageRecordOpts) ToQueryResourceUsageRecordMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

/**
 * This API is used to query monthly bills.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func QueryPartnerMonthlyBills(client *gophercloud.ServiceClient, opts QueryPartnerMonthlyBillsOptsBuilder)(r QueryPartnerMonthlyBillsResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryPartnerMonthlyBillsURL(client,domainID)
	if opts != nil {
		query, err := opts.ToQueryPartnerMonthlyBillsMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}

/**
 * This API can be used to query the expenditure summary bills of a customer on the customer platform. The bills summarize the summary data by month. The expenditure summary contains only the data generated before 24:00 of the previous day.
 * This API can be invoked using the customer AK/SK or token only.
 */
func QueryMonthlyExpenditureSummary(client *gophercloud.ServiceClient, opts QueryMonthlyExpenditureSummaryOptsBuilder)(r QueryMonthlyExpenditureSummaryResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryMonthlyExpenditureSummaryURL(client,domainID)
	if opts != nil {
		query, err := opts.ToQueryMonthlyExpenditureSummaryMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}

/**
 * This API can be used to query usage details of each resource for a customer on the customer platform. The resource details have a latency (a maximum of 24 hours).
 * This API can be invoked using the customer AK/SK or token only.
 */
func QueryResourceUsageDetails(client *gophercloud.ServiceClient, opts QueryResourceUsageDetailsOptsBuilder)(r QueryResourceUsageDetailsResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryResourceUsageDetailsURL(client,domainID)
	if opts != nil {
		query, err := opts.ToQueryResourceUsageDetailsMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}

/**
 * This API can be used to query the usage details of each resource for a customer on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func QueryResourceUsageRecord(client *gophercloud.ServiceClient, opts QueryResourceUsageRecordOptsBuilder)(r QueryResourceUsageRecordResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryResourceUsageRecordURL(client,domainID)
	if opts != nil {
		query, err := opts.ToQueryResourceUsageRecordMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}