package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/lifecyclehooks"
)

func main() {
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("Get privider failed")
		if err, ok := err.(gophercloud.UnifiedError); ok {
			fmt.Println("ErrorMessage:", err.Message())
			fmt.Println("ErrorCode:", err.ErrorCode())
		}
	}

	sc, err := openstack.NewASV1(provider, gophercloud.EndpointOpts{})
	TestHookCreate(sc)
	TestHookList(sc)
	TestHookGet(sc)
	TestHookUpdate(sc)
	TestHookDel(sc)
	TestHookListWithSuspend(sc)
	TestHookCallBack(sc)
}

func TestHookCreate(client *gophercloud.ServiceClient) {
	opts := lifecyclehooks.CreateOpts{
		LifecycleHookName:    "TEST",
		LifecycleHookType:    "INSTANCE_TERMINATING",
		NotificationTopicUrn: "urn:smn:southchina:128a7bf965154373a7b73c89eb6b65aa:newtest",
	}

	res, err := lifecyclehooks.Create(client, "c5c7636d-2567-41dd-ad22-62326b8e1dd8", opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecyclehooks: %+v\r\n", res)
	fmt.Println("Test Create success!")

}

func TestHookList(client *gophercloud.ServiceClient) {
	res, err := lifecyclehooks.List(client, "c5c7636d-2567-41dd-ad22-62326b8e1dd8").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecyclehooks: %+v\r\n", res)
	fmt.Println("Test List success!")

}

func TestHookGet(client *gophercloud.ServiceClient) {
	res, err := lifecyclehooks.Get(client, "f9642b84-06c8-4c8a-aa06-7f2bdd04667f", "as-hook-z1be").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecyclehooks: %+v\r\n", res)
	fmt.Println("Test Get success!")

}

func TestHookUpdate(client *gophercloud.ServiceClient) {
	opts := lifecyclehooks.UpdateOpts{LifecycleHookType: "INSTANCE_LAUNCHING", DefaultResult: "ABANDON"}
	res, err := lifecyclehooks.Update(client, "c5c7636d-2567-41dd-ad22-62326b8e1dd8", "TEST", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecyclehooks: %+v\r\n", res)
	fmt.Println("Test Update success!")

}

func TestHookDel(client *gophercloud.ServiceClient) {
	err := lifecyclehooks.Delete(client, "f9642b84-06c8-4c8a-aa06-7f2bdd04667f", "as-hook-ordm").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test Del success!")

}

func TestHookListWithSuspend(client *gophercloud.ServiceClient) {
	instanceid := lifecyclehooks.ListWithSuspensionOpts{InstanceId: "72c2d432-fd27-4460-8dc0-530c7dff8958"}
	res, err := lifecyclehooks.ListWithSuspension(client, "c5c7636d-2567-41dd-ad22-62326b8e1dd8", instanceid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(res.InstanceHangingInfo)
	fmt.Println("Test List With Suspend success!")

}

func TestHookCallBack(client *gophercloud.ServiceClient) {
	opts := lifecyclehooks.CallBackOpts{InstanceId: "72c2d432-fd27-4460-8dc0-530c7dff8958", LifecycleActionResult: "ABANDON"}
	err := lifecyclehooks.CallBack(client, "c5c7636d-2567-41dd-ad22-62326b8e1dd8", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test List With CallBack success!")

}
