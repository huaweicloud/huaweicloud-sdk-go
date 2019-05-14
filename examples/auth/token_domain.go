package examples

import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
)

//AuthTokenUserList using token auth method ,list users .
func AuthTokenUserList() {

	fmt.Println("main start...")

	// init token auth options
	akskOptions := token.TokenOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		DomainID:         "replace-your-domainID",
		Username:         "replace-your-username",
		Password:         "replace-your-password",
		AllowReauth:      true,
	}

	//init provider client
	provider, err := openstack.AuthenticatedClient(akskOptions)
	if err != nil {
		panic(err)
	}

	// init IAM client
	iamClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get IAM v3 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// build http request，list all users belong to this domain
	page, err := users.List(iamClient, users.ListOpts{}).AllPages()
	if err != nil {

		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// extract http response body
	userList, err := users.ExtractUsers(page)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// print result
	for _, d := range userList {

		fmt.Println("user id is :", d.ID)
		fmt.Println("user name is :", d.Name)

	}

	fmt.Println("main end...")
}
