package configures

import (
	"github.com/gophercloud/gophercloud"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_configuration")
}

func DeleteURL(c *gophercloud.ServiceClient, scalingConfigurationId string) string {
	return c.ServiceURL("scaling_configuration", scalingConfigurationId)
}

func DeleteWithBatchURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_configurations")
}

func GetURL(c *gophercloud.ServiceClient, scalingConfigurationId string) string {
	return c.ServiceURL("scaling_configuration", scalingConfigurationId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_configuration")
}
