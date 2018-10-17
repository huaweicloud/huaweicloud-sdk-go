package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"

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

	//TestListenerCreate(sc)
	//TestListenerGet(sc)
	//TestListenerUpdate(sc)
	//TestListenerDelete(sc)
	TestListenerList(sc)
	//

	fmt.Println("main end...")
}

func TestListenerCreate(sc *gophercloud.ServiceClient) {

	opts := listeners.CreateOpts{

		Name:           "new listener",
		Description:    "AAAAAAA",
		Protocol:       listeners.ProtocolHTTP,
		ProtocolPort:   8801,
		LoadbalancerID: "9eb9ef27-25d1-45d3-b860-84791d97f328",
	}

	resp, err := listeners.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("listener Create success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestListenerList(sc *gophercloud.ServiceClient) {
	allPages, err := listeners.List(sc, listeners.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test ListenerList success!")

	allData, _ := listeners.ExtractListeners(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestListenerGet(sc *gophercloud.ServiceClient) {

	id := "11d633e3-ea9f-4e14-bd87-2d2866d347fd"

	resp, err := listeners.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("listener get success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestListenerUpdate(sc *gophercloud.ServiceClient) {

	id := "1e5114d8-99bd-444d-8f17-4664ec9fc866"

	updatOpts := listeners.UpdateOpts{
		Name: "KAKAK A listener",
	}

	resp, err := listeners.Update(sc, id, updatOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("listener update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestListenerDelete(sc *gophercloud.ServiceClient) {

	id := "1e5114d8-99bd-444d-8f17-4664ec9fc866"
	err := listeners.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete listener success!")
}
