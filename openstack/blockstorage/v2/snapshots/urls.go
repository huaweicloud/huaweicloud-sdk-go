package snapshots

import "github.com/gophercloud/gophercloud"

// createURL generate url for creating snapshot
func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

// deleteURL generate url for deleting snapshot
func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

// getURL generate url for getting snaohsot
func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

// updateURL generate url for updating snapshot
func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

// listURL generate url for listing snapshots
func listURL(c *gophercloud.ServiceClient) string {
	return createURL(c)
}

// detailURL generate url for getting detail of snapshot
func detailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots", "detail")
}

// metadataURL generate url for snapshot's metadata
func metadataURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "metadata")
}

// metadataKeyURL generate url for getting key of snapshot's metadata
func metadataKeyURL(c *gophercloud.ServiceClient, id, key string) string {
	return c.ServiceURL("snapshots", id, "metadata", key)
}

