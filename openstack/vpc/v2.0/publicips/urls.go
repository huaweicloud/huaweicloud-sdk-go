package publicips

import "github.com/gophercloud/gophercloud"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}