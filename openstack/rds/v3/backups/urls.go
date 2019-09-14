package backups

import "github.com/gophercloud/gophercloud"

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("backups")
}

func updatepolicyURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "backups", "policy")
}

func getpolicyURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "backups", "policy")
}

func deleteURL(sc *gophercloud.ServiceClient, backupId string) string {
	return sc.ServiceURL("backups", backupId)
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("backups")
}

func restoreURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("instances")
}

func listfilesURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("backup-files")
}

func getrestoretimeURL(sc *gophercloud.ServiceClient, instanceId string) string {
	return sc.ServiceURL("instances", instanceId, "restore-time")
}

func recoveryURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("instances", "recovery")
}
