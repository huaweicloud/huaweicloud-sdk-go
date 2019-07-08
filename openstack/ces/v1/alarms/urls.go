package alarms

import (
	"github.com/gophercloud/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("alarms")
}

func deleteURL(c *gophercloud.ServiceClient, alarmId string) string {
	return c.ServiceURL("alarms", alarmId)
}

func getURL(c *gophercloud.ServiceClient, alarmId string) string {
	return c.ServiceURL("alarms", alarmId)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("alarms")
}

func updateURL(c *gophercloud.ServiceClient, alarmId string) string {
	return c.ServiceURL("alarms", alarmId, "action")
}
