package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/quotas"
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

	sc, err := openstack.NewASV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get as v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestListQuotas(sc)
	TestListWithInstances(sc)

	fmt.Println("main end...")
}

func TestListQuotas(client *gophercloud.ServiceClient) {

	result, err := quotas.List(client).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(result, "", " ")

	fmt.Println(string(b))
	fmt.Printf("quotas: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestListWithInstances(client *gophercloud.ServiceClient) {

	id := "719ef427-3918-442a-b417-6b8761c0414e"

	result, err := quotas.ListWithInstances(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(result, "", " ")

	fmt.Println(string(b))
	fmt.Printf("ListWithInstances: %+v\r\n", result)
	fmt.Println("Test List success!")
}
