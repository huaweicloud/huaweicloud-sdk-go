package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
)

func main() {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	//初始化服务 client
	client, err_client := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewVPCV1 client: ", err_client)
		return
	}

	//初始化过滤参数
	listOpts := subnets.ListOpts{
		Limit: 1,
	}
	//根据过滤条件获取子网列表
	allPages, err := subnets.List(client, listOpts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//解析所有返回值
	allSubnets, err_extract := subnets.ExtractSubnets(allPages)
	if err_extract != nil {
		fmt.Println("Unable to extract subnets: ", err_extract)
		return
	}
	fmt.Println("Succeed to list subnets!")
	if len(allSubnets) > 0 {
		fmt.Println("First subnets ID is:", allSubnets[0].ID)
	}
}
