package ports

import (
	"github.com/gophercloud/gophercloud"
	"net/url"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return baseUrl(c)
}

func DeleteURL(c *gophercloud.ServiceClient, portId string) string {
	return baseUrlWithID(c, portId)
}

func GetURL(c *gophercloud.ServiceClient, portId string) string {
	return baseUrlWithID(c, portId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return baseUrl(c)
}

func UpdateURL(c *gophercloud.ServiceClient, portId string) string {
	return baseUrlWithID(c, portId)
}

func baseUrl(c *gophercloud.ServiceClient) string {
	u, _ := url.Parse(c.ResourceBaseURL())
	return u.Scheme + "://" + u.Host + "/v1/" + "ports"
}

func baseUrlWithID(c *gophercloud.ServiceClient, portId string) string {

	return baseUrl(c) + "/" + portId
}
