package quotas

import (
	"github.com/gophercloud/gophercloud"
)

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
