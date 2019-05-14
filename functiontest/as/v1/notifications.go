package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/notifications"
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

	TestEnableNotification(sc)
	TestDeleteNotification(sc)
	TestListNotification(sc)

	fmt.Println("main end...")
}

func TestListNotification(client *gophercloud.ServiceClient) {
	id := "cfdd14a3-e24d-414c-a3ce-b191386b9b83"

	result, err := notifications.List(client, id).Extract()

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
	fmt.Printf("Notification: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestEnableNotification(client *gophercloud.ServiceClient) {
	id := "cfdd14a3-e24d-414c-a3ce-b191386b9b83"
	opts := notifications.ConfigNotificationOpts{
		TopicScene: []string{"SCALING_UP_FAIL"},
		TopicUrn:   "urn:smn:southchina:128a7bf965154373a7b73c89eb6b65aa:newtest",
	}

	result, err := notifications.ConfigNotification(client, id, opts).Extract()

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
	fmt.Printf("Notification: %+v\r\n", result)
	fmt.Println("Test TestEnableNotification success!")
}

func TestDeleteNotification(client *gophercloud.ServiceClient) {
	id := "f9642b84-06c8-4c8a-aa06-7f2bdd04667f"

	topicUrn := "urn:smn:southchina:128a7bf965154373a7b73c89eb6b65aa:newtest"

	err := notifications.Delete(client, id, topicUrn).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestDeleteNotification success!")
}
