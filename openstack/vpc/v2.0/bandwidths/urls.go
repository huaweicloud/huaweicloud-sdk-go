package bandwidths

import "github.com/gophercloud/gophercloud"

func UpdateURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}
