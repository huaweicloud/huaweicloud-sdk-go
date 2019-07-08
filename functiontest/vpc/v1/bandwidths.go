package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/bandwidths"
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

	UpdateBandWidth(sc)
	GetBandWidth(sc)
	ListBandWidth(sc)
	fmt.Println("main end...")
}

func UpdateBandWidth(client *gophercloud.ServiceClient) {
	result, err := bandwidths.Update(client, "00d7dedc-75aa-47aa-9c2f-f1c1f49bed19", bandwidths.UpdateOpts{
		Size: 10,
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("bandwidth: %+v\r\n", result)
	fmt.Println("bandwidth Id is:", result.ID)
	fmt.Println("bandwidth Name is:", result.Name)
	fmt.Println("bandwidth BandwidthType is:", result.BandwidthType)
	fmt.Println("bandwidth TenantId is:", result.TenantId)
	fmt.Println("bandwidth BillingInfo is:", result.BillingInfo)
	fmt.Println("bandwidth ChargeMode is:", result.ChargeMode)
	fmt.Println("bandwidth EnterpriseProjectID is:", result.EnterpriseProjectID)
	fmt.Println("bandwidth PublicIpInfo is:", result.PublicipInfo)
	fmt.Println("bandwidth ShareType is:", result.ShareType)
	fmt.Println("bandwidth Size is:", result.Size)

	fmt.Println("Update success!")
}

func GetBandWidth(client *gophercloud.ServiceClient) {
	result, err := bandwidths.Get(client, "00d7dedc-75aa-47aa-9c2f-f1c1f49bed19").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("bandwidth: %+v\r\n", result)
	fmt.Println("bandwidth Id is:", result.ID)
	fmt.Println("bandwidth Name is:", result.Name)
	fmt.Println("bandwidth BandwidthType is:", result.BandwidthType)
	fmt.Println("bandwidth TenantId is:", result.TenantId)
	fmt.Println("bandwidth BillingInfo is:", result.BillingInfo)
	fmt.Println("bandwidth ChargeMode is:", result.ChargeMode)
	fmt.Println("bandwidth EnterpriseProjectID is:", result.EnterpriseProjectID)
	fmt.Println("bandwidth PublicIpInfo is:", result.PublicipInfo)
	fmt.Println("bandwidth ShareType is:", result.ShareType)
	fmt.Println("bandwidth Size is:", result.Size)

	fmt.Println("Get success!")
}

func ListBandWidth(client *gophercloud.ServiceClient) {
	allPages, err := bandwidths.List(client, bandwidths.ListOpts{
		Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result, err1 := bandwidths.ExtractBandWidths(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("bandwidths: %+v\r\n", result)
	for _, resp := range result {
		fmt.Println("bandwidth Id is:", resp.ID)
		fmt.Println("bandwidth Name is:", resp.Name)
		fmt.Println("bandwidth BandwidthType is:", resp.BandwidthType)
		fmt.Println("bandwidth TenantId is:", resp.TenantId)
		fmt.Println("bandwidth BillingInfo is:", resp.BillingInfo)
		fmt.Println("bandwidth ChargeMode is:", resp.ChargeMode)
		fmt.Println("bandwidth EnterpriseProjectID is:", resp.EnterpriseProjectID)
		fmt.Println("bandwidth PublicIpInfo is:", resp.PublicipInfo)
		fmt.Println("bandwidth ShareType is:", resp.ShareType)
		fmt.Println("bandwidth Size is:", resp.Size)
	}

	fmt.Println("List success!")
}
