package periodorder

import "github.com/gophercloud/gophercloud"

type QueryOrderListOpts struct {
	//Order ID.
	OrderId string `q:"order_id"`

	//Customer account ID
	CustomerId string `q:"customer_id"`

	//Start time of order creation.
	CreateTimeBegin string `q:"create_time_begin"`

	//End time of order creation
	CreateTimeEnd string `q:"create_time_end"`

	//Cloud service type code
	ServiceType string `q:"service_type"`

	//Order status
	Status string `q:"status"`

	//Order type
	OrderType string `q:"order_type"`

	//Number of records per page.
	PageSize *int `q:"page_size" required:"true"`

	//Current page number.
	PageIndex *int `q:"page_index" required:"true"`

	//Sorting order of the orders.
	Sort string `q:"sort"`

	//Start time of order payment.
	PaymentTimeBegin string `q:"payment_time_begin"`

	//End time of order payment.
	PaymentTimeEnd string `q:"payment_time_end"`
}

type QueryOrderDetailOpts struct {
	//Indicates the page number
	Offset int `q:"offset"`

	//Indicates the number of records displayed on each page
	Limit int `q:"limit"`
}

type PayPeriodOrderOpts struct {
	//Order ID.
	OderId string `json:"orderId" required:"true"`

	//Payment account type
	PayAccountType *int `json:"payAccountType,omitempty"`

	//partner account ID
	BpId string `json:"bpId"`

	//Coupon IDs
	CouponIds []string `json:"couponIds"`
}

type UnsubscribePeriodOrderOpts struct {
	//Unsubscription type.
	UnsubType int `q:"unsub_type" required:"true"`

	//Unsubscription reason classification
	UnsubscribeReasonType int `q:"unsubscribe_reason_type"`

	//Unsubscription reason, which is generally specified by the customer.
	UnsubscribeReason string `q:"unsubscribe_reason"`
}

type CancelOrderOpts struct {
	//Order ID.
	OrderId string `json:"orderId" required:"true"`
}

type QueryResourceStatusByOrderIdOpts struct {
	//Page number.
	Offset int `q:"offset"`

	//Number of records per page.
	Limit int `q:"limit"`
}

type QueryRefundOrderAmountOpts struct {
	//ID of an unsubscription order or degrade order.
	OrderId string `q:"order_id"`
}

type QueryOrderListBuilder interface {
	ToQueryOrderListQuery() (string, error)
}

type QueryOrderDetailBuilder interface {
	ToQueryOrderDetailQuery() (string, error)
}

type PayPeriodOrderBuilder interface {
	ToPayPeriodOrderOptsMaps() (map[string]interface{}, error)
}

type UnsubscribePeriodOrderBuilder interface {
	ToUnsubscribePeriodOrderQuery() (string, error)
}

type CancelOrderBuilder interface {
	ToCancelOrderOptsMaps() (map[string]interface{}, error)
}

type QueryResourceStatusByOrderIdBuilder interface {
	ToQueryResourceStatusByOrderIdQuery() (string, error)
}

type QueryRefundOrderAmountBuilder interface {
	ToQueryRefundOrderAmountQuery() (string, error)
}

func (opts QueryOrderListOpts) ToQueryOrderListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func (opts QueryOrderDetailOpts) ToQueryOrderDetailQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func (opts PayPeriodOrderOpts) ToPayPeriodOrderOptsMaps() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func (opts UnsubscribePeriodOrderOpts) ToUnsubscribePeriodOrderQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func (opts CancelOrderOpts) ToCancelOrderOptsMaps() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func (opts QueryResourceStatusByOrderIdOpts) ToQueryResourceStatusByOrderIdQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func (opts QueryRefundOrderAmountOpts) ToQueryRefundOrderAmountQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

/**
 * After a customer purchases yearly/monthly resources, it can query the orders in different status on the customer platform, such as in the pending approval, processing, canceled, completed, and pending payment statuses.
 * This API can be invoked using the customer AK/SK or token.
 */
func QueryOrderList(client *gophercloud.ServiceClient, opts QueryOrderListBuilder) (r QueryOrderListResult) {
	domainID := client.ProviderClient.DomainID
	url := getQueryOrderListURL(client, domainID)
	if opts != nil {
		query, err := opts.ToQueryOrderListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

/**
 * Customers can query order details on the customer platform.
 * This API can be invoked using the customer AK/SK or token.
 */
func QueryOrderDetail(client *gophercloud.ServiceClient, opts QueryOrderDetailBuilder, orderId string) (r QueryOrderDetailResult) {
	domainID := client.ProviderClient.DomainID
	url := getQueryOrderDetailURL(client, domainID, orderId)
	if opts != nil {
		query, err := opts.ToQueryOrderDetailQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

/**
 * A customer can pay yearly-monthly product orders in the pending payment status on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func PayPeriodOrder(client *gophercloud.ServiceClient, opts PayPeriodOrderBuilder) (r PayPeriodOrderResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToPayPeriodOrderOptsMaps()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getPayPeriodOrderURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * A customer can unsubscribe yearly-monthly product orders in the subscribed, changing, or failed to be provisioned status on the customer platform.
 * This API can be invoked using the customer AK/SK or token only.
 */
func UnsubscribePeriodOrder(client *gophercloud.ServiceClient, opts UnsubscribePeriodOrderBuilder, orderId string) (r UnsubscribePeriodOrderResult) {
	domainID := client.ProviderClient.DomainID
	url := getUnsubscribePeriodOrderURL(client, domainID, orderId)
	if opts != nil {
		query, err := opts.ToUnsubscribePeriodOrderQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		JSONResponse: &r.Body,
	})
	return
}

/**
 * A customer can cancel subscription of yearly-monthly product orders in the pending payment status on the partner sales platform.
 * This API can be invoked using the customer token only.
 */
func CancelOrder(client *gophercloud.ServiceClient, opts CancelOrderBuilder,actionId string) (r CancelOrderResult) {
	domainID := client.ProviderClient.DomainID
	url := getCancelOrderURL(client, domainID)
	url += "?action_id="+actionId
	if opts != nil {
		body, err := opts.ToCancelOrderOptsMaps()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Put(url, body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * Customers can query resource details and provisioning status of an order on the customer platform.
 * This API can be invoked only by the customer AK/SK or token.
 */
func QueryResourceStatusByOrderId(client *gophercloud.ServiceClient, opts QueryResourceStatusByOrderIdBuilder, orderId string) (r QueryResourceStatusByOrderIdResult) {
	domainID := client.ProviderClient.DomainID
	url := getQueryResourceStatusByOrderIdURL(client, domainID, orderId)
	if opts != nil {
		query, err := opts.ToQueryResourceStatusByOrderIdQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

/**
 * A customer can query the resources and original orders of the unsubscription amount for an unsubscription order or degrade order.
 * This API can be invoked using the AK/SK or token of the partner or the token of the partner's customer.
 */
func QueryRefundOrderAmount(client *gophercloud.ServiceClient, opts QueryRefundOrderAmountBuilder) (r QueryRefundOrderAmountResult) {
	domainID := client.ProviderClient.DomainID
	url := getQueryRefundOrderAmountURL(client, domainID)
	if opts != nil {
		query, err := opts.ToQueryRefundOrderAmountQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
