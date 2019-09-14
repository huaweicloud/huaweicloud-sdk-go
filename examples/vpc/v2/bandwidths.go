package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/bandwidths"
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

	CreatePublicIPBandwidths(sc)
	BatchCreatePublicIPBandwidths(sc)
	InsertPublicIPBandwidths(sc)
	RemovePublicIPBandwidths(sc)
	DeletePublicIPBandwidths(sc)
	ModifyPublicIPBandwidths(sc)

	fmt.Println("main end...")
}

//Create a BandWidth
func CreatePublicIPBandwidths(sc *gophercloud.ServiceClient) {
	var size = 10
	opts := bandwidths.CreateOpts{
		Name: "xxxxxx",
		Size: &size,
	}

	result, err := bandwidths.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("bandwidth: %+v\r\n", result)

	fmt.Println("bandwidth id is:", result.ID)
	fmt.Println("bandwidth Size is:", result.Size)
	fmt.Println("bandwidth Name is:", result.Name)
	fmt.Println("bandwidth ShareType is:", result.ShareType)
	fmt.Println("bandwidth ChargeMode is:", result.ChargeMode)
	fmt.Println("bandwidth PublicipInfo is:", result.PublicipInfo)

	fmt.Println("Create success!")

}

//Batch Create BandWidth
func BatchCreatePublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var size = 10
	var count = 1
	opts := bandwidths.BatchCreateOpts{
		Name:  "xxxxxx",
		Size:  &size,
		Count: &count,
	}

	result, err := bandwidths.BatchCreate(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("bandwidth: %+v\r\n", result)

	resultBandwidth := *result

	for _, resp := range resultBandwidth {
		fmt.Println("bandwidth id is:", resp.ID)
		fmt.Println("bandwidth Size is:", resp.Size)
		fmt.Println("bandwidth Name is:", resp.Name)
		fmt.Println("bandwidth ShareType is:", resp.ShareType)
		fmt.Println("bandwidth ChargeMode is:", resp.ChargeMode)
		fmt.Println("bandwidth PublicipInfo is:", resp.PublicipInfo)
	}

	fmt.Println("BatchCreatePublicIPBandwidths success!")

}

// Insert ip into BandWidth
func InsertPublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var publicIPList []bandwidths.PublicIpInfoID

	opts := bandwidths.BandWidthInsertOpts{
		PublicipInfo: append(publicIPList, bandwidths.PublicIpInfoID{
			PublicIPID: "xxxxxx",
		}, bandwidths.PublicIpInfoID{
			PublicIPID: "xxxxxx",
		}),
	}

	result, err := bandwidths.Insert(sc, "xxxxxx", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("bandwidth: %+v\r\n", result)

	fmt.Println("bandwidth id is:", result.ID)
	fmt.Println("bandwidth Size is:", result.Size)
	fmt.Println("bandwidth Name is:", result.Name)
	fmt.Println("bandwidth ShareType is:", result.ShareType)
	fmt.Println("bandwidth ChargeMode is:", result.ChargeMode)
	fmt.Println("bandwidth PublicipInfo is:", result.PublicipInfo)

	fmt.Println("InsertPublicIPBandwidths success!")

}

// Remove ip into BandWidth
func RemovePublicIPBandwidths(sc *gophercloud.ServiceClient) {

	var size = 10

	var publicIPList []bandwidths.PublicIpInfoID

	opts := bandwidths.BandWidthRemoveOpts{
		ChargeMode: "traffic",
		Size:       &size,
		PublicipInfo: append(publicIPList, bandwidths.PublicIpInfoID{
			PublicIPID: "xxxxxx",
		}, bandwidths.PublicIpInfoID{
			PublicIPID: "xxxxxx",
		}),
	}

	err := bandwidths.Remove(sc, "xxxxxx", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("RemovePublicIPBandwidths success!")

}

// Delete a BandWidth
func DeletePublicIPBandwidths(sc *gophercloud.ServiceClient) {
	err := bandwidths.Delete(sc, "xxxxxx").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("RemovePublicIPBandwidths success!")
}

//Update a BandWidth
func ModifyPublicIPBandwidths(sc *gophercloud.ServiceClient) {

	size := 10
	//modify name
	//opts := bandwidths.UpdateOpts{
	//	Bandwidth: bandwidths.Bandwidth{
	//		Name: "xxxxxx",
	//	},
	//}

	//modify bandwidth size
	opts := bandwidths.UpdateOpts{
		Bandwidth: bandwidths.Bandwidth{
			Name: "xxxxxx",
			Size: size,
		},
		ExtendParam: &bandwidths.ExtendParam{
			IsAutoPay: "true",
		},
	}

	data, err := bandwidths.Update(sc, "xxxxxx", opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if order, ok := data.(bandwidths.PrePaid); ok {
		fmt.Println("its order id ")
		fmt.Println("order id is:", order.OrderID)
	}

	if on, ok := data.(bandwidths.PostPaid); ok {
		fmt.Println("its bandwidth info")
		fmt.Println("bandwidth id is:", on.ID)
		fmt.Println("bandwidth Size is:", on.Size)
		fmt.Println("bandwidth Name is:", on.Name)
		fmt.Println("bandwidth ShareType is:", on.ShareType)
		fmt.Println("bandwidth ChargeMode is:", on.ChargeMode)
		fmt.Println("bandwidth PublicipInfo is:", on.PublicipInfo)
	}

	fmt.Println("Update success!")

}
