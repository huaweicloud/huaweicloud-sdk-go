package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/lifecyclehooks"
)

func main() {
	fmt.Println("main start...")
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}
	//初始化provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	//初始化服务 client
	sc, errClient := openstack.NewASV1(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewASV1 client: ", errClient)
		return
	}
	//开始测试
	TestCreateHook(sc)
	TestListHook(sc)
	TestgetHook(sc)
	TestUpdateHook(sc)
	CallBackHook(sc)
	ListhangingHook(sc)
	DeleteHook(sc)
	fmt.Println("main end...")
}

//-------------------------------------------------删除挂钩------------------------------------------------------
func DeleteHook(client *gophercloud.ServiceClient) {
	err := lifecyclehooks.Delete(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", "as-hook-gpsr").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete lifecycle Hook success!")
}

//-------------------------------------------------查询伸缩实例挂起信息------------------------------------------------------
func ListhangingHook(client *gophercloud.ServiceClient) {
	opts := lifecyclehooks.ListWithSuspensionOpts{
		InstanceId: "b41e0d49-e480-45f0-8e19-5b37bf9c153d",
	}
	result, err := lifecyclehooks.ListWithSuspension(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecycle hanging hooks: %+v\r\n", result)
	fmt.Println("Test List hanging Hook success!")
}

//--------------------------------------------------lifecyclehook回调------------------------------------------------------
func CallBackHook(client *gophercloud.ServiceClient) {
	opts := lifecyclehooks.CallBackOpts{
		InstanceId:            "5b98b904-fa3a-4405-8c4d-59df8aad1d15",
		LifecycleHookName:     "as-hook-3dvt",
		LifecycleActionResult: "ABANDON",
	}
	err := lifecyclehooks.CallBack(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test CallBackHook success!")
}

//--------------------------------------------------修改lifecyclehook------------------------------------------------------
func TestUpdateHook(client *gophercloud.ServiceClient) {
	Timeout := 300
	opts := lifecyclehooks.UpdateOpts{
		LifecycleHookType:    "INSTANCE_LAUNCHING",
		DefaultResult:        "CONTINUE",
		DefaultTimeout:       &Timeout,
		NotificationTopicUrn: "urn:smn:southchina:e428559bbe67470e8bbccb1d24073510:3",
		NotificationMetadata: "YourMetadata",
	}
	result, err := lifecyclehooks.Update(client, "31cfb1b7-823c-4358-8336-9936a8891f01", "as-SDK-test01", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("LifecycleHookName is:", result.LifecycleHookName)
	fmt.Println("LifecycleHookType is:", result.LifecycleHookType)
	fmt.Println("DefaultResult is:", result.DefaultResult)
	fmt.Println("DefaultTimeout is:", result.DefaultTimeout)
	fmt.Println("NotificationMetadata is:", result.NotificationMetadata)
	fmt.Println("NotificationTopicUrn is:", result.NotificationTopicUrn)
}

//--------------------------------------------------查询lifecyclehook--详情----------------------------------------------------
func TestgetHook(client *gophercloud.ServiceClient) {
	result, err := lifecyclehooks.Get(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", "as-hook-3dvt").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("LifecycleHookName is:", result.LifecycleHookName)
	fmt.Println("LifecycleHookType is:", result.LifecycleHookType)
	fmt.Println("DefaultResult is:", result.DefaultResult)
	fmt.Println("DefaultTimeout is:", result.DefaultTimeout)
	fmt.Println("NotificationMetadata is:", result.NotificationMetadata)
	fmt.Println("NotificationTopicUrn is:", result.NotificationTopicUrn)
}

//--------------------------------------------------查询lifecyclehook--列表----------------------------------------------------
func TestListHook(client *gophercloud.ServiceClient) {
	result, err := lifecyclehooks.List(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("lifecyclehooks list: %+v\r\n", result)
	fmt.Println("Test List Hook success!")
}

//--------------------------------------------------添加lifecyclehook------------------------------------------------------
func TestCreateHook(client *gophercloud.ServiceClient) {
	Timeout := 300
	opts := lifecyclehooks.CreateOpts{
		LifecycleHookName:    "as-SDK-test01",
		LifecycleHookType:    "INSTANCE_LAUNCHING",
		DefaultResult:        "CONTINUE",
		DefaultTimeout:       &Timeout,
		NotificationTopicUrn: "urn:smn:southchina:e428559bbe67470e8bbccb1d24073510:3",
		NotificationMetadata: "YourMetadata",
	}
	result, err := lifecyclehooks.Create(client, "31cfb1b7-823c-4358-8336-9936a8891f01", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("LifecycleHookName is:", result.LifecycleHookName)
	fmt.Println("LifecycleHookType is:", result.LifecycleHookType)
	fmt.Println("DefaultResult is:", result.DefaultResult)
	fmt.Println("DefaultTimeout is:", result.DefaultTimeout)
	fmt.Println("NotificationMetadata is:", result.NotificationMetadata)
	fmt.Println("NotificationTopicUrn is:", result.NotificationTopicUrn)
}
