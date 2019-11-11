package storagetype

import "github.com/gophercloud/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("storage-type", databasename)
}
