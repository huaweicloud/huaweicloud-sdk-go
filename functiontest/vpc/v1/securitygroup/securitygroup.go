package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/security/groups"
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
	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	TestGroupsList(sc)
	fmt.Println("main end...")
}

func TestGroupsList(sc *gophercloud.ServiceClient) {

	allPages, err := groups.List(sc, groups.ListOpts{
		Marker:"199d019f-a742-4cf6-ae75-68f78d242b2c",
		Limit:3,
	}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestGroupsList success!")
	allData, _ := groups.ExtractGroups(allPages)
	fmt.Println(len(allData))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range allData {
		p,_:=json.MarshalIndent(v,""," ")
		fmt.Println(string(p))
	}

}
