package roles

import "github.com/gophercloud/gophercloud"

const (
	rolePath = "roles"
)

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func getURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func updateURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func deleteURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func assignURL(client *gophercloud.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}

func addRoleForGroupOnDomainURL(client *gophercloud.ServiceClient, domianId string, groupID string, roleId string) string {
	return client.ServiceURL("domains", domianId, "groups", groupID, "roles", roleId)
}

func addRoleForGroupOnProjectURL(client *gophercloud.ServiceClient, projectID string, groupID string, roleId string) string {
	return client.ServiceURL("projects", projectID, "groups", groupID, "roles", roleId)
}

func headRoleForGroupOnProjectURL(client *gophercloud.ServiceClient, projectID string, groupID string, roleId string) string {
	return client.ServiceURL("projects", projectID, "groups", groupID, "roles", roleId)
}

func headRoleForGroupOnDomainURL(client *gophercloud.ServiceClient, domainID string, groupID string, roleId string) string {
	return client.ServiceURL("domains", domainID, "groups", groupID, "roles", roleId)
}

func deleteRoleForGroupOnDomainURL(client *gophercloud.ServiceClient, domainID string, groupID string, roleId string) string {
	return client.ServiceURL("domains", domainID, "groups", groupID, "roles", roleId)
}

func deleteRoleForGroupOnProjectURL(client *gophercloud.ServiceClient, projectID string, groupID string, roleId string) string {
	return client.ServiceURL("projects", projectID, "groups", groupID, "roles", roleId)
}

func addAlRoleForGroupOnlProjectURL(client *gophercloud.ServiceClient, domainID string, groupID string, roleID string) string {
	return client.ServiceURL("OS-INHERIT", "domains", domainID, "groups", groupID, "roles", roleID, "inherited_to_projects")
}

func getRolesForGroupOnDomain(client *gophercloud.ServiceClient, domainID string, groupID string) string {
	return client.ServiceURL("domains", domainID, "groups", groupID, "roles")
}

func getRolesForGroupOnProject(client *gophercloud.ServiceClient, projectID string, groupID string) string {
	return client.ServiceURL("projects", projectID, "groups", groupID, "roles")
}
