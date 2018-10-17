package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/nat"
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

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//ELB
	TestPoolsListMembers(sc)
	//TestPoolsGetMember(sc)
	//TestPoolsUpdateMember(sc)
	//TestPoolsCreateMember(sc)
	//TestPoolsDeleteMember(sc)
	//TestNatGatwayCreate(sc)
	//TestGetNetworkIpAvailabilities(sc)

	//TestPoolsListPorts(sc)
	fmt.Println("main end...")
}



func TestNatGatwayCreate(sc *gophercloud.ServiceClient){


	opts:=nat.CreateOpts{
		Description:"thisistest",
		Name:"newgateway",
		Spec:"1",
		RouterId:"2342342134",
		InternalNetworkId:"3d34r23erd",
	}

	resp,err:=nat.CreateNatGateway(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools Create nat success!")
	fmt.Println("member:", resp)
}








func TestPoolsListPorts(sc *gophercloud.ServiceClient) {
	allPages, err := ports.List(sc, &ports.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Ports List success!")

	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range allPorts {
		fmt.Println("p:", p)
	}
}

func TestPoolsListMembers(sc *gophercloud.ServiceClient) {
	allPages, err := pools.ListMembers(sc, "3a412129-863e-430e-a03a-aa6c66a7827e", &pools.ListMembersOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools List Members success!")

	allMembers, err := pools.ExtractMembers(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range allMembers {
		fmt.Println("member.ID:", m.ID)
		fmt.Println("member.PoolID:", m.PoolID)
	}
}

func TestPoolsGetMember(sc *gophercloud.ServiceClient) {
	member, err := pools.GetMember(sc, "82fb38a1-65de-4bc6-b2ba-4b9e5d53acaa",
		"3d7961f4-dbdf-459e-9fb6-99fdd36b777f").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools Get Member success!")
	fmt.Println("member:", member)
}

func TestPoolsUpdateMember(sc *gophercloud.ServiceClient) {
	//adminStateUp := true

	opts := pools.UpdateMemberOpts{
		Weight: 5,
		//AdminStateUp: &adminStateUp,
	}

	member, err := pools.UpdateMember(sc, "3a412129-863e-430e-a03a-aa6c66a7827e",
		"2b395f9e-0c01-4041-8a6c-89d29b7cc6e2", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools Update Member success!")
	fmt.Println("member:", member)
}

func TestPoolsCreateMember(sc *gophercloud.ServiceClient) {
	weight := 0
	opts := pools.CreateMemberOpts{
		Address:      "192.168.1.38",
		ProtocolPort: 87,
		SubnetID:     "20b8a44b-e724-4103-8233-f70c7aa1bbc2",
		Name:         "member-xx2",
		Weight:       &weight,
	}

	member, err := pools.CreateMember(sc, "82fb38a1-65de-4bc6-b2ba-4b9e5d53acaa", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools Create Member success!")
	fmt.Println("member:", member)
}

func TestPoolsDeleteMember(sc *gophercloud.ServiceClient) {
	err := pools.DeleteMember(sc, "82fb38a1-65de-4bc6-b2ba-4b9e5d53acaa", "5ce64556-be25-4684-9b27-e52efbd4098e").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Pools Delete Member success!")
}

func TestGetNetworkIpAvailabilities(sc *gophercloud.ServiceClient) {
	iu, err := networks.GetNetworkIpAvailabilities(sc, "5689bda8-767d-4029-9c9e-460c3e05f46a")
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("get network ip Availabilities success!")
	fmt.Println("UsedIps:", iu.NetworkIpAvail.UsedIps)
	fmt.Println("TotalIps:", iu.NetworkIpAvail.TotalIps)
}
