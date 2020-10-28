package credentials

import "github.com/gophercloud/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("OS-CREDENTIAL", "credentials")
}

func createTemporaryAkURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("OS-CREDENTIAL", "securitytokens")
}

func deletePermanentAkURL(client *gophercloud.ServiceClient, ak string) string {
	return client.ServiceURL("OS-CREDENTIAL", "credentials", ak)
}

func getAllPermanentAksURL(client *gophercloud.ServiceClient, userID string) string {
	if userID == "" {
		return client.ServiceURL("OS-CREDENTIAL", "credentials")
	} else {
		return client.ServiceURL("OS-CREDENTIAL", "credentials") + "?" + "user_id=" + userID
	}
}

func getPermanentAkURL(client *gophercloud.ServiceClient, ak string) string {
	return client.ServiceURL("OS-CREDENTIAL", "credentials", ak)
}

func updatePermanentAkURL(client *gophercloud.ServiceClient, ak string) string {
	return client.ServiceURL("OS-CREDENTIAL", "credentials", ak)
}