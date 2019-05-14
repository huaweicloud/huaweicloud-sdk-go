package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/instances"
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
	TestInstanceAction(sc)
	TestListInstanceAction(sc)
	TestDelInstanceAction(sc)
	fmt.Println("main end...")
}

func TestInstanceAction(client *gophercloud.ServiceClient) {

	GroupID := "f9642b84-06c8-4c8a-aa06-7f2bdd04667f"

	opts := instances.ActionOpts{
		InstancesId: []string{"asdfasdfasdfasdfasd", "asdfasdfasdfasdfasd"},
	}

	err := instances.Action(client, GroupID, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestInstanceAction success!")
}

func TestListInstanceAction(client *gophercloud.ServiceClient) {
	GroupID := "f9642b84-06c8-4c8a-aa06-7f2bdd04667f"
	opts := instances.ListOpts{
		Limit: 1,
		StartNumber: 1,
	}
	page, err := instances.List(client, GroupID, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	resp, err := instances.ExtractInstances(page)
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

func TestDelInstanceAction(client *gophercloud.ServiceClient) {
	instanceID := "42d782f3-cb37-439c-9b73-73cf465dab4f"
	opts := instances.DeleteOpts{}
	err := instances.Delete(client, instanceID, opts).ExtractErr()
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
