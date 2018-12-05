package policies

import "github.com/gophercloud/gophercloud"

const (
	ROOTPATH     = "lbaas"
	RESOURCEPATH = "l7policies"
)

//GET list and post url
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH)
}

//GET details put delete url
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH, id)
}

//GET rules list and post url
func rulesrootURL(c *gophercloud.ServiceClient, policyId string) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH, policyId, "rules")
}

//GET rules details put delete urlrulesid
func rulesresourceURL(c *gophercloud.ServiceClient, policyId, ruleId string) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH, policyId, "rules", ruleId)
}
