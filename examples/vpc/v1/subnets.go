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
		Cloud:            "yyy.com",
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
		//VpcID: "xxxxxx",
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
		Name:      "xxxxxx",
		Cidr:      "xxx.xxx.xxx.xxx/xx",
		GatewayIP: "xxx.xxx.xxx.xxx",
		VpcID:     "xxxxxx",
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

	resp, err := subnets.Update(sc, "xxxxxx", "xxxxxx", subnets.UpdateOpts{
		Name: "xxxxxx",
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
	resp, err := subnets.Get(sc, "xxxxxx").Extract()
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

	resp := subnets.Delete(sc, "xxxxxx", "xxxxxx")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Delete success!")
}
