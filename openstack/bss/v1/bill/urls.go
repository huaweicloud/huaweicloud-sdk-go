package bill

import "github.com/gophercloud/gophercloud"

//GET /v1.0/{partner_id}/partner/account-mgr/subcustomer-bills  查询客户月度消费账单
func getQueryPartnerMonthlyBillsURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/account-mgr/subcustomer-bills")
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