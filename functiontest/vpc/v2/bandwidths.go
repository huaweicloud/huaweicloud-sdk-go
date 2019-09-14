package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/bandwidths"
	"encoding/json"
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
	TestCreatePublicIPBandwidths(sc)
	TestBatchCreatePublicIPBandwidths(sc)
	TestInsertPublicIPBandwidths(sc)
	TestRemovePublicIPBandwidths(sc)
	TestDeletePublicIPBandwidths(sc)
	TestModifyPublicIPBandwidths(sc)
	fmt.Println("main end...")
}

func TestCreatePublicIPBandwidths(sc *gophercloud.ServiceClient) {
	var size = 10
	opts := bandwidths.CreateOpts{
		Name: "eeeeeeeeeeeeeeeee",
		Size: &size,
	}

	on, err := bandwidths.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, err := json.MarshalIndent(on, "", " ")
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(string(b))
	fmt.Println("its bandwidth info")
	fmt.Println("bandwidth id is ", on.ID)
	fmt.Println("bandwidth Size is ", on.Size)
	fmt.Println("bandwidth Name is ", on.Name)
	fmt.Println("bandwidth ShareType is ", on.ShareType)
	fmt.Println("bandwidth ChargeMode is ", on.ChargeMode)
	fmt.Println("bandwidth PublicipInfo is ", on.PublicipInfo)

	fmt.Println("Test create bandwidths success!")

}
func TestBatchCreatePublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var size = 10
	var count = 1
	opts := bandwidths.BatchCreateOpts{
		Name:  "eeeeeeeee",
		Size:  &size,
		Count: &count,
	}

	data, err := bandwidths.BatchCreate(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(string(b))

	fmt.Println("Test TestBatchCreatePublicIPBandwidths success!")

}

func TestInsertPublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var publicIPList []bandwidths.PublicIpInfoID

	opts := bandwidths.BandWidthInsertOpts{
		PublicipInfo: append(publicIPList, bandwidths.PublicIpInfoID{
			PublicIPID: "da68cc40-cfab-4e04-b4f8-741d77c6216e",
		}, bandwidths.PublicIpInfoID{
			PublicIPID: "fbb87197-9482-4122-9b62-c4b4d653c467",
		}),
	}

	on, err := bandwidths.Insert(sc, "fc011016-c2f7-4dad-8573-41c468f97d89", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, err := json.MarshalIndent(on, "", " ")
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(string(b))

	fmt.Println("its bandwidth info")
	fmt.Println("bandwidth id is ", on.ID)
	fmt.Println("bandwidth Size is ", on.Size)
	fmt.Println("bandwidth Name is ", on.Name)
	fmt.Println("bandwidth ShareType is ", on.ShareType)
	fmt.Println("bandwidth ChargeMode is ", on.ChargeMode)
	fmt.Println("bandwidth PublicipInfo is ", on.PublicipInfo)

	fmt.Println("TestInsertPublicIPBandwidths success!")

}
func TestRemovePublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var size = 10

	var publicIPList []bandwidths.PublicIpInfoID

	opts := bandwidths.BandWidthRemoveOpts{
		ChargeMode: "traffic",
		Size:       &size,
		PublicipInfo: append(publicIPList, bandwidths.PublicIpInfoID{
			PublicIPID: "da68cc40-cfab-4e04-b4f8-741d77c6216e",
		}, bandwidths.PublicIpInfoID{
			PublicIPID: "fbb87197-9482-4122-9b62-c4b4d653c467",
		}),
	}

	err := bandwidths.Remove(sc, "fc011016-c2f7-4dad-8573-41c468f97d89", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestRemovePublicIPBandwidths success!")

}

func TestDeletePublicIPBandwidths(sc *gophercloud.ServiceClient) {
	err := bandwidths.Delete(sc, "545ece15-5302-4285-a704-6866cd3acd64").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestRemovePublicIPBandwidths success!")
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
