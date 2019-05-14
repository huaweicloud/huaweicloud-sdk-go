package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/publicips"
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
	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get vpc v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	TestCreatePublicIPs(sc)
	fmt.Println("main end...")
}

func TestCreatePublicIPs(sc *gophercloud.ServiceClient) {

	size := 1
	// create on demand opts
	opts := publicips.CreateOpts{
		PublicIP: publicips.PublicIP{
			Type: "5_bgp",
		},
		Bandwidth: publicips.Bandwidth{
			Name:       "kakakondemand",
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
