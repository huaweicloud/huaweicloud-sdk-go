package configurations

import "github.com/gophercloud/gophercloud"

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}
