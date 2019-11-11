package cloudservers

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"net/url"
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

func listDetailURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "detail")
}

func batchActionURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "action")
}

func batchUpdateURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "server-name")
}


func batchTagActionURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID, "tags", "action")
}

func listProjectTagsURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers", "tags")
}

func listServerTagsURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID, "tags")
}
