package periodresource

import "github.com/gophercloud/gophercloud"

//GET /v1.0/{partner_id}/common/order-mgr/resources/detail 查询客户包周期资源列表
func getQueryCustomerPeriodResourcesListURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/detail")
}

//POST /v1.0/{partner_id}/common/order-mgr/resources/renew 续订包周期资源
func getRenewSubscriptionByResourceIdURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/renew")
}

//POST /v1.0/{partner_id}/common/order-mgr/resources/delete 退订包周期资源
func getUnsubscribeByResourceIdURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/delete")
}

//POST /v1.0/{partner_id}/common/order-mgr/resources/{resource_id}/actions?action_id=autorenew 设置包周期资源自动续费
func getEnableAutoRenewURL(client *gophercloud.ServiceClient, domainId string,resourceId string,actionId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources",resourceId,"actions?action_id=")+actionId
}

//DELETE /v1.0/{partner_id}/common/order-mgr/resources/{resource_id}/actions?action_id=autorenew 退订包周期资源
func getDisableAutoRenewURL(client *gophercloud.ServiceClient, domainId string,resourceId string,actionId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources",resourceId,"actions?action_id=")+actionId
}