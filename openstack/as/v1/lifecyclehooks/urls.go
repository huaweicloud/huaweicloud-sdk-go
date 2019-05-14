package lifecyclehooks

import (
	"github.com/gophercloud/gophercloud"
)

func CallBackURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_instance_hook", scalingGroupId, "callback")
}

func CreateURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_lifecycle_hook", scalingGroupId)
}

func DeleteURL(c *gophercloud.ServiceClient, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}

func GetURL(c *gophercloud.ServiceClient, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_lifecycle_hook", scalingGroupId, "list")
}

func ListWithSuspensionURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_instance_hook", scalingGroupId, "list")
}

func UpdateURL(c *gophercloud.ServiceClient, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}
