package availabilityzones

import "github.com/gophercloud/gophercloud"

// listURL generates URL for list avaliabilityzones
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
