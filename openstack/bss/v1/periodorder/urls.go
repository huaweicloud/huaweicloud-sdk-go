package periodorder

import "github.com/gophercloud/gophercloud"

//GET /v1.0/{domain_id}/common/order-mgr/orders/detail
func getQueryOrderListURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/orders/detail")
}

//GET /v1.0/{domain_id}/common/order-mgr/orders/{order_id}
func getQueryOrderDetailURL(client *gophercloud.ServiceClient, domainId string, orderId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/orders", orderId)
}

//POST /v1.0/{domain_id}/customer/order-mgr/order/pay
func getPayPeriodOrderURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/order-mgr/order/pay")
}

//DELETE /v1.0/{domain_id}/customer/order-mgr/orders/{order_id}
func getUnsubscribePeriodOrderURL(client *gophercloud.ServiceClient, domainId string, orderId string) string {
	return client.ServiceURL(domainId, "customer/order-mgr/orders", orderId)
}

//PUT /v1.0/{domain_id}/customer/order-mgr/orders/actions
func getCancelOrderURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/order-mgr/orders/actions")
}

//GET /v1.0/{domain_id}/common/order-mgr/orders-resource/{order_id}
func getQueryResourceStatusByOrderIdURL(client *gophercloud.ServiceClient, domainId string, orderId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/orders-resource", orderId)
}

//GET /v1.0/{domain_id}/common/order-mgr/orders/refund-order
func getQueryRefundOrderAmountURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/orders/refund-order")
}
