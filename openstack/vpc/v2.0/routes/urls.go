package routes

import (
	"github.com/gophercloud/gophercloud"
	"net/url"

)

func CreateURL(c *gophercloud.ServiceClient) string {
	return baseURL(c)
}

func DeleteURL(c *gophercloud.ServiceClient, routeId string) string {
	return baseURLWithID(c,routeId)
}

func GetURL(c *gophercloud.ServiceClient, routeId string) string {
	return baseURLWithID(c,routeId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return baseURL(c)
}

func baseURL(c *gophercloud.ServiceClient) string {
	u, _ := url.Parse(c.ResourceBaseURL())
	return u.Scheme + "://" + u.Host + "/v2.0/" + "vpc/routes"
}

func baseURLWithID(c *gophercloud.ServiceClient, routeId string) string {
	return baseURL(c) + "/" + routeId
}
