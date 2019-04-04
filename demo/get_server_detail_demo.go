package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
)

func main()  {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	//初始化服务 client
	client, err_client := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewECSV1 client: ", err_client)
		return
	}

	//调用的后台ecs v1接口，获取server详情
	serverId := "17e8af52-22d7-4f5e-b10b-816a2d260842"
	resp,err:=cloudserversext.GetServerExt(client,serverId)
	if err!=nil{
		fmt.Printf("Failed to get the server detail of %s,Error: %s",serverId,err.Error())
		return
	}

	//打印server部分信息
	fmt.Println("CloudServer ID is:",resp.CloudServer.ID)
	fmt.Println("CloudServer Name is:",resp.CloudServer.Name)
	fmt.Println("CloudServer VMState is:",resp.CloudServer.VMState)

	fmt.Println("CloudServer Flavor ID is:",resp.CloudServer.Flavor.ID)
	fmt.Println("CloudServer Flavor Name is:",resp.CloudServer.Flavor.Name)
	fmt.Println("CloudServer Flavor RAM is:",resp.CloudServer.Flavor.RAM)
	fmt.Println("CloudServer Flavor Disk is:",resp.CloudServer.Flavor.Disk)
	fmt.Println("CloudServer Flavor Vcpus is:",resp.CloudServer.Flavor.Vcpus)

	fmt.Println("CloudServer Metadata ChargingMode is:",resp.CloudServer.Metadata.ChargingMode)
	fmt.Println("CloudServer Metadata OrderID is:",resp.CloudServer.Metadata.OrderID)
	fmt.Println("CloudServer Metadata VpcID is:",resp.CloudServer.Metadata.VpcID)
	fmt.Println("CloudServer Metadata ImageID is:",resp.CloudServer.Metadata.ImageID)
	fmt.Println("CloudServer Metadata Imagetype is:",resp.CloudServer.Metadata.Imagetype)

	for k,addresses := range resp.CloudServer.Addresses{
		for _,v:=range addresses{
			fmt.Printf("CloudServer Addresses %s : Addr is %s,MacAddr is %s\n",k,v.Addr,v.MacAddr)
		}
	}

	for _,v:=range resp.CloudServer.VolumeAttached{
		fmt.Println("CloudServer VolumeAttached:",v.ID)
	}
}