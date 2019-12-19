package enquiry

import "github.com/gophercloud/gophercloud"

// POST /v1.0/{domain_id}/customer/product-mgr/query-rating
func getQueryRatingURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/product-mgr/query-rating")
}
