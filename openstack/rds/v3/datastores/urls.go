package datastores

import "github.com/gophercloud/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("datastores", databasename)
}
