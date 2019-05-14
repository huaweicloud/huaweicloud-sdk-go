package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/logs"
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

	TestListLogs(sc)

	fmt.Println("main end...")
}

func TestListLogs(client *gophercloud.ServiceClient) {
	id := "c5c7636d-2567-41dd-ad22-62326b8e1dd8"

	opts := logs.ListOpts{
		Limit: 2,
	}
	err := logs.List(client, id, opts).EachPage(func(page pagination.Page) (bool, error) {

		result, err := logs.ExtractLogs(page)
		if err != nil {
			return false, err
		}

		b, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(b))
		fmt.Printf("logs: %+v\r\n", result)

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

	fmt.Println("Test TestListLogs success!")
}
