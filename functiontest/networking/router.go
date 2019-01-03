package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

var routerid string
var subnetid string

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestRouterList(sc)
	TestRouterCreate(sc)
	TestRouterGet(sc)
	TestRouterUpdate(sc)
	TestRouteAddInterface(sc)
	TestRouteRemoveInterface(sc)
	TestRouterDelete(sc)

	fmt.Println("main end...")
}

func TestRouterList(sc *gophercloud.ServiceClient) {
	allpages, err := routers.List(sc, routers.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	routers, err := routers.ExtractRouters(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get router list success!")
	p, _ := json.MarshalIndent(routers, "", " ")
	fmt.Println(string(p))
}

func TestRouterGet(sc *gophercloud.ServiceClient) {
	router, err := routers.Get(sc, routerid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get router detail success!")
	p, _ := json.MarshalIndent(router, "", " ")
	fmt.Println(string(p))
}

func TestRouterCreate(sc *gophercloud.ServiceClient) {
	opts := routers.CreateOpts{
		Name: "routertest",
	}

	router, err := routers.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Get create router success!")
	routerid = router.ID
	p, _ := json.MarshalIndent(router, "", " ")
	fmt.Println(string(p))
}

func TestRouterDelete(sc *gophercloud.ServiceClient) {
	err := routers.Delete(sc, routerid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete router success!")
}

func TestRouterUpdate(sc *gophercloud.ServiceClient) {
	opts := routers.UpdateOpts{
		Name: "routertest2",
	}

	router, err := routers.Update(sc, routerid, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update router success!")
	p, _ := json.MarshalIndent(router, "", " ")
	fmt.Println(string(p))
}

func TestRouteAddInterface(sc *gophercloud.ServiceClient) {
	opts := subnets.CreateOpts{
		Name:      "testsubnet",
		NetworkID: "021d431d-d430-41fc-a6df-9ce50b9e8169",
		CIDR:      "192.168.1.0/24",
	}

	subnet, err := subnets.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	subnetid = subnet.ID
	routeropts := routers.AddInterfaceOpts{
		SubnetID: subnetid,
	}

	interfaceInfo, err := routers.AddInterface(sc, routerid, routeropts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test router add interface success!")
	p, _ := json.MarshalIndent(interfaceInfo, "", " ")
	fmt.Println(string(p))
}

func TestRouteRemoveInterface(sc *gophercloud.ServiceClient) {
	opts := routers.RemoveInterfaceOpts{
		SubnetID: subnetid,
	}

	interfaceInfo, err := routers.RemoveInterface(sc, routerid, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test router remove interface success!")
	p, _ := json.MarshalIndent(interfaceInfo, "", " ")
	fmt.Println(string(p))
}
