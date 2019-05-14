package peerings

import (
	"github.com/gophercloud/gophercloud"
	"net/url"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return baseURL(c)
}

func DeleteURL(c *gophercloud.ServiceClient, peeringID string) string {
	return baseURL(c) + "/" + peeringID
}

func GetURL(c *gophercloud.ServiceClient, peeringID string) string {
	return baseURL(c) + "/" + peeringID
}

func ListURL(c *gophercloud.ServiceClient) string {
	return baseURL(c)
}

func UpdateURL(c *gophercloud.ServiceClient, peeringID string) string {
	return baseURL(c) + "/" + peeringID
}

func AcceptURL(c *gophercloud.ServiceClient, peeringID string) string {
	return baseURL(c) + "/" + peeringID + "/accept"
}
func RejectURL(c *gophercloud.ServiceClient, peeringID string) string {
	return baseURL(c) + "/" + peeringID + "/reject"
}

func baseURL(c *gophercloud.ServiceClient) string {
	u, _ := url.Parse(c.ResourceBaseURL())
	return u.Scheme + "://" + u.Host + "/v2.0/" + "vpc/peerings"
}
