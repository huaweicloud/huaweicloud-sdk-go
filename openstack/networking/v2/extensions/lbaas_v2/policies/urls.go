package policies



import "github.com/gophercloud/gophercloud"

const (
	rootPath     = "lbaas"
	resourcePath = "l7policies"
)
//GET list and post url
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

//GET details put delete url
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
