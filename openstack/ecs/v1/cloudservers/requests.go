package cloudservers

import (
	"github.com/gophercloud/gophercloud"
	"fmt"
)


// Get requests details on a single server, by ID.
func Get(client *gophercloud.ServiceClient, serverID string) (r GetResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err=err
		return r
	}
	_, r.Err = client.Get(getURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}
