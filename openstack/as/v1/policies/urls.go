package policies

import (
	"github.com/gophercloud/gophercloud"
)

func ActionURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId, "action")
}

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_policy")
}

func DeleteURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId)
}

func GetURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId)
}

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_policy", scalingGroupId, "list")
}

func UpdateURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy", scalingPolicyId)
}
