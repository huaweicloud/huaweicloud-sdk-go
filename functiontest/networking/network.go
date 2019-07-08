package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

var networkid string

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

	TestNetworkList(sc)
	TestNetworkCreate(sc)
	TestNetworkGet(sc)
	TestNetworkUpdate(sc)
	TestNetworkDelete(sc)
	TestGetIpUsed(sc)

	fmt.Println("main end...")
}

func TestNetworkList(sc *gophercloud.ServiceClient) {
	allpages, err := networks.List(sc, networks.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	networks, err := networks.ExtractNetworks(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get network list success!")
	p, _ := json.MarshalIndent(networks, "", " ")
	fmt.Println(string(p))
}

func TestNetworkGet(sc *gophercloud.ServiceClient) {

	network, err := networks.Get(sc, networkid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get network detail success!")
	p, _ := json.MarshalIndent(network, "", " ")
	fmt.Println(string(p))
}

func TestNetworkCreate(sc *gophercloud.ServiceClient) {
	opts := networks.CreateOpts{
		Name: "testnetwork",
	}

	network, err := networks.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create network success!")
	networkid = network.ID
	p, _ := json.MarshalIndent(network, "", " ")
	fmt.Println(string(p))
}

func TestNetworkUpdate(sc *gophercloud.ServiceClient) {
	opts := networks.UpdateOpts{
		Name: "testnetwork2",
	}

	network, err := networks.Update(sc, networkid, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update network success!")
	p, _ := json.MarshalIndent(network, "", " ")
	fmt.Println(string(p))
}

func TestNetworkDelete(sc *gophercloud.ServiceClient) {
	err := networks.Delete(sc, networkid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete network success!")
}

func TestGetIpUsed(sc *gophercloud.ServiceClient) {
	networkID := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	resp, err := networks.GetNetworkIpAvailabilities(sc, networkID)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("network Id is:", resp.NetworkIpAvail.NetworkId)
	fmt.Println("network Name is:", resp.NetworkIpAvail.NetworkName)
	fmt.Println("used_ips is:", resp.NetworkIpAvail.UsedIps)
	fmt.Println("total_ips is:", resp.NetworkIpAvail.TotalIps)
	fmt.Println("NetworkIpAvail.SubnetIpAvail is:", resp.NetworkIpAvail.SubnetIpAvail)

	fmt.Println("Get success!")

}
