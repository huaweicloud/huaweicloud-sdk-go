package custompolicy

import "github.com/gophercloud/gophercloud"

const (
	customPolicyPath = "OS-ROLE"
	rolesPath        = "roles"
)

// 查询自定义策略列表
func listCustomPoliciesUrl(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(customPolicyPath, rolesPath)
}

// 查询自定义策略详情
func queryAgencyDetailsUrl(client *gophercloud.ServiceClient, roleId string) string {
	return client.ServiceURL(customPolicyPath, rolesPath, roleId)
}

// 创建云服务自定义策略-创建委托自定义策略
func createCustomPolicyURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(customPolicyPath, rolesPath)
}

// 修改云服务自定义策略-修改委托自定义策略
func updateCustomPolicyURL(client *gophercloud.ServiceClient, roleId string) string {
	return client.ServiceURL(customPolicyPath, rolesPath, roleId)
}

// 删除委托
func deleteCustomPolicyURL(client *gophercloud.ServiceClient, roleId string) string {
	return client.ServiceURL(customPolicyPath, rolesPath, roleId)
}
