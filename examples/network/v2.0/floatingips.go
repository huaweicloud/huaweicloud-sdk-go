package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
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

	ListFloatingIP(sc)
	CreateFloatingIP(sc)
	GetFloatingIP(sc)
	UpdateFloatingIP(sc)
	DeleteFloatingIP(sc)

	fmt.Println("main end...")
}

// List FloatingIp
func ListFloatingIP(sc *gophercloud.ServiceClient) {
	allPages, err := floatingips.List(sc, floatingips.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := floatingips.ExtractFloatingIPs(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, resp := range result {
		fmt.Printf("floatingIp: %+v\r\n", resp)
		fmt.Println("floatingIp Id is:", resp.ID)
		fmt.Println("floatingIp FloatingNetworkID is:", resp.FloatingNetworkID)
		fmt.Println("floatingIp FloatingIP is:", resp.FloatingIP)
		fmt.Println("floatingIp PortID is:", resp.PortID)
		fmt.Println("floatingIp FixedIP is:", resp.FixedIP)
		fmt.Println("floatingIp TenantID is:", resp.TenantID)
		fmt.Println("floatingIp Status is:", resp.Status)
		fmt.Println("floatingIp RouterID is:", resp.RouterID)
	}

	fmt.Println("List success!")
}

// Get a FloatingIp
func GetFloatingIP(sc *gophercloud.ServiceClient) {
	result, err := floatingips.Get(sc, "xxxxxx").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("floatingIp: %+v\r\n", result)
	fmt.Println("floatingIp Id is:", result.ID)
	fmt.Println("floatingIp FloatingNetworkID is:", result.FloatingNetworkID)
	fmt.Println("floatingIp FloatingIP is:", result.FloatingIP)
	fmt.Println("floatingIp PortID is:", result.PortID)
	fmt.Println("floatingIp FixedIP is:", result.FixedIP)
	fmt.Println("floatingIp TenantID is:", result.TenantID)
	fmt.Println("floatingIp Status is:", result.Status)
	fmt.Println("floatingIp RouterID is:", result.RouterID)

	fmt.Println("Get success!")
}

// Create a FloatingIp
func CreateFloatingIP(sc *gophercloud.ServiceClient) {
	opts := floatingips.CreateOpts{
		FloatingNetworkID: "xxxxxx",
	}

	result, err := floatingips.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("floatingIp: %+v\r\n", result)
	fmt.Println("floatingIp Id is:", result.ID)
	fmt.Println("floatingIp FloatingNetworkID is:", result.FloatingNetworkID)
	fmt.Println("floatingIp FloatingIP is:", result.FloatingIP)
	fmt.Println("floatingIp PortID is:", result.PortID)
	fmt.Println("floatingIp FixedIP is:", result.FixedIP)
	fmt.Println("floatingIp TenantID is:", result.TenantID)
	fmt.Println("floatingIp Status is:", result.Status)
	fmt.Println("floatingIp RouterID is:", result.RouterID)

	fmt.Println("Create success!")
}

// Update a FloatingIp
func UpdateFloatingIP(sc *gophercloud.ServiceClient) {

	opts := floatingips.UpdateOpts{
		PortID: nil,
	}

	result, err := floatingips.Update(sc, "xxxxxx", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("floatingIp: %+v\r\n", result)
	fmt.Println("floatingIp Id is:", result.ID)
	fmt.Println("floatingIp FloatingNetworkID is:", result.FloatingNetworkID)
	fmt.Println("floatingIp FloatingIP is:", result.FloatingIP)
	fmt.Println("floatingIp PortID is:", result.PortID)
	fmt.Println("floatingIp FixedIP is:", result.FixedIP)
	fmt.Println("floatingIp TenantID is:", result.TenantID)
	fmt.Println("floatingIp Status is:", result.Status)
	fmt.Println("floatingIp RouterID is:", result.RouterID)

	fmt.Println("Update success!")
}

// Delete a FloatingIp
func DeleteFloatingIP(sc *gophercloud.ServiceClient) {
	err := floatingips.Delete(sc, "xxxxxx").ExtractErr()
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
