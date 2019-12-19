package utilities

import "github.com/gophercloud/gophercloud"

//POST /v1.0/{partner_id}/partner/common-mgr/verificationcode
func getSendVerificationCodeURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "partner/common-mgr/verificationcode")
}


