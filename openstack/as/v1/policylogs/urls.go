package policylogs

import (
	"github.com/gophercloud/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy_execute_log", scalingPolicyId)
}
