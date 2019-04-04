package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
)

func main() {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	//初始化服务 client
	sc, err_client := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewVPCV1 client: ", err_client)
		return
	}

	CreatePublicIP(sc)
	GetPublicIP(sc)
	UpdatePublicIP(sc)
	ListPublicIP(sc)
	DeletePublicIP(sc)
}

func CreatePublicIP(sc *gophercloud.ServiceClient) {

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
func ListPublicIP(sc *gophercloud.ServiceClient) {

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
func DeletePublicIP(sc *gophercloud.ServiceClient) {

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
