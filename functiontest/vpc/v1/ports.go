package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/ports"
)

func main() {
	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
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
	CreatePort(sc)
	UpdatePort(sc)
	GetPort(sc)
	ListPort(sc)
	DeletePort(sc)
	fmt.Println("main end...")
}
func CreatePort(client *gophercloud.ServiceClient) {

	result, err := ports.Create(client, ports.CreateOpts{
		Name:      "kakport",
		NetworkId: "dd7ab438-f0ed-468a-8d29-0f8522ad132a",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("port: %+v\r\n", result)
	fmt.Println("port Id is:", result.ID)
	fmt.Println("port Status is:", result.Status)
	fmt.Println("port DeviceOwner is:", result.DeviceOwner)
	fmt.Println("port Name is:", result.Name)
	fmt.Println("port AdminStateUp is:", result.AdminStateUp)
	fmt.Println("port TenantId is:", result.TenantId)
	fmt.Println("port FixedIps is:", result.FixedIps)

	fmt.Println("Create success!")
}

func UpdatePort(client *gophercloud.ServiceClient) {

	result, err := ports.Update(client, "a96ba238-4b7f-4928-8aac-f23d81b1d111", ports.UpdateOpts{
		Name: "ModifiedPort",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("port: %+v\r\n", result)
	fmt.Println("port Id is:", result.ID)
	fmt.Println("port Status is:", result.Status)
	fmt.Println("port DeviceOwner is:", result.DeviceOwner)
	fmt.Println("port Name is:", result.Name)
	fmt.Println("port AdminStateUp is:", result.AdminStateUp)
	fmt.Println("port TenantId is:", result.TenantId)
	fmt.Println("port FixedIps is:", result.FixedIps)
	fmt.Println("Update success!")
}

func GetPort(client *gophercloud.ServiceClient) {

	result, err := ports.Get(client, "016ed520-3cfb-4e2c-919d-ce5eb1071aac").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("port: %+v\r\n", result)
	fmt.Println("port Id is:", result.ID)
	fmt.Println("port Status is:", result.Status)
	fmt.Println("port DeviceOwner is:", result.DeviceOwner)
	fmt.Println("port Name is:", result.Name)
	fmt.Println("port AdminStateUp is:", result.AdminStateUp)
	fmt.Println("port TenantId is:", result.TenantId)
	fmt.Println("port FixedIps is:", result.FixedIps)
	fmt.Println("Get success!")
}

func ListPort(client *gophercloud.ServiceClient) {

	allPages, err := ports.List(client, ports.ListOpts{
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
	result, err1 := ports.ExtractPorts(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("port: %+v\r\n", result)
	for _, resp := range result {
		fmt.Println("port Id is:", resp.ID)
		fmt.Println("port Status is:", resp.Status)
		fmt.Println("port DeviceOwner is:", resp.DeviceOwner)
		fmt.Println("port Name is:", resp.Name)
		fmt.Println("port AdminStateUp is:", resp.AdminStateUp)
		fmt.Println("port TenantId is:", resp.TenantId)
		fmt.Println("port FixedIps is:", resp.FixedIps)
	}
	fmt.Println("List success!")
}

func DeletePort(client *gophercloud.ServiceClient) {
	err := ports.Delete(client, "a96ba238-4b7f-4928-8aac-f23d81b1d111").ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Delete success!")
}
