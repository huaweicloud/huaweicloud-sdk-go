package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"

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

	//TestPolicyList(sc)
	//TestPolicyCreate(sc)
	//TestPolicyGet(sc)
	//TestPolicyUpdate(sc)
	TestPolicyDelete(sc)

	fmt.Println("main end...")
}

func TestPolicyCreate(sc *gophercloud.ServiceClient) {

	opts := policies.CreateOpts{

		Name:           "asd",
		RedirectPoolID: "3a412129-863e-430e-a03a-aa6c66a7827e",
		ListenerID:     "11d633e3-ea9f-4e14-bd87-2d2866d347fd",
		Action:         "REDIRECT_TO_POOL",
	}

	resp, err := policies.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("policy Create success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestPolicyList(sc *gophercloud.ServiceClient) {
	allPages, err := policies.List(sc, policies.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test policy List success!")

	allData, _ := policies.ExtractPolcies(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestPolicyGet(sc *gophercloud.ServiceClient) {

	id := "8d5f4b3b-2ffc-47d4-8b94-b73eb36fd61b"

	resp, err := policies.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("policy get success!", resp)

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestPolicyUpdate(sc *gophercloud.ServiceClient) {

	id := "b03cbb53-27af-4660-9b44-b56356d9320b"

	updatOpts := policies.UpdateOpts{
		Description: "asdddddddddddddd",
	}

	resp, err := policies.Update(sc, id, updatOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("policy update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestPolicyDelete(sc *gophercloud.ServiceClient) {

	id := "b03cbb53-27af-4660-9b44-b56356d9320b"
	err := policies.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete policy success!")
}
