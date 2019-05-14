package cloudservers

import (
	"github.com/gophercloud/gophercloud"
	"net/url"
	"fmt"
)

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers")
}

func deleteURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "delete")
}

func getURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID)
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "details")
}

func updateURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID)
}

func autorecoveryURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID, "autorecovery")
}

func actionURL(sc *gophercloud.ServiceClient, serverID string) string {

	u, _ := url.Parse(sc.ResourceBaseURL())
	return fmt.Sprintf("%s://%s/v1.0/servers/%s/action", u.Scheme, u.Host, serverID)

}

func batchChangeURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "batch-changeos")
}
