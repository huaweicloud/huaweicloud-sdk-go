package logs

import (
	"github.com/gophercloud/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_activity_log", scalingGroupId)
}
