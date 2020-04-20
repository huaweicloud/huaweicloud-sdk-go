package regions

import "github.com/gophercloud/gophercloud"

// listURL generate url to list regions
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("regions")
}

// getURL generate url to show region details
func getURL(client *gophercloud.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}

// createURL generate url for creating regions
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("regions")
}

// updateURL generate url for updating region
func updateURL(client *gophercloud.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}

// deleteURL generate url for deleting region
func deleteURL(client *gophercloud.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}
