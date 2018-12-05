package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
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

	poolId:=PoolCreate(sc)
	PoolList(sc)
	PoolGet(sc,poolId)
	PoolUpdate(sc,poolId)
	PoolDelete(sc,poolId)


	fmt.Println("main end...")
}



func PoolCreate(sc *gophercloud.ServiceClient) (poolId string)  {

	prisistenct:=pools.SessionPersistenceRequest { Type:"HTTP_COOKIE"}
	TrueValue:=true

	opts:=pools.CreateOpts{
		Name:"kaka new",
		LBMethod:"ROUND_ROBIN",
		Protocol:"HTTP",
		LoadbalancerID:"165b6a38-5278-4569-b747-b2ee65ea84a4",
		Description:"pool test",
		Persistence:&prisistenct,
		AdminStateUp:&TrueValue,
	}

	resp,err:=pools.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("pool Create success!")
	poolId = (*resp).ID
	return poolId
}


func PoolList(sc *gophercloud.ServiceClient)(allPages pagination.Page)  {
	allPages, err := pools.List(sc,pools.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test pool List success!")
	return allPages

}


func PoolGet(sc *gophercloud.ServiceClient, id string) (resp *pools.Pool)  {

	resp,err:= pools.Get(sc,id).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("pool get success!")
	return resp
}

func PoolUpdate(sc *gophercloud.ServiceClient, id string) (resp *pools.Pool) {

	prisistenct:=pools.SessionPersistence {"APP_COOKIE","test_cookie"}
	updatOpts:=pools.UpdateOpts{
		Name:"KAKAK A pool",
		Description:"LEAST_CONNECTIONS",
		LBMethod:"LEAST_CONNECTIONS",
		Persistence:&prisistenct,
	}

	resp,err:=pools.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("pool update success!")
	return resp
}

func PoolDelete(sc *gophercloud.ServiceClient, id string)  {

	err:=pools.Delete(sc,id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete pool success!")
}



