package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/policies"
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
	TestCreatePolicy(sc)
	TestGetPolicy(sc)
	TestListPolicy(sc)
	TestUpdatePolicy(sc)
	TestDeletePolicy(sc)
	TestPolicyAction(sc)

	fmt.Println("main end...")
}
func TestCreatePolicy(client *gophercloud.ServiceClient) {

	var num = 100
	opts := policies.CreateOpts{
		ScalingGroupId:    "719ef427-3918-442a-b417-6b8761c0414e",
		ScalingPolicyName: "kaka",
		ScalingPolicyType: "SCHEDULED",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime: "2019-03-15T03:02Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:      "ADD",
			InstanceNumber: &num,
		},
		CoolDownTime: &num,
	}

	result, err := policies.Create(client, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("policies: %+v\r\n", result)
	fmt.Println("Test Create success!")
}

func TestGetPolicy(client *gophercloud.ServiceClient) {

	id := "abcfab21-db13-4064-8a4c-bccc3637efbf"

	result, err := policies.Get(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("policies: %+v\r\n", result)
	fmt.Println("Test Get success!")
}

func TestListPolicy(client *gophercloud.ServiceClient) {
	id := "c5c7636d-2567-41dd-ad22-62326b8e1dd8"
	opts := policies.ListOpts{
		Limit: 30,
	}

	err := policies.List(client, id, opts).EachPage(func(page pagination.Page) (bool, error) {

		result, err := policies.ExtractPolicies(page)
		if err != nil {
			return false, err
		}
		b, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(b))
		fmt.Printf("policies: %+v\r\n", result)
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

func TestUpdatePolicy(client *gophercloud.ServiceClient) {
	id := "672740a3-2f67-4489-ad3c-545852cdc1d3"
	opts := policies.UpdateOpts{
		ScalingPolicyName: "asdfasd",
	}

	result, err := policies.Update(client, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(result, "", " ")

	fmt.Println(string(b))
	fmt.Printf("policies: %+v\r\n", result)
	fmt.Println("Test TestUpdatePolicy success!")
}

func TestPolicyAction(client *gophercloud.ServiceClient) {
	id := "672740a3-2f67-4489-ad3c-545852cdc1d3"

	opts := policies.ActionOpts{
		Action: "execute",
	}

	err := policies.Action(client, id, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestPolicyAction success!")
}

func TestDeletePolicy(client *gophercloud.ServiceClient) {
	id := "266399c3-30fe-432e-9ae0-49ce4e83e6c7"

	err := policies.Delete(client, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestDeletePolicy success!")
}
