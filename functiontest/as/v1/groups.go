package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/groups"
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
	TestCreateASGroup(sc)
	TestListAsGroup(sc)
	TestUpdateAsGroup(sc)
	TestGetAsGroup(sc)
	TestDelAsGroup(sc)
	fmt.Println("main end...")
}

func TestCreateASGroup(client *gophercloud.ServiceClient) {

	net := groups.Network{
		ID: "6df498a2-3480-4faf-b6e7-ac25a053bbbc",
	}

	opts := groups.CreateOpts{
		ScalingGroupName: "KAKAnew",
		VpcId:            "90de16ce-3bd5-42e8-a6b3-d275d26ceb33",
		Networks:         []groups.Network{net},
	}

	result, err := groups.Create(client, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("groups: %+v\r\n", result)
	fmt.Println("Test Create success!")
}

func TestListAsGroup(client *gophercloud.ServiceClient) {
	opts := groups.ListOpts{
		Limit: 1,
		//StartNumber: 1,
	}
	page, err := groups.List(client, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	resp, err := groups.ExtractGroups(page)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test List success!")
}

func TestListAsGroupEach(client *gophercloud.ServiceClient) {
	opts := groups.ListOpts{
		Limit:       1,
		StartNumber: 1,
	}

	err := groups.List(client, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, err := groups.ExtractGroups(page)
		if err != nil {
			return false, err
		}
		b, _ := json.MarshalIndent(resp, "", " ")
		fmt.Println(string(b))
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
func TestGetAsGroup(client *gophercloud.ServiceClient) {

	groupID := "f9642b84-06c8-4c8a-aa06-7f2bdd04667f"
	result, err := groups.Get(client, groupID).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("groups: %+v\r\n", result)
	fmt.Println("Test Get success!")
}

func TestUpdateAsGroup(client *gophercloud.ServiceClient) {

	groupID := "2558d0b9-4ed0-4ec3-a603-449cfce36322"
	opts := groups.UpdateOpts{
		ScalingGroupName: "myName",
	}
	result, err := groups.Update(client, groupID, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("groups: %+v\r\n", result)
	fmt.Println("Test update success!")

}

func TestDelAsGroup(client *gophercloud.ServiceClient) {
	groupID := "42d782f3-cb37-439c-9b73-73cf465dab4f"
	opts := groups.DeleteOpts{
		ForceDelete: "no",
	}
	err := groups.Delete(client, groupID, opts).ExtractErr()
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
