package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/instances"
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
	TestListInstance(sc)
	RemoveInstance(sc)
	BathActionInstance(sc)
	fmt.Println("main end...")
}

//--------------------------------------------------查询伸缩组的实例列表------------------------------------------------------
func TestListInstance(client *gophercloud.ServiceClient) {
	opts := instances.ListOpts{
		LifeCycleState:         "INSERVICE",
		HealthStatus:           "NORMAL",
		StartNumber:            0,
		Limit:                  20,
		ProtectFromScalingDown: "true",
	}
	result, err := instances.List(client, "1c6c066b-f0f4-4312-9e98-a9663432f20b", opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List,err:= instances.ExtractInstances(result)
	for _, resp := range result1List.ScalingGroupInstances{
		fmt.Println("Instance detail:")
		fmt.Println("ScalingGroupName is:", resp.ScalingGroupName)
		fmt.Println("ScalingGroupId is:", resp.ScalingGroupId)
		fmt.Println("ScalingConfigurationId is:", resp.ScalingConfigurationId)
		fmt.Println("ScalingConfigurationName is:", resp.ScalingConfigurationName)
		fmt.Println("HealthStatus is:", resp.HealthStatus)
		fmt.Println("InstanceId is:", resp.InstanceId)
		fmt.Println("LifeCycleState is:", resp.LifeCycleState)
		fmt.Println("ProtectFromScalingDown is:", resp.ProtectFromScalingDown)
		fmt.Println("InstanceName is:", resp.InstanceName)
		fmt.Println("CreateTime is:", resp.CreateTime)
		fmt.Println("------------------------")
	}
}

//--------------------------------------------------移除实例------------------------------------------------------
func RemoveInstance(client *gophercloud.ServiceClient) {
	opts := instances.DeleteOpts{
		InstanceDelete: "no",
	}
	err := instances.Delete(client, "5b98b904-fa3a-4405-8c4d-59df8aad1d15", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Remove Instance success!")
}

//--------------------------------------------------批量操作实例------------------------------------------------------
func BathActionInstance(client *gophercloud.ServiceClient) {
	var InstancesId1  = "b41e0d49-e480-45f0-8e19-5b37bf9c153d"
	var InstancesId2  = "5b98b904-fa3a-4405-8c4d-59df8aad1d15"
	var InstancesId  = []string{InstancesId1,InstancesId2}
	opts := instances.ActionOpts{
		InstancesId:    InstancesId,
		InstanceDelete: "",
		Action:         "REMOVE",
	}
	err := instances.Action(client, "c34bf676-217c-4752-9907-5a6a0ff85748", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Bath Action Instance success!")
}
