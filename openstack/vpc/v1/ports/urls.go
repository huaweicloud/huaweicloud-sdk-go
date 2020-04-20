package ports

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ports")
}

func DeleteURL(c *gophercloud.ServiceClient, portId string) string {
	return c.ServiceURL("ports", portId)
}

func GetURL(c *gophercloud.ServiceClient, portId string) string {
	return c.ServiceURL("ports", portId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ports")
}

func UpdateURL(c *gophercloud.ServiceClient, portId string) string {
	return c.ServiceURL("ports", portId)
}