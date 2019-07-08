package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/privateips"
)

func main() {
	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreatePrivateIp(sc)
	GetPrivateIp(sc)
	ListPrivateIp(sc)
	DeletePrivateIp(sc)
	fmt.Println("main end...")
}
func CreatePrivateIp(client *gophercloud.ServiceClient) {
	result, err := privateips.Create(client, privateips.CreateOpts{
		Privateips: []privateips.PrivateIpCreate{
			{
				SubnetId: "xxxxxx",
				//IpAddress: "xxx.xxx.xxx.xxx",
			},
		},
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("privateips: %+v\r\n", result)
	for _, resp := range *result {
		fmt.Println("PrivateIps Id is:", resp.ID)
		fmt.Println("PrivateIps Status is:", resp.Status)
		fmt.Println("PrivateIps DeviceOwner is:", resp.DeviceOwner)
		fmt.Println("PrivateIps IpAddress is:", resp.IpAddress)
		fmt.Println("PrivateIps SubnetId is:", resp.SubnetId)
		fmt.Println("PrivateIps TenantId is:", resp.TenantId)
	}

}

func GetPrivateIp(client *gophercloud.ServiceClient) {
	result, err := privateips.Get(client, "xxxxxx").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("privateips: %+v\r\n", result)
	fmt.Println("PrivateIps Id is:", result.ID)
	fmt.Println("PrivateIps Status is:", result.Status)
	fmt.Println("PrivateIps DeviceOwner is:", result.DeviceOwner)
	fmt.Println("PrivateIps IpAddress is:", result.IpAddress)
	fmt.Println("PrivateIps SubnetId is:", result.SubnetId)
	fmt.Println("PrivateIps TenantId is:", result.TenantId)
	fmt.Println("Get success!")
}

func ListPrivateIp(client *gophercloud.ServiceClient) {
	subnetID := "xxxxxx"
	allPages, err := privateips.List(client, subnetID, privateips.ListOpts{
		Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result, err1 := privateips.ExtractPrivateIps(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("privateips: %+v\r\n", result)
	for _, resp := range result {
		fmt.Println("PrivateIps Id is:", resp.ID)
		fmt.Println("PrivateIps Status is:", resp.Status)
		fmt.Println("PrivateIps DeviceOwner is:", resp.DeviceOwner)
		fmt.Println("PrivateIps IpAddress is:", resp.IpAddress)
		fmt.Println("PrivateIps SubnetId is:", resp.SubnetId)
		fmt.Println("PrivateIps TenantId is:", resp.TenantId)
	}
	fmt.Println("List success!")
}

func DeletePrivateIp(client *gophercloud.ServiceClient) {
	err := privateips.Delete(client, "xxxxxx").ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Delete success!")
}
