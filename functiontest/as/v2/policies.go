package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v2/policies"
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

	sc, err := openstack.NewASV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get as v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestCreatePolicy(sc)
	//TestUpdatePolicy(sc)
	//TestGetPolicy(sc)
	TestListPolicy(sc)
	TestListPolicyByResourceID(sc)

	fmt.Println("main end...")
}
func TestCreatePolicy(client *gophercloud.ServiceClient) {
	var num = 10
	opts := policies.CreateOpts{
		ScalingResourceId:   "719ef427-3918-442a-b417-6b8761c0414e",
		ScalingResourceType: "SCALING_GROUP",
		ScalingPolicyName:   "kaka",
		ScalingPolicyType:   "SCHEDULED",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime: "2019-03-15T03:02Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation: "ADD",
			Size:      &num,
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

	id := "4c8a7c5d-78e6-4fa1-835f-fa0387175569"

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
	opts := policies.ListOpts{
		Limit: 30,
	}

	err := policies.List(client, opts).EachPage(func(page pagination.Page) (bool, error) {

		result, err := policies.ExtractPolicies(page)
		if err != nil {
			return false, err
		}
		b, _ := json.MarshalIndent(result, "", " ")
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

func TestListPolicyByResourceID(client *gophercloud.ServiceClient) {
	id := "719ef427-3918-442a-b417-6b8761c0414e"
	opts := policies.ResourceListOpts{
		Limit: 30,
	}

	err := policies.GetPolicyListByResourceID(client, id, opts).EachPage(func(page pagination.Page) (bool, error) {

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
	id := "4c8a7c5d-78e6-4fa1-835f-fa0387175569"
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
