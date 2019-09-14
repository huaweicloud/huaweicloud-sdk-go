package instances

import "github.com/gophercloud/gophercloud"

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("instances")
}

func deleteURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("instances", serverID)
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("instances")
}

func restartURL(sc *gophercloud.ServiceClient, instancesId string) string {
	return sc.ServiceURL("instances", instancesId, "action")
}

func singletohaURL(sc *gophercloud.ServiceClient, instancesId string) string {
	return sc.ServiceURL("instances", instancesId, "action")
}

func resizeURL(sc *gophercloud.ServiceClient, instancesId string) string {
	return sc.ServiceURL("instances", instancesId, "action")
}

func enlargeURL(sc *gophercloud.ServiceClient, instancesId string) string {
	return sc.ServiceURL("instances", instancesId, "action")
}

func listerrorlogURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "errorlog")
}

func listslowlogURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "slowlog")
}

