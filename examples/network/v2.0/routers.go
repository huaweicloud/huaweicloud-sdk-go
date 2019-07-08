package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
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

	ListRouter(sc)
	CreateRouter(sc)
	GetRouter(sc)
	UpdateRouter(sc)
	AddInterfaceRoute(sc)
	RemoveInterfaceRoute(sc)
	DeleteRouter(sc)

	fmt.Println("main end...")
}

// List Router
func ListRouter(sc *gophercloud.ServiceClient) {
	allPages, err := routers.List(sc, routers.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := routers.ExtractRouters(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, resp := range result {

		fmt.Printf("router: %+v\r\n", resp)
		fmt.Println("router ID is:", resp.ID)
		fmt.Println("router TenantID is:", resp.TenantID)
		fmt.Println("router Name is:", resp.Name)
		fmt.Println("router AdminStateUp is:", resp.AdminStateUp)
		fmt.Println("router Status is:", resp.Status)
		fmt.Println("router AvailabilityZoneHints is:", resp.AvailabilityZoneHints)
		fmt.Println("router Distributed is:", resp.Distributed)
		fmt.Println("router GatewayInfo is:", resp.GatewayInfo)
		fmt.Println("router Routes is:", resp.Routes)
	}
	fmt.Println("List success!")

}

// Get a Router
func GetRouter(sc *gophercloud.ServiceClient) {
	result, err := routers.Get(sc, "xxxxxx").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("router: %+v\r\n", result)
	fmt.Println("router ID is:", result.ID)
	fmt.Println("router TenantID is:", result.TenantID)
	fmt.Println("router Name is:", result.Name)
	fmt.Println("router AdminStateUp is:", result.AdminStateUp)
	fmt.Println("router Status is:", result.Status)
	fmt.Println("router AvailabilityZoneHints is:", result.AvailabilityZoneHints)
	fmt.Println("router Distributed is:", result.Distributed)
	fmt.Println("router GatewayInfo is:", result.GatewayInfo)
	fmt.Println("router Routes is:", result.Routes)

	fmt.Println("Get success!")

}

// Create a Router
func CreateRouter(sc *gophercloud.ServiceClient) {
	opts := routers.CreateOpts{
		Name: "xxxxxx",
	}

	result, err := routers.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("router: %+v\r\n", result)
	fmt.Println("router ID is:", result.ID)
	fmt.Println("router TenantID is:", result.TenantID)
	fmt.Println("router Name is:", result.Name)
	fmt.Println("router AdminStateUp is:", result.AdminStateUp)
	fmt.Println("router Status is:", result.Status)
	fmt.Println("router AvailabilityZoneHints is:", result.AvailabilityZoneHints)
	fmt.Println("router Distributed is:", result.Distributed)
	fmt.Println("router GatewayInfo is:", result.GatewayInfo)
	fmt.Println("router Routes is:", result.Routes)

	fmt.Println("Create success!")
}

// Delete a Router
func DeleteRouter(sc *gophercloud.ServiceClient) {
	err := routers.Delete(sc, "xxxxxx").ExtractErr()
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

// Update a Router
func UpdateRouter(sc *gophercloud.ServiceClient) {
	opts := routers.UpdateOpts{
		Name: "xxxxxx",
	}

	result, err := routers.Update(sc, "xxxxxx", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("router: %+v\r\n", result)
	fmt.Println("router ID is:", result.ID)
	fmt.Println("router TenantID is:", result.TenantID)
	fmt.Println("router Name is:", result.Name)
	fmt.Println("router AdminStateUp is:", result.AdminStateUp)
	fmt.Println("router Status is:", result.Status)
	fmt.Println("router AvailabilityZoneHints is:", result.AvailabilityZoneHints)
	fmt.Println("router Distributed is:", result.Distributed)
	fmt.Println("router GatewayInfo is:", result.GatewayInfo)
	fmt.Println("router Routes is:", result.Routes)

	fmt.Println("Update success!")
}

// Add Interface Router
func AddInterfaceRoute(sc *gophercloud.ServiceClient) {

	routerAddOpts := routers.AddInterfaceOpts{
		SubnetID: "xxxxxx",
	}

	result, err := routers.AddInterface(sc, "xxxxxx", routerAddOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TenantID is:", result.TenantID)
	fmt.Println("ID is:", result.ID)
	fmt.Println("PortID is:", result.PortID)
	fmt.Println("SubnetID is:", result.SubnetID)

	fmt.Println("Add success!")
}

// Remove Interface Router
func RemoveInterfaceRoute(sc *gophercloud.ServiceClient) {
	routerRemoveOpts := routers.RemoveInterfaceOpts{
		SubnetID: "xxxxxx",
	}

	result, err := routers.RemoveInterface(sc, "xxxxxx", routerRemoveOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TenantID is:", result.TenantID)
	fmt.Println("ID is:", result.ID)
	fmt.Println("PortID is:", result.PortID)
	fmt.Println("SubnetID is:", result.SubnetID)

	fmt.Println("Remove success!")

}
