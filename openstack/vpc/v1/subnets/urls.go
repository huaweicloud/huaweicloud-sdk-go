package subnets

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("subnets")
}

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("subnets")
}

func DeleteURL(c *gophercloud.ServiceClient, vpcId string, subnetId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets", subnetId)
}

func GetURL(c *gophercloud.ServiceClient, subnetId string) string {
	return c.ServiceURL("subnets", subnetId)
}

func UpdateURL(c *gophercloud.ServiceClient, vpcId string, subnetId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets", subnetId)
}
