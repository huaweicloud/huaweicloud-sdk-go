package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/publicips"
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
	client, err_client := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewVPCV2 client: ", err_client)
		return
	}

	size := 1
	// create on demand opts
	//createOpts := publicips.CreateOpts{
	//	PublicIP: publicips.PublicIP{
	//		Type: "5_bgp",
	//	},
	//	Bandwidth: publicips.Bandwidth{
	//		Name:       "kakakondemand",
	//		Size:       size,
	//		ShareType:  "WHOLE",
	//		ChargeMode: "bandwidth",
	//	},
	//}

	// create common opts with bssprama
	//
	pernum := 1
	createOpts := publicips.CreateOpts{
		PublicIP: publicips.PublicIP{
			Type: "5_bgp",
		},
		Bandwidth: publicips.Bandwidth{
			Name:       "kakakbss",
			Size:       size,
			ShareType:  "PER",
			ChargeMode: "bandwidth",
		},
		ExtendParam: publicips.ExtendParam{
			ChargeMode:  "prePaid",
			PeriodType:  "month",
			PeriodNum:   pernum,
			IsAutoRenew: "false",
			IsAutoPay:   "true",
		},
	}

	data, err := publicips.Create(client, createOpts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test publicips Create success!")

	if order, ok := data.(publicips.PrePaid); ok {
		fmt.Println("its order public ip ")
		fmt.Println(order.OrderID)
		fmt.Println(order.PublicipID)
	}

	if on, ok := data.(publicips.PostPaid); ok {
		fmt.Println("its on demand  public ip ")
		fmt.Println(on.ID)
		fmt.Println(on.Status)
	}

}
