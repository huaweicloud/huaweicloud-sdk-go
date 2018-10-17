package cloudservers

import "github.com/gophercloud/gophercloud"

func resetPwdURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("servers", serverID, "os-reset-password")
}

func changeURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID, "changeos")
}
