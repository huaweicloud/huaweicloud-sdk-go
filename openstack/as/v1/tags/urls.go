package tags

import (
	"github.com/gophercloud/gophercloud"
)

func ListResourceTagsURL(c *gophercloud.ServiceClient, resourceType string, resourceId string) string {
	return c.ServiceURL(resourceType, resourceId, "tags")
}

func ListTenantTagsURL(c *gophercloud.ServiceClient, resourceType string) string {
	return c.ServiceURL(resourceType, "tags")
}

func ListInstanceTagsURL(c *gophercloud.ServiceClient, resourceType string) string {
	return c.ServiceURL(resourceType, "resource_instances", "action")
}

func UpdateURL(c *gophercloud.ServiceClient, resourceType string, resourceId string) string {
	return c.ServiceURL(resourceType, resourceId, "tags", "action")
}
