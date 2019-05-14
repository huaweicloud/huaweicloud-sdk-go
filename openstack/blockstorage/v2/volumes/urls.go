package volumes

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
)

// `createURL` is a pure function. `createURL(c)` is a URL for which a POST
// request will response with a blockstorage struct in the service `c`.
func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

// `listURL` is a pure function. `detailURL(c)` is a URL for which a GET
// request will response with detail of blockstorage in the service `c`.
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

// `listBriefURL` is a pure function. `listURL(c)` is a URL for which a GET request
// will response with a list of blockstorages in the service `c`.
func listBriefURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

// `deleteURL` is a pure function. `deleteURL(c)` is a URL for which a DELETE
// request will response with a status code in the service  `c`.
func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

// `getURL` is a pure function. `getURL(c)` is a URL for which a GET request
// will response with a blockstorage in the service  `c`.
func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

// `updateURL` is a pure function. `updateURL(c)` is a URL for which a PUT
// request will response with a update blockstorage in the service `c`.
func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

// `metadataURL` is a pure function. `metadataURL(c)` is a URL for which a GET
// request will response with metadatas in the service `c`
func metadataURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "metadata")
}

// `metadataKeyURL` is a pure function. `metadataKeyURL(c, id, key)` is a URL
// for which a GET request will response with value of metadata.
func metadataKeyURL(c *gophercloud.ServiceClient, id, key string) string {
	return c.ServiceURL("volumes", id, "metadata", key)
}

// `actionURL` is a pure function. `actionURL(c, id)` is a URL for which a POST
// request will response with a status code.
func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}

func getQuotaSetURL(c *gophercloud.ServiceClient, projectId string) string {
	newStr := fmt.Sprintf("%s?usage=True", projectId)
	return c.ServiceURL("os-quota-sets", newStr)
}
