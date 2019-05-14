package bandwidths

import "github.com/gophercloud/gophercloud"

func PostURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("bandwidths")
}

func BatchPostURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("batch-bandwidths")
}
func UpdateURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}

func DeleteURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}

func InsertURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID, "insert")
}

func RemoveURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID, "remove")
}
