package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
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

	ListNetwork(sc)
	CreateNetwork(sc)
	GetNetwork(sc)
	UpdateNetwork(sc)
	DeleteNetwork(sc)
	GetIpUsed(sc)

	fmt.Println("main end...")
}

// List Network
func ListNetwork(sc *gophercloud.ServiceClient) {
	allPages, err := networks.List(sc, networks.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := networks.ExtractNetworks(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, resp := range result {

		fmt.Printf("network: %+v\r\n", resp)
		fmt.Println("network Id is:", resp.ID)
		fmt.Println("network Status is:", resp.Status)
		fmt.Println("network RouterExternal is:", resp.RouterExternal)
		fmt.Println("network Name is:", resp.Name)
		fmt.Println("network AdminStateUp is:", resp.AdminStateUp)
		fmt.Println("network TenantID is:", resp.TenantID)
		fmt.Println("network AvailabilityZoneHints is:", resp.AvailabilityZoneHints)
	}
	fmt.Println("List success!")

}

// Get a Network
func GetNetwork(sc *gophercloud.ServiceClient) {

	result, err := networks.Get(sc, "xxxxxx").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("network: %+v\r\n", result)
	fmt.Println("network Id is:", result.ID)
	fmt.Println("network Status is:", result.Status)
	fmt.Println("network RouterExternal is:", result.RouterExternal)
	fmt.Println("network Name is:", result.Name)
	fmt.Println("network AdminStateUp is:", result.AdminStateUp)
	fmt.Println("network TenantID is:", result.TenantID)
	fmt.Println("network AvailabilityZoneHints is:", result.AvailabilityZoneHints)

	fmt.Println("Get success!")
}

// Create a Network
func CreateNetwork(sc *gophercloud.ServiceClient) {
	opts := networks.CreateOpts{
		Name: "xxxxxx",
	}

	result, err := networks.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("network: %+v\r\n", result)
	fmt.Println("network Id is:", result.ID)
	fmt.Println("network Status is:", result.Status)
	fmt.Println("network RouterExternal is:", result.RouterExternal)
	fmt.Println("network Name is:", result.Name)
	fmt.Println("network AdminStateUp is:", result.AdminStateUp)
	fmt.Println("network TenantID is:", result.TenantID)
	fmt.Println("network AvailabilityZoneHints is:", result.AvailabilityZoneHints)

	fmt.Println("Create success!")
}

// Update a Network
func UpdateNetwork(sc *gophercloud.ServiceClient) {
	opts := networks.UpdateOpts{
		Name: "xxxxxx",
	}

	result, err := networks.Update(sc, "xxxxxx", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("network: %+v\r\n", result)
	fmt.Println("network Id is:", result.ID)
	fmt.Println("network Status is:", result.Status)
	fmt.Println("network RouterExternal is:", result.RouterExternal)
	fmt.Println("network Name is:", result.Name)
	fmt.Println("network AdminStateUp is:", result.AdminStateUp)
	fmt.Println("network TenantID is:", result.TenantID)
	fmt.Println("network AvailabilityZoneHints is:", result.AvailabilityZoneHints)

	fmt.Println("Update success!")
}

// Delete a Network
func DeleteNetwork(sc *gophercloud.ServiceClient) {
	err := networks.Delete(sc, "xxxxxx").ExtractErr()
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

// Get NetworkIpAvailabilities
func GetIpUsed(sc *gophercloud.ServiceClient) {
	networkID := "xxxxxx"

	resp, err := networks.GetNetworkIpAvailabilities(sc, networkID)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("network_id is:", resp.NetworkIpAvail.NetworkId)
	fmt.Println("network_name is:", resp.NetworkIpAvail.NetworkName)
	fmt.Println("used_ips is:", resp.NetworkIpAvail.UsedIps)
	fmt.Println("total_ips is:", resp.NetworkIpAvail.TotalIps)
	fmt.Println("subnet_ip_availability is:", resp.NetworkIpAvail.SubnetIpAvail)

	fmt.Println("Get success!")

}
