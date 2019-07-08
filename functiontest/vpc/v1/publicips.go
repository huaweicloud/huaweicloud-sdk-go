package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
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
		fmt.Println("get vpc v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreatePublicIP(sc)
	GetPublicIP(sc)
	UpdatePublicIP(sc)
	ListPublicIP(sc)
	DeletePublicIP(sc)

	fmt.Println("main end...")
}

func CreatePublicIP(sc *gophercloud.ServiceClient) {

	resp, err := publicips.Create(sc, publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type: "5_bgp",
			//IPVersion: 4,
		},
		Bandwidth: publicips.BandWidth{
			Name:      "1t1t",
			ShareType: "PER",
			Size:      10,
		},
	}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("PublicIP: %+v\r\n", resp)
	fmt.Println("PublicIP Id is:", resp.ID)
	fmt.Println("PublicIP Status is:", resp.Status)
	fmt.Println("PublicIP EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("PublicIP BandwidthSize is:", resp.BandwidthSize)
	fmt.Println("PublicIP IPVersion is:", resp.IPVersion)
	fmt.Println("PublicIP PublicIpAddress is:", resp.PublicIpAddress)
	fmt.Println("PublicIP Type is:", resp.Type)
	fmt.Println("PublicIP CreateTime is:", resp.CreateTime)
	fmt.Println("PublicIP TenantId is:", resp.TenantId)
	fmt.Println("Create success!")

}

func UpdatePublicIP(sc *gophercloud.ServiceClient) {

	resp, err := publicips.Update(sc, "6ffdbd50-1425-4901-9383-09993304db61", publicips.UpdateOpts{
		IPVersion: 4,
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("PublicIP: %+v\r\n", resp)
	fmt.Println("PublicIP Id is:", resp.ID)
	fmt.Println("PublicIP Status is:", resp.Status)
	fmt.Println("PublicIP BandwidthSize is:", resp.BandwidthSize)
	fmt.Println("PublicIP IPVersion is:", resp.IPVersion)
	fmt.Println("PublicIP PublicIpAddress is:", resp.PublicIpAddress)
	fmt.Println("PublicIP Type is:", resp.Type)
	fmt.Println("PublicIP CreateTime is:", resp.CreateTime)
	fmt.Println("PublicIP TenantId is:", resp.TenantId)
	fmt.Println("Update success!")

}

func GetPublicIP(sc *gophercloud.ServiceClient) {
	resp, err := publicips.Get(sc, "ef3a7d23-f22b-40a1-8559-b281defb768f").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("PublicIP: %+v\r\n", resp)
	fmt.Println("PublicIP Id is:", resp.ID)
	fmt.Println("PublicIP Status is:", resp.Status)
	fmt.Println("PublicIP EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("PublicIP BandwidthSize is:", resp.BandwidthSize)
	fmt.Println("PublicIP IPVersion is:", resp.IPVersion)
	fmt.Println("PublicIP PublicIpAddress is:", resp.PublicIpAddress)
	fmt.Println("PublicIP Type is:", resp.Type)
	fmt.Println("PublicIP CreateTime is:", resp.CreateTime)
	fmt.Println("PublicIP TenantId is:", resp.TenantId)
	fmt.Println("Get success!")

}

func ListPublicIP(sc *gophercloud.ServiceClient) {

	allPages, err := publicips.List(sc, publicips.ListOpts{
		Limit: 10,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	publicipList, err1 := publicips.ExtractPublicIPs(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("PublicIP: %+v\r\n", publicipList)
	for _, resp := range publicipList {
		fmt.Println("PublicIP Id is:", resp.ID)
		fmt.Println("PublicIP Status is:", resp.Status)
		fmt.Println("PublicIP EnterpriseProjectId is:", resp.EnterpriseProjectId)
		fmt.Println("PublicIP BandwidthSize is:", resp.BandwidthSize)
		fmt.Println("PublicIP IPVersion is:", resp.IPVersion)
		fmt.Println("PublicIP PublicIpAddress is:", resp.PublicIpAddress)
		fmt.Println("PublicIP Type is:", resp.Type)
		fmt.Println("PublicIP CreateTime is:", resp.CreateTime)
		fmt.Println("PublicIP TenantId is:", resp.TenantId)
	}
	fmt.Println("List success!")

}

func DeletePublicIP(sc *gophercloud.ServiceClient) {

	resp := publicips.Delete(sc, "0071410e-a9c5-41be-bee0-5c69db7aba31")
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
