package account

import "github.com/gophercloud/gophercloud"

// /v1.0/{domain_id}/customer/account-mgr/bill/resource-daily
func getURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/account-mgr/bill/resource-daily")
}
