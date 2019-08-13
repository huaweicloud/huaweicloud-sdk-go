package ptrs

import (
	"github.com/gophercloud/gophercloud"
)

func GetURL(c *gophercloud.ServiceClient, region string, floatingipId string) string {
	return c.ServiceURL("reverse", "floatingips", region+":"+floatingipId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("reverse", "floatingips")
}

func RestoreURL(c *gophercloud.ServiceClient, region string, floatingipId string) string {
	return c.ServiceURL("reverse", "floatingips", region+":"+floatingipId)
}

func SetupURL(c *gophercloud.ServiceClient, region string, floatingipId string) string {
	return c.ServiceURL("reverse", "floatingips", region+":"+floatingipId)
}

func UpdateURL(c *gophercloud.ServiceClient, region string, floatingipId string) string {
	return c.ServiceURL("reverse", "floatingips", region+":"+floatingipId)
}
