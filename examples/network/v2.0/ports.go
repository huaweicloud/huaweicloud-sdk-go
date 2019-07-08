package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
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
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get Network client failed")
		fmt.Println(err)
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

// List Port
func ListPort(sc *gophercloud.ServiceClient) {
	allPages, err := ports.List(sc, ports.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := ports.ExtractPorts(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, resp := range result {
		fmt.Printf("port: %+v\r\n", resp)
		fmt.Println("port Id is:", resp.ID)
		fmt.Println("port Status is:", resp.Status)
		fmt.Println("port DeviceOwner is:", resp.DeviceOwner)
		fmt.Println("port Name is:", resp.Name)
		fmt.Println("port AdminStateUp is:", resp.AdminStateUp)
		fmt.Println("port MACAddress is:", resp.MACAddress)
		fmt.Println("port FixedIPs is:", resp.FixedIPs)
	}
	fmt.Println("List success!")

}

// Get a Port
func GetPort(sc *gophercloud.ServiceClient) {
	result, err := ports.Get(sc, "xxxxxx").Extract()
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
	fmt.Println("port MACAddress is:", result.MACAddress)
	fmt.Println("port FixedIPs is:", result.FixedIPs)

	fmt.Println("Get success!")
}

// Create a Port
func CreatePort(sc *gophercloud.ServiceClient) {
	opts := ports.CreateOpts{
		Name:      "xxxxxx",
		NetworkID: "xxxxxx",
	}

	result, err := ports.Create(sc, opts).Extract()
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
	fmt.Println("port MACAddress is:", result.MACAddress)
	fmt.Println("port FixedIPs is:", result.FixedIPs)

	fmt.Println("Create success!")
}

// Update a Port
func UpdatePort(sc *gophercloud.ServiceClient) {
	opts := ports.UpdateOpts{
		Name: "testport2",
	}

	result, err := ports.Update(sc, "xxxxxx", opts).Extract()
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
	fmt.Println("port MACAddress is:", result.MACAddress)
	fmt.Println("port FixedIPs is:", result.FixedIPs)

	fmt.Println("Update success!")
}

// Delete a Port
func DeletePort(sc *gophercloud.ServiceClient) {
	err := ports.Delete(sc, "xxxxxx").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Delete success!")
}
