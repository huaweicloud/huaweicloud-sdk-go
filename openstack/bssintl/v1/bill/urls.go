package bill

import "github.com/gophercloud/gophercloud"

//GET /v1.0/{partner_id}/partner/account-mgr/postpaid-bill-summary
func getQueryPartnerMonthlyBillsURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/account-mgr/postpaid-bill-summary")
}

//GET /v1.0/{partner_id}/customer/account-mgr/bill/monthly-sum
func getQueryMonthlyExpenditureSummaryURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/account-mgr/bill/monthly-sum")
}

//GET /v1.0/{partner_id}/customer/account-mgr/bill/res-records
func getQueryResourceUsageDetailsURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/account-mgr/bill/res-records")
}

//GET /v1.0/{partner_id}/customer/account-mgr/bill/res-fee-records
func getQueryResourceUsageRecordURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "customer/account-mgr/bill/res-fee-records")
}