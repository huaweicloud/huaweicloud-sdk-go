package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}


	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	poolId:="13a887d0-cce3-4d2a-8961-7ad855d054c9"

	memId:=MemberCreate(sc, poolId)
	MemberList(sc, poolId)
	MemberGet(sc, poolId, memId)
	MemberUpdate(sc, poolId, memId)
	MemberDelete(sc, poolId, memId)


	fmt.Println("main end...")
}



func MemberCreate(sc *gophercloud.ServiceClient, poolId string) (memId string) {

	weight:=100
	TrueValue:=true

	opts:=pools.CreateMemberOpts{
		SubnetID:"5de13914-bd0c-4387-81a7-2d6618cd4824",
		Address:"192.168.0.50",
		ProtocolPort:1234,
		Name:"kaka new",
		TenantID:"601240b9c5c94059b63d484c92cfe308",
		AdminStateUp:&TrueValue,
		Weight:&weight,
	}

	resp,err:=pools.CreateMember(sc, poolId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("member Create success!")
	memId = (*resp).ID
	return memId
}


func MemberList(sc *gophercloud.ServiceClient, poolId string) (allMembers []pools.Member) {

	allPages, err := pools.ListMembers(sc, poolId, pools.ListMembersOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	allMembers, err = pools.ExtractMembers(allPages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test member List success!")
	return allMembers

}

func MemberGet(sc *gophercloud.ServiceClient, poolId string, memId string) (resp *pools.Member) {

	resp,err:= pools.GetMember(sc, poolId, memId).Extract()

	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("member get success!")
	return resp
}

func MemberUpdate(sc *gophercloud.ServiceClient, poolId string, memId string) (resp *pools.Member) {

	TrueValue:=true
	wei := 10
	updatOpts:=pools.UpdateMemberOpts{
		Name:"KAKAK A member",
		Weight:&wei,
		AdminStateUp:&TrueValue,
	}

	resp,err:=pools.UpdateMember(sc, poolId, memId, updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("member update success!")
	return resp

}

func MemberDelete(sc *gophercloud.ServiceClient, poolId string, memId string)  {

	err:=pools.DeleteMember(sc, poolId, memId).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete member success!")
}



