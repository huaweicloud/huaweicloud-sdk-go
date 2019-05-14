package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/tags"
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

	TestListResourceTag(sc)
	TestUpdateTag(sc)
	TestListTenantTag(sc)
	TestListInstanceTag(sc)
	fmt.Println("main end...")
}

func TestListInstanceTag(client *gophercloud.ServiceClient) {

	resourceType := "scaling_group_tag"
	opts := tags.InstanceOpts{
		Action: "filter",
	}
	result, err := tags.ListInstanceTags(client, resourceType, opts).Extract()

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
	fmt.Printf("InstanceTag: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestListResourceTag(client *gophercloud.ServiceClient) {

	resourceID := "719ef427-3918-442a-b417-6b8761c0414e"
	resourceType := "scaling_group_tag"

	result, err := tags.ListResourceTags(client, resourceType, resourceID).Extract()

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
	fmt.Printf("ResourceTag: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestListTenantTag(client *gophercloud.ServiceClient) {

	resourceType := "scaling_group_tag"

	result, err := tags.ListTenantTags(client, resourceType).Extract()

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
	fmt.Printf("TenantTag: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestUpdateTag(client *gophercloud.ServiceClient) {
	resourceID := "719ef427-3918-442a-b417-6b8761c0414e"
	resourceType := "scaling_group_tag"

	var tagList []tags.Tag

	p := append(tagList, tags.Tag{
		Key:   "ka",
		Value: "zhou",
	}, tags.Tag{
		Key:   "ka1",
		Value: "zhou1",
	}, )

	opts := tags.UpdateOpts{
		Tags:   p,
		Action: "create",
	}
	err := tags.Update(client, resourceType, resourceID, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update tags success!")
}
