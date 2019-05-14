package policies

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_policy")
}

func GetURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId)
}

func ListURL(c *gophercloud.ServiceClient, scalingResourceId string) string {
	return c.ServiceURL("scaling_policy", scalingResourceId, "list")
}

func UpdateURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId)
}
