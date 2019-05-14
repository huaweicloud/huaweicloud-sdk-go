package groups

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_group")
}

func DeleteURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group", scalingGroupId)
}

func EnableURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group", scalingGroupId, "action")
}

func GetURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group", scalingGroupId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_group")
}

func UpdateURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group", scalingGroupId)
}
