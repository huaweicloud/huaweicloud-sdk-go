package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
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
		fmt.Println("get network client failed")
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	ListSubnet(sc)
	CreateSubnet(sc)
	GetSubnet(sc)
	UpdateSubnet(sc)
	DeleteSubnet(sc)

	fmt.Println("main end...")
}

// List Subnet
func ListSubnet(sc *gophercloud.ServiceClient) {
	allPages, err := subnets.List(sc, subnets.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := subnets.ExtractSubnets(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("subnets: %+v\r\n", result)

	for _, resp := range result {
		fmt.Printf("subnets: %+v\r\n", resp)
		fmt.Println("subnet Id is:", resp.ID)
		fmt.Println("subnet EnableDHCP is:", resp.EnableDHCP)
		fmt.Println("subnet Name is:", resp.Name)
		fmt.Println("subnet CIDR is:", resp.CIDR)
		fmt.Println("subnet HostRoutes is:", resp.HostRoutes)
		fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
		fmt.Println("subnet IPv6AddressMode is:", resp.IPv6AddressMode)
		fmt.Println("subnet IPv6RAMode is:", resp.IPv6RAMode)
		fmt.Println("subnet NetworkID is:", resp.NetworkID)
	}
	fmt.Println("List success!")

}

// Get a Subnet
func GetSubnet(sc *gophercloud.ServiceClient) {
	resp, err := subnets.Get(sc, "xxxxxx").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("subnet: %+v\r\n", resp)
	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet EnableDHCP is:", resp.EnableDHCP)
	fmt.Println("subnet Name is:", resp.Name)
	fmt.Println("subnet CIDR is:", resp.CIDR)
	fmt.Println("subnet HostRoutes is:", resp.HostRoutes)
	fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
	fmt.Println("subnet IPv6AddressMode is:", resp.IPv6AddressMode)
	fmt.Println("subnet IPv6RAMode is:", resp.IPv6RAMode)
	fmt.Println("subnet NetworkID is:", resp.NetworkID)

	fmt.Println("Get success!")
}

// Create a Subnet
func CreateSubnet(sc *gophercloud.ServiceClient) {
	opts := subnets.CreateOpts{
		Name:      "xxxxxx",
		NetworkID: "xxxxxx",
		CIDR:      "xxx.xxx.xxx/xx",
	}

	resp, err := subnets.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("subnet: %+v\r\n", resp)
	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet EnableDHCP is:", resp.EnableDHCP)
	fmt.Println("subnet Name is:", resp.Name)
	fmt.Println("subnet CIDR is:", resp.CIDR)
	fmt.Println("subnet HostRoutes is:", resp.HostRoutes)
	fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
	fmt.Println("subnet IPv6AddressMode is:", resp.IPv6AddressMode)
	fmt.Println("subnet IPv6RAMode is:", resp.IPv6RAMode)
	fmt.Println("subnet NetworkID is:", resp.NetworkID)

	fmt.Println("Create success!")
}

// Update a Subnet
func UpdateSubnet(sc *gophercloud.ServiceClient) {
	opts := subnets.UpdateOpts{
		Name: "xxxxxx",
	}

	resp, err := subnets.Update(sc, "xxxxxx", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("subnet: %+v\r\n", resp)
	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet EnableDHCP is:", resp.EnableDHCP)
	fmt.Println("subnet Name is:", resp.Name)
	fmt.Println("subnet CIDR is:", resp.CIDR)
	fmt.Println("subnet HostRoutes is:", resp.HostRoutes)
	fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
	fmt.Println("subnet IPv6AddressMode is:", resp.IPv6AddressMode)
	fmt.Println("subnet IPv6RAMode is:", resp.IPv6RAMode)
	fmt.Println("subnet NetworkID is:", resp.NetworkID)

	fmt.Println("Update success!")
}

// Delete a Subnet
func DeleteSubnet(sc *gophercloud.ServiceClient) {
	err := subnets.Delete(sc, "xxxxxx").ExtractErr()
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
