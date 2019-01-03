package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
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

	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestSubnetCreate(sc)
	//TestSubnetUpdate(sc)
	//TestSubnetGet(sc)
	TestSubnetDelete(sc)
	//TestSubnetList(sc)

	fmt.Println("main end...")
}

func TestSubnetList(sc *gophercloud.ServiceClient) {

	allPages, err := subnets.List(sc, subnets.ListOpts{
		VpcID: "1d79d5ce-bc4c-48c6-88cd-4a8619f6ad2c",
		Limit: 1,
	}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestSubnetList success!")

	allData, _ := subnets.ExtractSubnets(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, resp := range allData {

		fmt.Println("subnet Id is:", resp.ID)
		fmt.Println("subnet Status is:", resp.Status)
		fmt.Println("subnet Name is:", resp.Name)
		fmt.Println("subnet Cidr is:", resp.Cidr)
		fmt.Println("subnet VpcId is:", resp.VpcID)
		fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
		fmt.Println("subnet DnsList is:", resp.DNSList)
		fmt.Println("subnet DhcpEnable is:", resp.DhcpEnable)
		fmt.Println("subnet PrimaryDns is:", resp.PrimaryDNS)
		fmt.Println("subnet SecondaryDns is:", resp.SecondaryDNS)
		fmt.Println("subnet NeutronNetworkId is:", resp.NeutronNetworkID)
		fmt.Println("subnet AvailabilityZone is:", resp.AvailabilityZone)

	}

}

func TestSubnetCreate(sc *gophercloud.ServiceClient) {

	resp, err := subnets.Create(sc, subnets.CreateOpts{
		Name:      "ABC",
		Cidr:      "192.168.1.0/24",
		GatewayIP: "192.168.1.1",
		VpcID:     "20cd8567-5b6b-46d3-b270-9619069880d9",
	}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet Status is:", resp.Status)
	fmt.Println("subnet Name is:", resp.Name)
	fmt.Println("subnet Cidr is:", resp.Cidr)
	fmt.Println("subnet VpcId is:", resp.VpcID)
	fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
	fmt.Println("subnet DnsList is:", resp.DNSList)
	fmt.Println("subnet DhcpEnable is:", resp.DhcpEnable)
	fmt.Println("subnet PrimaryDns is:", resp.PrimaryDNS)
	fmt.Println("subnet SecondaryDns is:", resp.SecondaryDNS)
	fmt.Println("subnet NeutronNetworkId is:", resp.NeutronNetworkID)
	fmt.Println("subnet AvailabilityZone is:", resp.AvailabilityZone)

}

func TestSubnetUpdate(sc *gophercloud.ServiceClient) {

	resp, err := subnets.Update(sc, "1d79d5ce-bc4c-48c6-88cd-4a8619f6ad2c", "9a56640e-5503-4b8d-8231-963fc59ff91c", subnets.UpdateOpts{
		Name: "ABC-baaaaaaaaaaaack",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet Status is:", resp.Status)

}

func TestSubnetGet(sc *gophercloud.ServiceClient) {
	resp, err := subnets.Get(sc, "b767b7d6-fb42-4762-94f4-addda56d8a9a").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet Status is:", resp.Status)
	fmt.Println("subnet Name is:", resp.Name)
	fmt.Println("subnet Cidr is:", resp.Cidr)
	fmt.Println("subnet VpcId is:", resp.VpcID)
	fmt.Println("subnet GatewayIp is:", resp.GatewayIP)
	fmt.Println("subnet DnsList is:", resp.DNSList)
	fmt.Println("subnet DhcpEnable is:", resp.DhcpEnable)
	fmt.Println("subnet PrimaryDns is:", resp.PrimaryDNS)
	fmt.Println("subnet SecondaryDns is:", resp.SecondaryDNS)
	fmt.Println("subnet NeutronNetworkId is:", resp.NeutronNetworkID)
	fmt.Println("subnet AvailabilityZone is:", resp.AvailabilityZone)

}

func TestSubnetDelete(sc *gophercloud.ServiceClient) {

	resp := subnets.Delete(sc, "20cd8567-5b6b-46d3-b270-9619069880d9", "d13bdccc-bed9-44fb-920f-dfd618fee327")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete subnet success!")
}
