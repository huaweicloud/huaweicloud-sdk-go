package securitygrouprules

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-group-rules")
}

func DeleteURL(c *gophercloud.ServiceClient, securityGroupsRulesId string) string {
	return c.ServiceURL("security-group-rules", securityGroupsRulesId)
}

func GetURL(c *gophercloud.ServiceClient, securityGroupsRulesId string) string {
	return c.ServiceURL("security-group-rules", securityGroupsRulesId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-group-rules")
}
