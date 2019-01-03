package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
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
	TestGetIpUsed(sc)
	fmt.Println("main end...")
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

	fmt.Println("network_id is ", resp.NetworkIpAvail.NetworkId)
	fmt.Println("network_name is ", resp.NetworkIpAvail.NetworkName)
	fmt.Println("used_ips is ", resp.NetworkIpAvail.UsedIps)
	fmt.Println("total_ips is ", resp.NetworkIpAvail.TotalIps)

	fmt.Println("used_ips is ", resp.NetworkIpAvail.SubnetIpAvail.UsedIps)
	fmt.Println("total_ips is ", resp.NetworkIpAvail.SubnetIpAvail.TotalIps)
	fmt.Println("subnet_id is ", resp.NetworkIpAvail.SubnetIpAvail.SubnetId)
	fmt.Println("subnet_name is ", resp.NetworkIpAvail.SubnetIpAvail.SubnetName)
}
