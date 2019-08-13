package quotas

import (
	"github.com/gophercloud/gophercloud"
)

func Get(client *gophercloud.ServiceClient) (r GetResult) {
    _, r.Err = client.Get(getURL(client), &r.Body, &gophercloud.RequestOpts{
        OkCodes: []int{200},
    })

    return
}
