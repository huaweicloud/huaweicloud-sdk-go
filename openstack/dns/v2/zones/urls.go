/*
package zones

import "github.com/gophercloud/gophercloud"

func baseURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *gophercloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}


 */

package zones

import "github.com/gophercloud/gophercloud"

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func DeleteURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId)
}

func DisassociateRouterURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId, "disassociaterouter")
}

func GetURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func ListNameServersURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId, "nameservers")
}
func UpdateURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId)
}

func AssociateRouterURL(c *gophercloud.ServiceClient, zoneId string) string {
	return c.ServiceURL("zones", zoneId, "associaterouter")
}

