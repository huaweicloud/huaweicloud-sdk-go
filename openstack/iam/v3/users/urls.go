package users

import "github.com/gophercloud/gophercloud"

// 查询用户详情
func queryUserDetailUrl(client *gophercloud.ServiceClient, userId string) string {
	return client.ServiceURL("OS-USER", "users", userId)
}

// 创建用户
func createUserUrl(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("OS-USER", "users")
}

// 更新用户信息
func updateUserInfoUrl(client *gophercloud.ServiceClient, userId string) string {
	return client.ServiceURL("OS-USER", "users", userId, "info")
}

// 更新用户信息
func updateUserUrl(client *gophercloud.ServiceClient, userId string) string {
	return client.ServiceURL("OS-USER", "users", userId)
}
