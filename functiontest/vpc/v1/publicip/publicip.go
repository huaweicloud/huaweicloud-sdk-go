package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
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
		fmt.Println("get vpc v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestCreatePublicIP(sc)
	TestGetPublicIP(sc)
	//TestUpdatePublicIP(sc)
	//TestListPublicIP(sc)
	TestDeletePublicIP(sc)

	fmt.Println("main end...")
}

func TestCreatePublicIP(sc *gophercloud.ServiceClient) {

	resp, err := publicips.Create(sc, publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type:      "5_bgp",
			IPVersion: 4,
		},
		Bandwidth: publicips.BandWidth{
			Name:      "test1t1t",
			ShareType: "PER",
			Size:      5,
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

//
func TestUpdatePublicIP(sc *gophercloud.ServiceClient) {

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
	fmt.Println("PublicIP Id is:", resp.ID)
	fmt.Println("PublicIP Status is:", resp.Status)
	fmt.Println("PublicIP BandwidthSize is:", resp.BandwidthSize)
	fmt.Println("PublicIP IPVersion is:", resp.IPVersion)
	fmt.Println("PublicIP PublicIpAddress is:", resp.PublicIpAddress)
	fmt.Println("PublicIP Type is:", resp.Type)
	fmt.Println("PublicIP CreateTime is:", resp.CreateTime)
	fmt.Println("PublicIP TenantId is:", resp.TenantId)

}

//
func TestGetPublicIP(sc *gophercloud.ServiceClient) {
	resp, err := publicips.Get(sc, "ef3a7d23-f22b-40a1-8559-b281defb768f").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

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

//
func TestListPublicIP(sc *gophercloud.ServiceClient) {

	allpages, err := publicips.List(sc, publicips.ListOpts{
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
	publicipList, err := publicips.ExtractPublicIPs(allpages)
	if err != nil {
		fmt.Println(err)
		return
	}
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
}

//
func TestDeletePublicIP(sc *gophercloud.ServiceClient) {

	resp := publicips.Delete(sc, "ef3a7d23-f22b-40a1-8559-b281defb768f")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete public ip success!")
}
