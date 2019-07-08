package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
)

func main() {

	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreateSubnet(sc)
	UpdateSubnet(sc)
	GetSubnet(sc)
	DeleteSubnet(sc)
	ListSubnet(sc)

	fmt.Println("main end...")
}

func ListSubnet(sc *gophercloud.ServiceClient) {

	allPages, err := subnets.List(sc, subnets.ListOpts{
		//VpcID: "71969f1a-2527-4751-ab7f-94a435b8eacc",
		//Limit: 1,
	}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allData, err1 := subnets.ExtractSubnets(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("subnets: %+v\r\n", allData)
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
	fmt.Println("List success!")

}

func CreateSubnet(sc *gophercloud.ServiceClient) {

	resp, err := subnets.Create(sc, subnets.CreateOpts{
		Name:      "ABC",
		Cidr:      "192.168.1.0/24",
		GatewayIP: "192.168.1.1",
		VpcID:     "4bf9a3fc-ad22-44aa-8097-30b2c272b0d2",
	}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("subnet: %+v\r\n", resp)
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
	fmt.Println("Create success!")

}

func UpdateSubnet(sc *gophercloud.ServiceClient) {

	resp, err := subnets.Update(sc, "83587a60-6f68-4beb-84e4-4360df4bbd49", "0aed2b44-b846-4d15-8047-84bbc6aaf46c", subnets.UpdateOpts{
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
	fmt.Printf("subnet: %+v\r\n", resp)
	fmt.Println("subnet Id is:", resp.ID)
	fmt.Println("subnet Status is:", resp.Status)
	fmt.Println("Update success!")

}

func GetSubnet(sc *gophercloud.ServiceClient) {
	resp, err := subnets.Get(sc, "008ce66f-ff4a-430c-ae7f-d9959ebcde00").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("subnet: %+v\r\n", resp)
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
	fmt.Println("Get success!")

}

func DeleteSubnet(sc *gophercloud.ServiceClient) {

	resp := subnets.Delete(sc, "83587a60-6f68-4beb-84e4-4360df4bbd49", "0aed2b44-b846-4d15-8047-84bbc6aaf46c")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete success!")
}
