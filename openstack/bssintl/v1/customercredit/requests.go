package customercredit

import (
	"encoding/json"
	"github.com/gophercloud/gophercloud"
)

type QueryCreditOpts struct {
	//Customer ID
	CustomerId string `q:"customer_id" required:"true"`
}

type SetCreditOpts struct {
	//客户ID。
	CustomerId string `json:"customerId" required:"true"`

	//调整的目标金额。
	AdjustmentAmount float64 `json:"adjustmentAmount" required:"true"`

	//金额单位。
	MeasureId *int `json:"measureId" required:"true"`
}

type QueryCreditOptsBuilder interface {
	ToQueryCreditMap() (string, error)
}

func (opts QueryCreditOpts) ToQueryCreditMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type SetCreditOptsBuilder interface {
	ToSetCreditMap() (map[string]interface{}, error)
}

func (opts SetCreditOpts) ToSetCreditMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
* You can use REST APIs to configure customers' credit limits on your sales platform.
* This API can be invoked only by the partner account AK/SK or token.
 */
func SetCredit(client *gophercloud.ServiceClient, opts SetCreditOptsBuilder) (r SetCreditResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToSetCreditMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getSetCreditURL(client, domainID), body, nil, &gophercloud.RequestOpts{
			OkCodes: []int{204},
			HandleError: func(httpStatus int, responseContent string) error {
				var setCrediterr SetCreditError
				message := responseContent
				err := json.Unmarshal([]byte(responseContent), &setCrediterr)
				if err == nil {
					return &setCrediterr
				}
				return &gophercloud.UnifiedError{
					ErrCode:    gophercloud.MatchErrorCode(httpStatus, message),
					ErrMessage: message,
				}
			},
		})
	}

	return
}

/**
 * You can use REST APIs to query customers' credit limits on your sales platform.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func QueryCredit(client *gophercloud.ServiceClient, opts QueryCreditOptsBuilder)(r QueryCreditResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryCreditURL(client,domainID)
	if opts != nil {
		query, err := opts.ToQueryCreditMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}