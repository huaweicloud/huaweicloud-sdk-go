package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/bandwidths"
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
	TestModifyPublicIPBandwidths(sc)
	fmt.Println("main end...")
}

func TestModifyPublicIPBandwidths(sc *gophercloud.ServiceClient) {

	size := 10
	//modify name
	//opts := bandwidths.UpdateOpts{
	//	Bandwidth: bandwidths.Bandwidth{
	//		Name: "fffffffffffff",
	//	},
	//}

	//modify bandwidth size
	opts := bandwidths.UpdateOpts{
		Bandwidth: bandwidths.Bandwidth{
			Name: "eeeeeeeeeeeeeeeee",
			Size: size,
		},
		ExtendParam: &bandwidths.ExtendParam{
			IsAutoPay: "true",
		},
	}

	data, err := bandwidths.Update(sc, "2a2ebbe0-a9c3-475a-b1ac-089aa435a426", opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test Update  bandwidths success!")

	if order, ok := data.(bandwidths.PrePaid); ok {
		fmt.Println("its order id ")
		fmt.Println("order id is", order.OrderID)
	}

	if on, ok := data.(bandwidths.PostPaid); ok {
		fmt.Println("its bandwidth info")
		fmt.Println("bandwidth id is ", on.ID)
		fmt.Println("bandwidth Size is ", on.Size)
		fmt.Println("bandwidth Name is ", on.Name)
		fmt.Println("bandwidth ShareType is ", on.ShareType)
		fmt.Println("bandwidth ChargeMode is ", on.ChargeMode)
		fmt.Println("bandwidth PublicipInfo is ", on.PublicipInfo)
	}

}
