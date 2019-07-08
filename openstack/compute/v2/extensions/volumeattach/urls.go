package volumeattach

import (
	"github.com/gophercloud/gophercloud"
	"fmt"
	"strconv"
)

const resourcePath = "os-volume_attachments"

func resourceURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, resourcePath)
}

func listURL(c *gophercloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

func createURL(c *gophercloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

func getURL(c *gophercloud.ServiceClient, serverID, aID string) string {
	return c.ServiceURL("servers", serverID, resourcePath, aID)
}

func deleteURL(c *gophercloud.ServiceClient, serverID, aID string) string {
	return getURL(c, serverID, aID)
}

func deleteWithFlagURL(c *gophercloud.ServiceClient, serverID, volumeID string, deleteFlag int) string {
	deleteFlagStr := strconv.Itoa(deleteFlag)
	newStr := fmt.Sprintf("%s?delete_flag=%s", volumeID, deleteFlagStr)
	return c.ServiceURL("servers", serverID, resourcePath, newStr)
}
