package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/publicips"
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
		fmt.Println(err)
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network vpc v2 client failed")
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	CreatePublicIPs(sc)

	fmt.Println("main end...")
}

// Create a publicIp
func CreatePublicIPs(sc *gophercloud.ServiceClient) {

	size := 1
	// create on demand opts
	opts := publicips.CreateOpts{
		PublicIP: publicips.PublicIP{
			Type: "5_bgp",
		},
		Bandwidth: publicips.Bandwidth{
			Name:       "xxxxxx",
			Size:       size,
			ShareType:  "WHOLE",
			ChargeMode: "bandwidth",
		},
	}

	//create common opts with bssprama

	//pernum := 1
	//opts := publicips.CreateOpts{
	//	PublicIP: publicips.PublicIP{
	//		Type: "5_bgp",
	//	},
	//	Bandwidth: publicips.Bandwidth{
	//		Name:       "kakakbss",
	//		Size:       size,
	//		ShareType:  "PER",
	//		ChargeMode: "bandwidth",
	//	},
	//	ExtendParam: publicips.ExtendParam{
	//		ChargeMode:  "prePaid",
	//		PeriodType:  "month",
	//		PeriodNum:   pernum,
	//		IsAutoRenew: "false",
	//		IsAutoPay:   "true",
	//	},
	//}

	data, err := publicips.Create(sc, opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if order, ok := data.(publicips.PrePaid); ok {

		fmt.Println("OrderID is:", order.OrderID)
		fmt.Println("PublicIpId is:", order.PublicipID)
	}

	if publicIp, ok := data.(publicips.PostPaid); ok {

		fmt.Println("publicIp Type is:", publicIp.Type)
		fmt.Println("publicIp TenantID is:", publicIp.TenantID)
		fmt.Println("publicIp Status is:", publicIp.Status)
		fmt.Println("publicIp ID is:", publicIp.ID)
		fmt.Println("publicIp BandwidthSize is:", publicIp.BandwidthSize)

	}

	fmt.Println("Create success!")

}
