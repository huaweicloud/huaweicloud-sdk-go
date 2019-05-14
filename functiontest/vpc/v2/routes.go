package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/routes"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network vpc v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestCreateRoute(sc)
	TestGetRoute(sc)
	TestListRoute(sc)
	TestDeleteRoute(sc)
	fmt.Println("main end...")
}
func TestCreateRoute(client *gophercloud.ServiceClient) {

	result, err := routes.Create(client, routes.CreateOpts{
		Type:        "peering",
		Nexthop:     "9d7ab42f-015e-4688-ad5a-d38bc97c045a",
		Destination: "192.168.1.0/24",
		VpcID:       "90de16ce-3bd5-42e8-a6b3-d275d26ceb33",
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
	fmt.Println("Test Create success!")
}

func TestGetRoute(client *gophercloud.ServiceClient) {

	id := "c311a64a-049a-4054-a307-d76f1f8ea62f"

	result, err := routes.Get(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("routes: %+v\r\n", result)
	fmt.Println("Test Get success!")
}

func TestListRoute(client *gophercloud.ServiceClient) {

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
	fmt.Printf("Route: %+v\r\n", result)
	fmt.Println("Test List success!")
}

func TestDeleteRoute(client *gophercloud.ServiceClient) {
	err := routes.Delete(client, "b194c615-8e14-4353-bb11-78c1d2d4b4cf").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete success!")
}
