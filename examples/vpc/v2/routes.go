package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/routes"
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
		fmt.Println(err)
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println(err)
		fmt.Println("get network vpc v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreateRoute(sc)
	GetRoute(sc)
	ListRoute(sc)
	DeleteRoute(sc)
	fmt.Println("main end...")
}

// Create a Route
func CreateRoute(client *gophercloud.ServiceClient) {

	result, err := routes.Create(client, routes.CreateOpts{
		Type:        "peering",
		Nexthop:     "xxxxxx",
		Destination: "xxx.xxx.xxx.xxx/xx",
		VpcID:       "xxxxxx",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("routes: %+v\r\n", result)

	fmt.Println("Route Id is:", result.ID)
	fmt.Println("Route TenantID is:", result.TenantID)
	fmt.Println("Route Destination is:", result.Destination)
	fmt.Println("Route Nexthop is:", result.Nexthop)
	fmt.Println("Route Type is:", result.Type)
	fmt.Println("Route VpcID is:", result.VpcID)

	fmt.Println("Create success!")
}

// Get a Route
func GetRoute(client *gophercloud.ServiceClient) {

	result, err := routes.Get(client, "xxxxxx").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("routes: %+v\r\n", result)

	fmt.Println("Route Id is:", result.ID)
	fmt.Println("Route TenantID is:", result.TenantID)
	fmt.Println("Route Destination is:", result.Destination)
	fmt.Println("Route Nexthop is:", result.Nexthop)
	fmt.Println("Route Type is:", result.Type)
	fmt.Println("Route VpcID is:", result.VpcID)

	fmt.Println("Get success!")
}

// List Routes
func ListRoute(client *gophercloud.ServiceClient) {

	allPages, err := routes.List(client, routes.ListOpts{
		Limit: 30,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result, err := routes.ExtractRoutes(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, resp := range result {

		fmt.Println("Route Id is:", resp.ID)
		fmt.Println("Route TenantID is:", resp.TenantID)
		fmt.Println("Route Destination is:", resp.Destination)
		fmt.Println("Route Nexthop is:", resp.Nexthop)
		fmt.Println("Route Type is:", resp.Type)
		fmt.Println("Route VpcID is:", resp.VpcID)
	}
	fmt.Printf("Route: %+v\r\n", result)
	fmt.Println("List success!")
}

// Delete a Route
func DeleteRoute(client *gophercloud.ServiceClient) {
	err := routes.Delete(client, "xxxxxx").ExtractErr()
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
