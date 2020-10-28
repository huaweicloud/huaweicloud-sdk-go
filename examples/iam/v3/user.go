package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/iam/v3/users"
)

func main() {
	fmt.Println("main start...")

	// AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		DomainID:         "{domainID}",
	}

	//init provider client
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	//初始化服务 client
	sc, errClient := openstack.NewIAMV3(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewIdentityV3 client: ", errClient)
		return
	}

	// 接口测试
	QueryUserDetails(sc)
	CreateUser(sc)
	UpdateUserInfo(sc)
	UpdateUserInfoByAdmin(sc)
	fmt.Println("main end--")
}

// 查询用户详情
// Query the user detail
// GET  /v3.0/OS-USER/users/{user_id}
func QueryUserDetails(client *gophercloud.ServiceClient) {
	user, err := users.Get(client, "").ExtractGet()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestQueryUserDetails success!")
}

// 创建用户
// Create a user
// POST  /v3.0/OS-USER/users
func CreateUser(client *gophercloud.ServiceClient) {
	opts := users.CreateUserOpts{
		Areacode:    "",
		Description: "",
		DomainId:    "",
		Email:       "",
		Enabled:     nil,
		Name:        "",
		Password:    "",
		Phone:       "",
		PwdStatus:   nil,
		XuserId:     "",
		XuserType:   "",
	}

	user, err := users.Create(client, opts).ExtractCreate()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestCreateUser success!")
}

// 更新用户信息
// Update the user info
// PUT  /v3.0/OS-USER/users/{user_id}/info
func UpdateUserInfo(client *gophercloud.ServiceClient) {
	opts := users.UpdateUserInfoOpts{
		Email:  "",
		Mobile: "",
	}

	_, err := users.UpdateUserInfo(client, "", opts).ExtractUpdate()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestUpdateUserInfo success!")

}

// 更新用户信息
// Update the user info
// PUT  /v3.0/OS-USER/users/{user_id}
func UpdateUserInfoByAdmin(client *gophercloud.ServiceClient) {
	opts := users.UpdateUserOpts{
		Areacode:    "",
		Description: "",
		Email:       "",
		Enabled:     nil,
		Name:        "",
		Password:    "",
		Phone:       "",
		PwdStatus:   nil,
		XuserId:     "",
		XuserType:   "",
	}

	user, err := users.UpdateUserInfoByAdmin(client, "", opts).ExtractUpdate()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(b))

	fmt.Println("TestUpdateUserInfoByAdmin success!")
}
