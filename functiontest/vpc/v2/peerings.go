package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/peerings"
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

	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network vpc v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestCreatePeering(sc)
	//TestGetPeering(sc)
	TestListPeering(sc)
	TestDeletePeering(sc)
	//
	//TestUpdatePeering(sc)
	//TestAcceptPeering(sc)
	//TestRejectPeering(sc)

	fmt.Println("main end...")
}
func TestCreatePeering(client *gophercloud.ServiceClient) {

	opts := peerings.CreateOpts{
		Name: "KAKAnew",
		RequestVpcInfo: peerings.VPCInfo{
			VpcID:    "d06ee937-4a98-4007-9542-83576ab1464e",
			TenantID: "128a7bf965154373a7b73c89eb6b65aa",
		},
		AcceptVpcInfo: peerings.VPCInfo{
			VpcID:    "f3a99f52-20d7-4dc0-852c-994804f30c51",
			TenantID: "f8b60643bb8e44349b75da40923cbcd3",
		},
	}

	result, err := peerings.Create(client, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peerings: %+v\r\n", result)
	fmt.Println("Test Create success!")
}

func TestGetPeering(client *gophercloud.ServiceClient) {

	id := "9d7ab42f-015e-4688-ad5a-d38bc97c045a"

	result, err := peerings.Get(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peerings: %+v\r\n", result)
	fmt.Println("Test Get success!")
}

func TestListPeering(client *gophercloud.ServiceClient) {

	result, err := peerings.List(client, peerings.ListOpts{
		Limit: 30,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	peerList, err := peerings.ExtractPeerings(result)

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(peerList, "", " ")

	fmt.Println(string(b))
	fmt.Printf("peerings: %+v\r\n", peerList)
	fmt.Println("Test List success!")
}

func TestDeletePeering(client *gophercloud.ServiceClient) {
	id:="e7d4b677-ba19-4678-b20f-8440ed2e36b7"
	err := peerings.Delete(client, id).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete success!")
}

func TestUpdatePeering(client *gophercloud.ServiceClient) {

	id := "9d7ab42f-015e-4688-ad5a-d38bc97c045a"
	opts := peerings.UpdateOpts{
		Name: "kaka",
	}

	result, err := peerings.Update(client, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peerings: %+v\r\n", result)
	fmt.Println("Test update success!")
}

func TestAcceptPeering(client *gophercloud.ServiceClient) {

	id := "5b131746-d727-4867-86fd-f3c70cbd4c88"

	result, err := peerings.Accept(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peerings: %+v\r\n", result)
	fmt.Println("Test TestAcceptPeering success!")
}
func TestRejectPeering(client *gophercloud.ServiceClient) {

	id := "e7d4b677-ba19-4678-b20f-8440ed2e36b7"

	result, err := peerings.Reject(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peerings: %+v\r\n", result)
	fmt.Println("Test TestRejectPeering success!")
}
