package metrics

import "github.com/gophercloud/gophercloud"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
