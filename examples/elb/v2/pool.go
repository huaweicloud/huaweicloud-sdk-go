package main

import (
	"encoding/json"
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

	poolId:=PoolCreate(sc)
	PoolList(sc)
	PoolGet(sc,poolId)
	PoolUpdate(sc,poolId)
	PoolUpdateCloseSP(sc,poolId)
	PoolDelete(sc,poolId)


	fmt.Println("main end...")
}



func PoolCreate(sc *gophercloud.ServiceClient) (poolId string)  {

	prisistenct:=pools.SessionPersistenceRequest {
		Type:			"APP_COOKIE",
		CookieName:		"test_cookie_name",
	}
	TrueValue:=true

	opts:=pools.CreateOpts{
		Name:				"kaka new",
		LBMethod:			"ROUND_ROBIN",
		Protocol:			"HTTP",
		LoadbalancerID:		"6bb85e33-4953-457a-85a9-336d76125b7b",
		Description:		"pool test",
		Persistence:		&prisistenct,
		AdminStateUp:		&TrueValue,
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

	poolId = (*resp).ID
	return poolId
}


func PoolList(sc *gophercloud.ServiceClient)(allPools []pools.Pool) {
	allPages, err := pools.List(sc,pools.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allPools, err = pools.ExtractPools(allPages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test pool List success!")
	return allPools

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
	// close session persistence
	prisistence:= &pools.SessionPersistence {
		Type:               "HTTP_COOKIE",
	}
	updatOpts:=pools.UpdateOpts{
		Name:				"KAKAK A pool",
		Description:		"LEAST_CONNECTIONS",
		LBMethod:			"LEAST_CONNECTIONS",
		Persistence:		prisistence,
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
	return resp
}

func PoolUpdateCloseSP(sc *gophercloud.ServiceClient, id string) (resp *pools.Pool) {
	// close session persistence
	prisistence:= &pools.SessionPersistence {}
	updatOpts:=pools.UpdateOpts{
		Name:				"KAKAK A pool",
		Description:		"LEAST_CONNECTIONS",
		LBMethod:			"LEAST_CONNECTIONS",
		Persistence:		prisistence,
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
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



