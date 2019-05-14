package quotas

import (
	"github.com/gophercloud/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
