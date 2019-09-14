package trigger

import "github.com/gophercloud/gophercloud"

const (
    FGS     = "fgs"
    TRIGGER = "triggers"
)

func listURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, TRIGGER, functionUrn)
}

func createURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return listURL(c, functionUrn)
}

func deleteAllURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return listURL(c, functionUrn)
}

func deleteURL(c *gophercloud.ServiceClient, functionUrn, triggerTypeCode, triggerId string) string {
    return getURL(c, functionUrn, triggerTypeCode, triggerId)
}

func getURL(c *gophercloud.ServiceClient, functionUrn, triggerTypeCode, triggerId string) string {
    return c.ServiceURL(FGS, TRIGGER, functionUrn, triggerTypeCode, triggerId)
}
