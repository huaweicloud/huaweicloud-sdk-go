package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/policylogs"
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
	TestListPolicyLogs(sc)
	fmt.Println("main end...")
}

func TestListPolicyLogs(client *gophercloud.ServiceClient) {
	id := "519d1798-9764-4465-901c-8e4a9c9934c3"
	opts := policylogs.ListOpts{
		Limit: 30,
	}
	err := policylogs.List(client, id, opts).EachPage(func(page pagination.Page) (bool, error) {

		result, err := policylogs.ExtractPolicyLogs(page)
		if err != nil {
			return false, err
		}
		b, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(b))
		fmt.Printf("TestListPolicyLogs: %+v\r\n", result)
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test List success!")
}
