package realnameauth

import "github.com/gophercloud/gophercloud"

//POST /v1.0/{partner_id}/partner/customer-mgr/realname-auth/individual
func getIndividualRealNameAuthURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/realname-auth/individual")
}

//POST /v1.0/{partner_id}/partner/customer-mgr/realname-auth/enterprise
func getEnterpriseRealNameAuthURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/realname-auth/enterprise")
}

//GET /v1.0/{partner_id}/partner/customer-mgr/realname-auth/result
func getQueryRealNameAuthURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/realname-auth/result")
}

//PUT /v1.0/{partner_id}/partner/customer-mgr/realname-auth/enterprise
func getChangeEnterpriseRealNameAuthURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/customer-mgr/realname-auth/enterprise")
}