package volumetransfer

import (
	"github.com/gophercloud/gophercloud"
)


func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer")
}
func acceptURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id,"accept")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer")
}

func detailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer", "detail")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}

