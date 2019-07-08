package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/notifications"
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
		fmt.Println("Failed to get the AS client: ", errClient)
		return
	}
	//开始测试
	CreateNotification(sc)
	ListNotification(sc)
	DeleteNotification(sc)
	fmt.Println("main end...")
}

//------------------配置通知---------------------------------------------
func CreateNotification(client *gophercloud.ServiceClient) {
	var TopicScene = []string{"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN"}
	opts := notifications.ConfigNotificationOpts{
		TopicUrn:   "urn:smn:southchina:e428559bbe67470e8bbccb1d24073510:3",
		TopicScene: TopicScene,
	}
	result, err := notifications.ConfigNotification(client, "04ffce3e-2424-49ef-859d-885e85ee1fde", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("TopicName is:", result.TopicName)
	fmt.Println("TopicScene is:", result.TopicScene)
	fmt.Println("TopicUrn is:", result.TopicUrn)
}

//------------------查询通知列表---------------------------------------------
func ListNotification(client *gophercloud.ServiceClient) {
	result, err := notifications.List(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("Notification list: %+v\r\n", result)
	fmt.Println("Test List Notification success!")
}
//------------------删除查询通知---------------------------------------------
func DeleteNotification(client *gophercloud.ServiceClient) {
	err := notifications.Delete(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", "urn:smn:southchina:1e914c4db904409996ad23c07c94a31b:test_topic5").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Delete Notification success!")
}
