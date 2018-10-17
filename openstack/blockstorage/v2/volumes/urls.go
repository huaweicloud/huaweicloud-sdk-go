package volumes

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func getQuotaSetURL(c *gophercloud.ServiceClient, projectId string) string {
	newStr := fmt.Sprintf("%s?usage=True", projectId)
	return c.ServiceURL("os-quota-sets", newStr)
}
