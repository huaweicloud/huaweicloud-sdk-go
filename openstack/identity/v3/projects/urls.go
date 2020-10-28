package projects

import "github.com/gophercloud/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

func getURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

func deleteURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func updateURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func listProjectsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "projects" )
}

func listForUserURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "projects" )
}

func getDetailsURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}