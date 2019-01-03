package publicips

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("publicips")
}

func DeleteURL(c *gophercloud.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}

func GetURL(c *gophercloud.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("publicips")
}

func UpdateURL(c *gophercloud.ServiceClient, publicipId string) string {
	return c.ServiceURL("publicips", publicipId)
}
