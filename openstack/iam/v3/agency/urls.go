package agency

import "github.com/gophercloud/gophercloud"

const (
	agencyPath = "OS-AGENCY"
)

// 查询指定条件下的委托列表
func listAgenciesUrl(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(agencyPath, "agencies")
}

// 查询委托详情
func queryAgencyDetailsUrl(client *gophercloud.ServiceClient, agencyId string) string {
	return client.ServiceURL(agencyPath, "agencies", agencyId)
}

//  创建委托
func createAgencyURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(agencyPath, "agencies")
}

//  修改委托
func updateAgencyURL(client *gophercloud.ServiceClient, agencyId string) string {
	return client.ServiceURL(agencyPath, "agencies", agencyId)
}

//  删除委托
func deleteAgencyURL(client *gophercloud.ServiceClient, agencyId string) string {
	return client.ServiceURL(agencyPath, "agencies", agencyId)
}

//  查询全局服务中的委托权限
func listPermissionsForAgencyOnDomainURL(client *gophercloud.ServiceClient, domainId string, agencyId string) string {
	return client.ServiceURL(agencyPath, "domains", domainId, "agencies", agencyId, "roles")
}

//  查询项目服务中的委托权限
func listPermissionsForAgencyOnProjectURL(client *gophercloud.ServiceClient, projectId string, agencyId string) string {
	return client.ServiceURL(agencyPath, "projects", projectId, "agencies", agencyId, "roles")
}

//  为委托授予全局服务权限
func grantPermissionToAgencyOnDomainURL(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "domains", domainId, "agencies", agencyId, "roles", roleId)
}

//  为委托授予项目服务权限
func grantPermissionToAgencyOnProjectURL(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "projects", projectId, "agencies", agencyId, "roles", roleId)
}

//  查询委托是否拥有全局服务权限
func checkPermissionForAgencyOnDomainURL(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "domains", domainId, "agencies", agencyId, "roles", roleId)
}

//  查询委托是否拥有项目服务权限
func checkPermissionForAgencyOnProjectURL(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "projects", projectId, "agencies", agencyId, "roles", roleId)
}

//  为委托授予全局服务权限
func removePermissionFromAgencyOnDomainURL(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "domains", domainId, "agencies", agencyId, "roles", roleId)
}

//  为委托授予项目服务权限
func removePermissionFromAgencyOnProjectURL(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) string {
	return client.ServiceURL(agencyPath, "projects", projectId, "agencies", agencyId, "roles", roleId)
}
