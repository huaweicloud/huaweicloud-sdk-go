package db_user

import "github.com/gophercloud/gophercloud"

func createURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_user")
}

func listURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_user", "detail")
}

func deleteURL(sc *gophercloud.ServiceClient, instanceID string, dbuser string) string {
	return sc.ServiceURL("instances", instanceID, "db_user",dbuser)
}
