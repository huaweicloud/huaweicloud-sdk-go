package payperuseresource

import "github.com/gophercloud/gophercloud"

//POST /v1.0/{partner_id}/partner/customer-mgr/customer-resource/query-resources
func getQueryCustomerResourceURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/customer-resource/query-resources")
}


