package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/policylogs"
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
	gophercloud.EnableDebug = true

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
	PoliceLog(sc)
	fmt.Println("main end...")
}

//--------------------查询策略日志---------------------------------------------------------------
func PoliceLog(client *gophercloud.ServiceClient) {
	opts := policylogs.ListOpts{
		LogId:               "e6280917-47cc-4be0-9859-b48ff50009d5",
		ScalingResourceType: "BANDWIDTH",
		ScalingResourceId:   "3793aa0b-e59a-43ba-b398-33698826fe7b",
		ExecuteType:         "SCHEDULED",
		StartTime:           "2019-02-20T09:50:03Z",
		EndTime:             "2019-05-20T09:50:03Z",
		StartNumber:         0,
		Limit:               0,
	}
	result, err := policylogs.List(client, "5cba0719-d424-46a7-9053-82ee643c3e69", opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List,err:= policylogs.ExtractPolicyLogs(result)
	fmt.Println("TotalNumber:",result1List.TotalNumber)
	for _, resp := range result1List.ScalingPolicyExecuteLog{
		fmt.Println("PolicyExecuteLog detail:")
		fmt.Println("ScalingPolicyId is:", resp.ScalingPolicyId)
		fmt.Println("Status is:", resp.Status)
		fmt.Println("DesireValue is:", resp.DesireValue)
		fmt.Println("ID is:", resp.ID)
		fmt.Println("ScalingResourceId is:", resp.ScalingResourceId)
		fmt.Println("Type is:", resp.Type)
		fmt.Println("ExecuteTime is:", resp.ExecuteTime)
		fmt.Println("ExecuteType is:", resp.ExecuteType)
		fmt.Println("FailedReason is:", resp.FailedReason)
		fmt.Println("LimitValue is:", resp.LimitValue)
		fmt.Println("MetaData is:", resp.MetaData)
		fmt.Println("OldValue is:", resp.OldValue)
		fmt.Println("ScalingResourceType is:", resp.ScalingResourceType)
		fmt.Println("TenantId is:", resp.TenantId)
		fmt.Println("------------------------")
	}
}
