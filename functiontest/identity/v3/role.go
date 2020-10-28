package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
)

func main() {
	fmt.Println(" main start--")
	// AKSK 认证
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		DomainID:         "{domainID}",
	}

	// init provider client
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Fail to get the provider: ", errAuth)
		return
	}

	sc, errClient := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NeIdentityV3 client: ", errClient)
		return
	}

	// 测试接口
	ListPermissions(sc)
	QueryPermissionDetails(sc)
	ListPermissionsForGroupOnDomain(sc)
	ListPermissionsForGroupOnProject(sc)
	GrantPermissionForGroupOnDomain(sc)
	GrantPermissionForGroupOnProject(sc)
	CheckPermissionForGroupOnProject(sc)
	CheckPermissionForGroupOnDomain(sc)
	RemovePermissionForGroupOnDomain(sc)
	RemovePermissionForGroupOnProject(sc)
	GrantPermissionToGroupOnAllProject(sc)
	fmt.Println("main end--")
}

// 查询权限列表
// Query a permission list
// GET  /v3/roles
func ListPermissions(client *gophercloud.ServiceClient) {
	opts := roles.ListOpts{
		DomainID: "",
		Name:     "",
	}
	allPages, err := roles.List(client, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode", ue.ErrCode)
				fmt.Println("ErrMessage", ue.ErrMessage)
			}
			return
		}
	}

	roles, err := roles.ExtractListRoles(allPages)
	for _, role := range roles {
		b, _ := json.MarshalIndent(role, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("TestListPermissions success!")
}

// 查询权限详情
// Query a permission detail
// GET  /v3/roles/{role_id}
func QueryPermissionDetails(client *gophercloud.ServiceClient) {
	role, err := roles.Get(client, "").ExtractGroup()

	if err != nil {
		fmt.Println(err)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode", ue.ErrCode)
				fmt.Println("ErrMessage", ue.ErrMessage)
			}
			return
		}
	}

	b, err := json.MarshalIndent(role, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestQueryPermissionDetails success!")
}

// 查询全局服务中的用户组权限
// Query permissions for group on global
// GET  /v3/domains/{domain_id}/groups/{group_id}/roles
func ListPermissionsForGroupOnDomain(client *gophercloud.ServiceClient) {
	result, err := roles.ListPermissionsForGroupOnDomain(client, "", "").ExtractList()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, role := range result.Roles {
		b, _ := json.MarshalIndent(role, "", " ")
		fmt.Println(string(b))
	}
	c, _ := json.MarshalIndent(result.Links, "", " ")
	fmt.Println(string(c))

	fmt.Println("TestListPermissionsForGroupOnDomain success!")
}

// 查询项目服务中的用户组权限
// Query permissions for group on project
// GET  /v3/projects/{project_id}/groups/{group_id}/roles
func ListPermissionsForGroupOnProject(client *gophercloud.ServiceClient) {
	result, err := roles.ListPermissionsForGroupOnProject(client, "", "").ExtractList()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, role := range result.Roles {
		b, _ := json.MarshalIndent(role, "", " ")
		fmt.Println(string(b))
	}
	c, _ := json.MarshalIndent(result.Links, "", " ")
	fmt.Println(string(c))

	fmt.Println("TestListPermissionsForGroupOnProject success!")
}

// 为用户组授予全局服务权限
// Grant permissions for group on global
// PUT  /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}
func GrantPermissionForGroupOnDomain(client *gophercloud.ServiceClient) {
	err := roles.GrantPermissionForGroupOnDomain(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestGrantPermissionForGroupOnDomain success!")
}

// 用户组授予项目服务权限
// Grant permissions for group on project
// PUT  /v3/projects/{project_id}/groups/{group_id}/roles/{role_id}
func GrantPermissionForGroupOnProject(client *gophercloud.ServiceClient) {
	err := roles.GrantPermissionForGroupOnProject(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestGrantPermissionForGroupOnProject success!")
}

// 查询用户组是否拥有项目服务权限
// Check whether the user group has the project service permission
// HEAD  /v3/projects/{project_id}/groups/{group_id}/roles/{role_id}
func CheckPermissionForGroupOnProject(client *gophercloud.ServiceClient) {
	err := roles.CheckPermissionForGroupOnProject(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestCheckPermissionForGroupOnProject success!")
}

// 查询用户组是否拥有全局服务权限
// Check whether the group has the global permission
// HEAD  /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}
func CheckPermissionForGroupOnDomain(client *gophercloud.ServiceClient) {
	err := roles.CheckPermissionForGroupOnDomain(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestCheckPermissionForGroupOnDomain success!")
}

// 移除用户组的全局服务权限
// Removing the global permission from the group
// DELETE  /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}
func RemovePermissionForGroupOnDomain(client *gophercloud.ServiceClient) {
	err := roles.RemovePermissionForGroupOnDomain(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestRemovePermissionForGroupOnDomain success!")
}

// 移除用户组的项目服务权限
// Removing the project permission from the group
// DELETE  /v3/projects/{project_id}/groups/{group_id}/roles/{role_id}
func RemovePermissionForGroupOnProject(client *gophercloud.ServiceClient) {
	err := roles.RemovePermissionForGroupOnProject(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestRemovePermissionForGroupOnProject success!")
}

// 用户组授予所有项目服务权限
// Grant all project permissions to the group.
// PUT  /v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects
func GrantPermissionToGroupOnAllProject(client *gophercloud.ServiceClient) {
	err := roles.GrantPermissionToGroupOnAllProject(client, "", "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestGrantPermissionToGroupOnAllProject success!")
}
