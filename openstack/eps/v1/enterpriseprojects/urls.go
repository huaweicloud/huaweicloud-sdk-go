package enterpriseprojects

import "github.com/gophercloud/gophercloud"

func createEPURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("enterprise-projects")
}

func updateURL(sc *gophercloud.ServiceClient, epID string) string {
	return sc.ServiceURL("enterprise-projects", epID)
}

func getURL(sc *gophercloud.ServiceClient, epID string) string {
	return sc.ServiceURL("enterprise-projects", epID)
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("enterprise-projects")
}

func getQuotasURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("enterprise-projects", "quotas")
}

func actionURL(sc *gophercloud.ServiceClient, epID string) string {
	return sc.ServiceURL("enterprise-projects", epID, "action")
}

func filterResourcesURL(sc *gophercloud.ServiceClient, epID string) string {
	return sc.ServiceURL("enterprise-projects", epID, "resources", "filter")
}

func migrateResourcesURL(sc *gophercloud.ServiceClient, epID string) string {
	return sc.ServiceURL("enterprise-projects", epID ,"resources-migrate")
}
