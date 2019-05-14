package privateips

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("privateips")
}

func DeleteURL(c *gophercloud.ServiceClient, privateipId string) string {
	return c.ServiceURL("privateips", privateipId)
}

func GetURL(c *gophercloud.ServiceClient, privateipId string) string {
	return c.ServiceURL("privateips", privateipId)
}

func ListURL(c *gophercloud.ServiceClient, subnetId string) string {
	return c.ServiceURL("subnets", subnetId, "privateips")
}
