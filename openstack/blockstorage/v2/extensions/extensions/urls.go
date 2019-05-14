package extensions

import "github.com/gophercloud/gophercloud"

// ListExtensionURL generates the URL for the extensions resource collection.
func ListExtensionURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("extensions")
}
