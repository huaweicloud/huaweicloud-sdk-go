package instances

import (
	"github.com/gophercloud/gophercloud"
)

func ActionURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group_instance", scalingGroupId, "action")
}

func DeleteURL(c *gophercloud.ServiceClient, instanceId string) string {
	return c.ServiceURL("scaling_group_instance", instanceId)
}

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_group_instance", scalingGroupId, "list")
}
