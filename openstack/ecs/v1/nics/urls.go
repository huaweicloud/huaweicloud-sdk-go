package nics

import "github.com/gophercloud/gophercloud"

//add nics url
func addUrl(sc *gophercloud.ServiceClient, seviceId string) string {
	return sc.ServiceURL("cloudservers", seviceId, "nics")
}

//delete nics url
func deleteUrl(sc *gophercloud.ServiceClient, seviceId string) string {
	return sc.ServiceURL("cloudservers", seviceId, "nics", "delete")
}

//Bind and Unbind Virtual IP url
func putURL(sc *gophercloud.ServiceClient, nicId string) string {
	return sc.ServiceURL("cloudservers", "nics", nicId)
}
