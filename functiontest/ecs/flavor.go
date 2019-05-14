package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/flavor"
)

func main()  {
	fmt.Println("main start...")

	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	sc, err := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ecs v1 client failed")
		fmt.Println(err.Error())
		return
	}

	TestFlavorList(sc)
	TestResize(sc)

	fmt.Println("main end...")
}

func TestFlavorList(sc *gophercloud.ServiceClient) {
	allPages, err := flavor.List(sc,&flavor.ListOpts{AvailabilityZone:"az1.dc1(obt)"}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	flavors ,err := flavor.ExtractFlavors(allPages)
	if err != nil {
		fmt.Println("Test get flavor list error:",err)
		return
	}

	b, _ := json.MarshalIndent(flavors, "", "   ")
	fmt.Println(string(b))
}

func TestResize(sc *gophercloud.ServiceClient) {
	opts := flavor.ResizeOpts{
		FlavorRef:"c3.15xlarge.2",
		DedicatedHostId:"459a2b9d-804a-4745-ab19-a113bb1b4ddc",
	}

	jobId, err := flavor.Resize(sc,"1cdd5621-93b7-4c49-be6a-500229c196f2",opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test resize success,jobId is: ",jobId)
}