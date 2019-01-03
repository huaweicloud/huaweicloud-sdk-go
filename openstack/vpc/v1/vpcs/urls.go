package vpcs

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("vpcs")
}

func DeleteURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId)
}

func GetURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("vpcs")
}

func UpdateURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId)
}
