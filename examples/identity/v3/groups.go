package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
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

	//init provider client
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
	ListGroups(sc)
	QueryGroupDetails(sc)
	CreateGroup(sc)
	UpdateGroup(sc)
	DeleteGroup(sc)
	CheckUserInGroup(sc)
	AddUserToGroup(sc)
	RemoveUserFromGroup(sc)
	fmt.Println("main end--")
}

// 查询用户组列表
// Query a group list
// GET  /v3/groups
func ListGroups(client *gophercloud.ServiceClient) {
	opts := groups.ListOpts{
		DomainID: "",
		Name:     "",
	}

	allPages, err := groups.List(client, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	groups, err := groups.ExtractGroups(allPages)

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	for _, group := range groups {
		b, _ := json.MarshalIndent(group, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("TestListGroups success!")
}

// 查询用户组详情
// Query the group detail
// GET  /v3/groups/{group_id}
func QueryGroupDetails(client *gophercloud.ServiceClient) {
	group, err := groups.Get(client, "").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	b, _ := json.MarshalIndent(group, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestQueryGroupDetails success!")
}

// 创建用户组
// Create a group
// POST  /v3/groups
func CreateGroup(client *gophercloud.ServiceClient) {
	opts := groups.CreateGroupOpts{
		Name:        "",
		Description: "",
		DomainID:    "",
	}
	group, err := groups.CreateGroup(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	b, _ := json.MarshalIndent(group, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestCreateGroup success!")
}

// 更新用户组
// Update the group info
// PATCH  /v3/groups/{group_id}
func UpdateGroup(client *gophercloud.ServiceClient) {
	opts := groups.UpdateGropupOpts{
		Name:        "",
		Description: "",
		DomainID:    "",
	}

	group, err := groups.UpdateGroup(client, "", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	b, _ := json.MarshalIndent(group, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestUpdateGroup success!")
}

// 删除用户组
// Delete the group
// DELETE  /v3/groups/{group_id}
func DeleteGroup(client *gophercloud.ServiceClient) {
	err := groups.Delete(client, "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestDeleteGroup success!")
}

// 查询用户是否在用户组
// Check whether the user is in the group
// HEAD  /v3/groups/{group_id}/users/{user_id}
func CheckUserInGroup(client *gophercloud.ServiceClient) {
	err := groups.CheckUserInGroup(client, "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestCheckUserInGroup success!")
}

// 添加用户到用户组
// Add the user to the group
// PUT  /v3/groups/{group_id}/users/{user_id}
func AddUserToGroup(client *gophercloud.ServiceClient) {
	err := groups.AddUserToGroup(client, "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestAddUserToGroup success!")
}

// 移除用户组中的用户
// Remove the user in the group
// DELETE  /v3/groups/{group_id}/users/{user_id}
func RemoveUserFromGroup(client *gophercloud.ServiceClient) {
	err := groups.RemoveUserFromGroup(client, "", "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestRemoveUserFromGroup success!")
}
