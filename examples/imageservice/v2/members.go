package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/members"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
)

func main() {
	//设置认证参数
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//初始化provider client
	provider, err := openstack.AuthenticatedClient(tokenOpts)
	if err != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", err)
		return
	}
	//初始化service client
	sc, clientErr := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})

	if clientErr != nil {
		fmt.Println("Failed to get the NewImageServiceV2 client: ", clientErr)
		return
	}

	memberCreate(sc)
	memberGet(sc)
	memberDelete(sc)
	memberUpdate(sc)
	memberDelete(sc)

}

func memberCreate(sc *gophercloud.ServiceClient) {
	//添加成员
	memberId := ""
	imageId := ""
	member, err := members.Create(sc, imageId, memberId).Extract()
	fmt.Printf("err: %s", err)
	if ue, ok := err.(*gophercloud.UnifiedError); ok {
		fmt.Println("ErrCode:", ue.ErrorCode())
		fmt.Println("Message:", ue.Message())
	}

	fmt.Println("member status is :", member.Status)
	fmt.Println("member ID is :", member.MemberID)
	return
}

func memberGet(sc *gophercloud.ServiceClient) {
	//查询成员
	imageId := ""
	memberId := ""
	member, err := members.Get(sc, imageId, memberId).Extract()
	fmt.Printf("err: %s", err)
	if ue, ok := err.(*gophercloud.UnifiedError); ok {
		fmt.Println("ErrCode:", ue.ErrorCode())
		fmt.Println("Message:", ue.Message())
	}

	fmt.Println("member status is :", member.Status)
	fmt.Println("member ID is :", member.MemberID)
	return
}

func memberUpdate(sc *gophercloud.ServiceClient) {
	//更新成员状态

	imageId := ""
	memberId := ""

	updateOpts := members.UpdateOpts{
		Status: "accepted",
	}

	member, err := members.Update(sc, imageId, memberId, updateOpts).Extract()
	fmt.Printf("err: %s", err)
	if ue, ok := err.(*gophercloud.UnifiedError); ok {
		fmt.Println("ErrCode:", ue.ErrorCode())
		fmt.Println("Message:", ue.Message())
	}

	fmt.Println("member status is :", member.Status)
	fmt.Println("member ID is :", member.MemberID)
	return
}

func memberDelete(sc *gophercloud.ServiceClient) {
	// 删除成员
	imageId := ""
	memberId := ""

	err := members.Delete(sc, imageId, memberId).ExtractErr()
	fmt.Printf("err: %s", err)
	if ue, ok := err.(*gophercloud.UnifiedError); ok {
		fmt.Println("ErrCode:", ue.ErrorCode())
		fmt.Println("Message:", ue.Message())
	}
	return
}
