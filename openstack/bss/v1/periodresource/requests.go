package periodresource

import (
	"github.com/gophercloud/gophercloud"
)

type QueryCustomerPeriodResourcesListOpts struct {
	//Resource IDs.
	ResourceIds string `q:"resource_ids,omitempty"`

	//Order ID.
	OrderId string `q:"order_id,omitempty"`

	//Whether to query only primary resources.
	OnlyMainResource *int `q:"only_main_resource,omitempty"`

	//Resource status
	StatusList string `q:"status_list,omitempty"`

	//Page number.
	PageNo *int `q:"page_no,omitempty"`

	//Number of records per page.
	PageSize *int `q:"page_size,omitempty"`
}

type QueryCustomerPeriodResourcesListOptsBuilder interface {
	ToQueryCustomerPeriodResourcesListMap() (string, error)
}

func (opts QueryCustomerPeriodResourcesListOpts) ToQueryCustomerPeriodResourcesListMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type RenewSubscriptionByResourceIdOpts struct {
	//Resource IDs.
	ResourceIds []string `json:"resource_ids"  required:"true"`

	//Period type.
	PeriodType *int `json:"period_type" required:"true"`

	//Number of periods
	PeriodNum *int `json:"period_num" required:"true"`

	//Expiration policy.
	ExpireMode *int `json:"expire_mode" required:"true"`

	//Whether enable automatic payment.
	IsAutoPay *int `json:"isAutoPay,omitempty"`

}

type RenewSubscriptionByResourceIdOptsBuilder interface {
	ToRenewSubscriptionByResourceIdMap() (map[string]interface{}, error)
}

func (opts RenewSubscriptionByResourceIdOpts) ToRenewSubscriptionByResourceIdMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type UnsubscribeByResourceIdOpts struct {
	//Resource IDs.
	ResourceIds []string `json:"resourceIds" required:"true"`

	//Unsubscription type.
	UnSubType *int `json:"unSubType" required:"true"`

	//Unsubscription cause
	UnsubscribeReasonType *int `json:"unsubscribeReasonType,omitempty"`

	//Unsubscription reason, which is generally specified by the customer.
	UnsubscribeReason string `json:"unsubscribeReason,omitempty"`
}

type UnsubscribeByResourceIdOptsBuilder interface {
	ToUnsubscribeByResourceIdMap() (map[string]interface{}, error)
}

func (opts UnsubscribeByResourceIdOpts) ToUnsubscribeByResourceIdMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type EnableAutoRenewOpts struct {
	//Operation ID.
	ActionId string `json:"action_id" required:"true"`
}

type EnableAutoRenewOptsBuilder interface {
	ToEnableAutoRenewMap() (map[string]interface{}, error)
}

func (opts EnableAutoRenewOpts) ToEnableAutoRenewMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type DisableAutoRenewOpts struct {
	//Operation ID.
	ActionId string `json:"action_id" required:"true"`
}

type DisableAutoRenewOptsBuilder interface {
	ToDisableAutoRenewMap() (map[string]interface{}, error)
}

func (opts DisableAutoRenewOpts) ToDisableAutoRenewMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}


/**
 * A customer can query one or all yearly/monthly resources on the customer platform.
 * This API can be invoked only by the customer AK/SK or token.
 */
func QueryCustomerPeriodResourcesList(client *gophercloud.ServiceClient, opts QueryCustomerPeriodResourcesListOptsBuilder) (r QueryCustomerPeriodResourcesListResult) {
	domainID := client.ProviderClient.DomainID
	url := getQueryCustomerPeriodResourcesListURL(client, domainID)
	if opts != nil {
		query, err := opts.ToQueryCustomerPeriodResourcesListMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url,&r.Body,nil)
	}

	return
}

/**
 * A customer can renew its yearly/monthly resources on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func RenewSubscriptionByResourceId(client *gophercloud.ServiceClient, opts RenewSubscriptionByResourceIdOptsBuilder) (r RenewSubscriptionByResourceIdResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToRenewSubscriptionByResourceIdMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getRenewSubscriptionByResourceIdURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}


/**
 * A customer can enable automatic renewal for its yearly/monthly resources on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func EnableAutoRenew(client *gophercloud.ServiceClient, opts EnableAutoRenewOptsBuilder,resourceId string) (r EnableAutoRenewResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToEnableAutoRenewMap()
		if err != nil {
			r.Err = err
			return
		}

		_, r.Err = client.Post(getEnableAutoRenewURL(client, domainID,resourceId,body["action_id"].(string)), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * A customer can unsubscribe from its yearly/monthly resources on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func UnsubscribeByResourceId(client *gophercloud.ServiceClient, opts UnsubscribeByResourceIdOptsBuilder) (r UnsubscribeByResourceIdResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToUnsubscribeByResourceIdMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getUnsubscribeByResourceIdURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * A customer can disable automatic renewal for its yearly/monthly resources on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func DisableAutoRenew(client *gophercloud.ServiceClient, opts DisableAutoRenewOptsBuilder,resourceId string) (r DisableAutoRenewResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToDisableAutoRenewMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Delete(getDisableAutoRenewURL(client, domainID,resourceId,body["action_id"].(string)), &gophercloud.RequestOpts{
			OkCodes: []int{200},
			JSONResponse: &r.Body,
		})
	}

	return
}


