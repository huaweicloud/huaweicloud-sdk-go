package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/logs"
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
	Log(sc)
	fmt.Println("main end...")
}

//--------------------------------------------------查询日志------------------------------------------------------
func Log(client *gophercloud.ServiceClient) {
	opts := logs.ListOpts{
		StartTime:   "",
		EndTime:     "",
		StartNumber: 0,
		Limit:       0,
	}
	result, err := logs.List(client, "9e209302-4dce-41a0-b1b4-c337c2cb4b6f", opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List,err:= logs.ExtractLogs(result)
	for _, resp := range result1List.ScalingActivityLog{
		fmt.Println("StartTime is:", resp.StartTime)
		fmt.Println("EndTime is:", resp.EndTime)
		fmt.Println("ID is:", resp.ID)
		fmt.Println("Description is:", resp.Description)
		fmt.Println("DesireValue is:", resp.DesireValue)
		fmt.Println("InstanceAddedList is:", resp.InstanceAddedList)
		fmt.Println("InstanceDeletedList is:", resp.InstanceDeletedList)
		fmt.Println("InstanceRemovedList is:", resp.InstanceRemovedList)
		fmt.Println("InstanceValue is:", resp.InstanceValue)
		fmt.Println("ScalingValue is:", resp.ScalingValue)
		fmt.Println("Status is:", resp.Status)
		fmt.Println("------------------------")
	}
}
