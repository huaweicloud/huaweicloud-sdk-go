package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
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
	ListUsers(sc)
	QueryUserDetails(sc)
	ListGroupsForUser(sc)
	CreateUser(sc)
	UpdateUserPassword(sc)
	UpdateUserInformationByAdmin(sc)
	DeleteUser(sc)
	ListUsersForGroupByAdmin(sc)
	fmt.Println("main end--")
}

// 查询用户列表
// Query a user list
// GET  /v3/users
func ListUsers(client *gophercloud.ServiceClient) {
	opts := users.UserListOpts{
		DomainID:          "",
		Enabled:           nil,
		Name:              "",
		PasswordExpiresAt: "",
	}

	allPages, err := users.ListUser(client, opts).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	extractUsers, err := users.ExtractListUsers(allPages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, user := range extractUsers {
		b, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(b))
	}
	fmt.Println("TestListUsers Success!")
}

// 查询用户详情
// Query the user detail
// GET  /v3/users/{user_id}
func QueryUserDetails(client *gophercloud.ServiceClient) {
	user, err := users.Get(client, "").ExtractDetail()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, err := json.MarshalIndent(user, "", "")
	fmt.Println(string(b))
	fmt.Println("TestQueryUserDetails Success!")
}

// 查询用户所属的用户组
// Query the group to which the User belongs
// GET  /v3/users/{user_id}/groups
func ListGroupsForUser(client *gophercloud.ServiceClient) {
	allPages, err := users.ListGroupsForUser(client, "").AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	extractGroups, err := groups.ExtractGroups(allPages)
	for _, group := range extractGroups {
		b, _ := json.MarshalIndent(group, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("TestListGroupsForUser Success!")
}

// 查询用户组中所包含的用户
// Query all users in the group
// GET  /v3/groups/{group_id}/users
func ListUsersForGroupByAdmin(client *gophercloud.ServiceClient) {
	result, err := users.ListUsersForGroupByAdmin(client, "").ExtractList()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	for _, user := range result.Users {
		b, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(b))
	}
	c, _ := json.MarshalIndent(result.Links, "", " ")
	fmt.Println(string(c))

	fmt.Println("TestListUsersForGroupByAdmin Success!")
}

// 创建用户
// Create a user
// POST  /v3/users
func CreateUser(client *gophercloud.ServiceClient) {
	opts := users.CreateUserOpts{
		Name:        "",
		Description: "",
		DomainID:    "",
		Enabled:     nil,
		Password:    "",
	}

	user, err := users.CreateUser(client, opts).ExtractCreateUser()
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestCreateUser Success!")
}

// 更新用户密码
// Update the user password
// POST  /v3/users/{user_id}/password
func UpdateUserPassword(client *gophercloud.ServiceClient) {
	opts := users.UpdatePasswordOpts{
		Password:         "",
		OriginalPassword: "",
	}

	err := users.UpdateUserPassword(client, "", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestUpdateUserPassword Success!")
}

// 更新用户信息
// Update the user info
// PATCH  /v3/users/{user_id}
func UpdateUserInformationByAdmin(client *gophercloud.ServiceClient) {
	opts := users.UpdateUserOpts{
		Name:        "",
		Description: "",
		DomainID:    "",
		Enabled:     nil,
		Password:    "",
		PwdStatus:   nil,
	}

	user, err := users.UpdateUser(client, "", opts).UpdateExtract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, err := json.MarshalIndent(user, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestUpdateUserInformationByAdmin Success!")
}

// 删除用户
// Delete the user
// DELETE  /v3/users/{user_id}
func DeleteUser(client *gophercloud.ServiceClient) {
	err := users.Delete(client, "").ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestDeleteUser Success!")
}
