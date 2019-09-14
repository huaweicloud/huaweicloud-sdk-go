package database

import "github.com/gophercloud/gophercloud"


func createURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "database")
}

func listURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "database","detail")
}

func deleteURL(sc *gophercloud.ServiceClient, instanceID string,dbName string) string {
	return sc.ServiceURL("instances", instanceID, "database",dbName)
}
