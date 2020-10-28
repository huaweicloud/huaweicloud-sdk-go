package predefinetags

import "github.com/gophercloud/gophercloud"

func createOrDeleteURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("predefine_tags/action")
}

func updateURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("predefine_tags")
}

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("predefine_tags")
}
