package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/ptrs"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	sc, err := openstack.NewDNSV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get DNS v2 client failed")
		fmt.Println(err.Error())
		return
	}

	//TestSetupPtr(sc)
	//TestGetPtr(sc)
	TestListPtr(sc)
	//TestRestorePtr(sc)

	fmt.Println("main end...")
}

func TestGetPtr(sc *gophercloud.ServiceClient) {
	regionID := "southchina"
	fip := "bfa4116c-7347-4b54-929a-38d64be3dbd5"

	resp, err := ptrs.Get(sc, regionID, fip).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestGetPtr success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

func TestListPtr(sc *gophercloud.ServiceClient) {

	opts := ptrs.ListOpts{}

	resp, err := ptrs.List(sc, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	listptrs, err := ptrs.ExtractPtrs(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range listptrs.Floatingips {

		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListPtr success!")
}

func TestRestorePtr(sc *gophercloud.ServiceClient) {

	regionID := "southchina"
	fip := "bfa4116c-7347-4b54-929a-38d64be3dbd5"
	err := ptrs.Restore(sc, regionID, fip).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestRestorePtr success!")
}

func TestSetupPtr(sc *gophercloud.ServiceClient) {
	opts := ptrs.SetupOpts{
		Ptrdname: "www.kaksssa.com",
	}
	region := "northchina"
	fip := "bfa4116c-7347-4b54-929a-38d64be3dbd5"
	resp, err := ptrs.Setup(sc, region, fip, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestSetupPtr success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestUpdatePtr(sc *gophercloud.ServiceClient) {
	opts := ptrs.UpdateOpts{
		Ptrdname: "www.aaa.com",
	}
	region := "northchina"
	fip := "bfa4116c-7347-4b54-929a-38d64be3dbd5"
	resp, err := ptrs.Update(sc, region, fip, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestUpdatePtr success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}
