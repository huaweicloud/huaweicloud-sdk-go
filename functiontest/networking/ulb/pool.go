package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"

	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
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

	//TestPoolList(sc)
	TestPoolCreate(sc)
	//TestPoolGet(sc)
	//TestPoolUpdate(sc)
	//TestPoolDelete(sc)

	fmt.Println("main end...")
}

func TestPoolCreate(sc *gophercloud.ServiceClient) {

	opts := pools.CreateOpts{
		Name:           "kaka new",
		LBMethod:       "ROUND_ROBIN",
		Protocol:       "UDP",
		LoadbalancerID: "fd18b88e-e75f-46d6-984e-753eb56d7b17",
	}

	resp, err := pools.Create(sc, opts).Extract()

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

}

func TestPoolList(sc *gophercloud.ServiceClient) {
	allPages, err := pools.List(sc, pools.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test pool List success!")

	allData, _ := pools.ExtractPools(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestPoolGet(sc *gophercloud.ServiceClient) {

	id := "7036109a-4099-4437-a4bd-7001e7f40f79"

	resp, err := pools.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("pool get success!", resp)

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestPoolUpdate(sc *gophercloud.ServiceClient) {

	id := "7036109a-4099-4437-a4bd-7001e7f40f79"

	updatOpts := pools.UpdateOpts{
		Name:     "KAKAK A pool",
		LBMethod: "ROUND_ROBIN",
	}

	resp, err := pools.Update(sc, id, updatOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("pool update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestPoolDelete(sc *gophercloud.ServiceClient) {

	id := "356b6933-3add-4502-a40b-203dd871c973"
	err := pools.Delete(sc, id).ExtractErr()

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
