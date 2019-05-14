package quotas

import (
	"github.com/gophercloud/gophercloud"
)

func List(client *gophercloud.ServiceClient) (r ListResult) {
	url := ListURL(client)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListWithInstances(client *gophercloud.ServiceClient, scalingGroupId string) (r ListWithInstancesResult) {
	url := ListWithInstancesURL(client, scalingGroupId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
