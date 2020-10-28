package domains

import "github.com/gophercloud/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("domains")
}

func getURL(client *gophercloud.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("domains")
}

func deleteURL(client *gophercloud.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func updateURL(client *gophercloud.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func listDomainsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "domains")
}
func getPwdStrengthPolicyURL(client *gophercloud.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID, "config", "security_compliance")
}

func getPwdStrengthPolicyURLByOption(client *gophercloud.ServiceClient, domainID string, option string) string {
	if option == "" {
		return client.ServiceURL("domains", domainID, "config", "security_compliance")
	} else {
		return client.ServiceURL("domains", domainID, "config", "security_compliance", option)
	}
}
