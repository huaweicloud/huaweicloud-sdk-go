package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"encoding/json"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get IAM v3 failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//TestGetRegion(sc)
	TestListRegion(sc)

	fmt.Println("main end...")
}

func TestGetRegion(client *gophercloud.ServiceClient) {

	result, err := regions.Get(client, "ccf1c29b-add0-4869-9569-1d7b47e36b39").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("regions: %+v\r\n", result)
	fmt.Println("Test Get  regions success!")
}

func TestListRegion(sc *gophercloud.ServiceClient) {

	opts := regions.ListOpts{}

	resp, err := regions.List(sc, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	regionslist, err := regions.ExtractRegions(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range regionslist {

		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListRegion success!")
}
