package cloudservers

import "github.com/gophercloud/gophercloud"

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers")
}

func deleteURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers","delete")
}

func getURL(sc *gophercloud.ServiceClient,serverID string) string {
	return sc.ServiceURL("cloudservers",serverID)
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers","details")
}

func updateURL(sc *gophercloud.ServiceClient,serverID string) string {
	return sc.ServiceURL("cloudservers",serverID)
}

