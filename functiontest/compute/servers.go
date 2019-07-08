package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
)

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestServerCreate(sc)
	TestServerListV226(sc)
	TestServerList(sc)
	//TestServerListBrief(sc)
	//TestServerGet(sc)
	//TestServerUpdate(sc)
	//TestServerDelete(sc)
	TestServerListInstanceActions(sc)
	TestServerGetInstanceActions(sc)
	TestServerGetConsoleLog(sc)

	fmt.Println("main end...")
}
func TestServerCreate(sc *gophercloud.ServiceClient) {
	baseOpts := servers.CreateOpts{
		Name:      "ECS_xx2",
		FlavorRef: "c1.xlarge",
		Networks: []servers.Network{
			servers.Network{UUID: "9a56640e-5503-4b8d-8231-963fc59ff91c"},
		},
		AvailabilityZone: "az1.dc1",
	}

	bd := []bootfromvolume.BlockDevice{
		bootfromvolume.BlockDevice{
			BootIndex:       0,
			DestinationType: "volume",
			SourceType:      "image",
			VolumeSize:      40,
			UUID:            "ee5c7dc8-acb8-4d93-8d47-b27610b3477d",
		},
	}

	sh := schedulerhints.SchedulerHints{
		CheckResources: "true",
	}

	bsOpts := bootwithscheduler.CreateOptsExt{
		CreateOptsBuilder: baseOpts,
		BlockDevice:       bd,
		SchedulerHints:    sh,
	}

	resp, err := servers.Create(sc, bsOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test server create success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestServerListV226(sc *gophercloud.ServiceClient) {
	sc.SetMicroversion("2.26")
	defer sc.UnsetMicroversion()
	allPages, err := servers.List(sc, servers.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test server List success!")

	allData, err := servers.ExtractServers(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range allData {
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}
}
func TestServerList(sc *gophercloud.ServiceClient) {
	allPages, err := servers.List(sc, servers.ListOpts{Limit: 2}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test server List success!")

	allData, err := servers.ExtractServers(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

}

func TestServerListBrief(sc *gophercloud.ServiceClient) {
	allPages, err := servers.ListBrief(sc, servers.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test server List brief success!")

	allData, err := servers.ExtractServers(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestServerGet(sc *gophercloud.ServiceClient) {
	id := "73a9f3f2-c30a-4941-9ad1-85d7e95e1295"

	resp, err := servers.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestServerUpdate(sc *gophercloud.ServiceClient) {
	id := "9efe661a-7c36-48d8-9df0-0ec5e049d726"

	updatOpts := servers.UpdateOpts{
		Name:        "KAKAK server",
		Description: "new kaka server",
	}

	resp, err := servers.Update(sc, id, updatOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestServerDelete(sc *gophercloud.ServiceClient) {
	id := "1a414482-f70c-4327-8142-e172ab821c4b"
	err := servers.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test server delete success!")
}

func TestServerListInstanceActions(sc *gophercloud.ServiceClient) {
	id := "4a5e7286-8da5-4bf8-9658-88f5fc604d2e"

	resp, err := servers.ListInstanceActions(sc, id).ExtractInstanceActionsListResult()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get ServerActions success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestServerGetInstanceActions(sc *gophercloud.ServiceClient) {
	id := "4a5e7286-8da5-4bf8-9658-88f5fc604d2e"
	reqId := "req-d21d1fee-3fae-4691-bead-bee607700000"
	resp, err := servers.GetInstanceActions(sc, id, reqId).ExtractInstanceActionsResult()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get ServerActionsByRequestId success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestServerGetConsoleLog(sc *gophercloud.ServiceClient) {
	id := "a447418b-7d79-4f9d-82e1-17380b78ee2c"

	resp,err := servers.GetConsoleLog(sc, id,"20").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println(resp)
	fmt.Println("Test server get TestServerGetConsoleLog success!")

}