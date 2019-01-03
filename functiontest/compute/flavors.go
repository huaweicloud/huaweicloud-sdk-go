package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	GetFlavorList(sc)
	GetFlavorExtraSpecsList(sc)
	fmt.Println("main end...")
}

func GetFlavorList(sc *gophercloud.ServiceClient) {

	fpage, err := flavors.ListDetail(sc, nil).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	f, err := flavors.ExtractFlavors(fpage)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("get Flavor details success")

	for _, data := range f {
		fmt.Println("get flavor name", data.Name)
		fmt.Println("get flavor id ", data.ID)
		fmt.Println("get flavor Disk", data.Disk)
		fmt.Println("get flavor Ephemeral", data.Ephemeral)
		fmt.Println("get flavor IsPublic ", data.IsPublic)
		fmt.Println("get flavor RAM", data.RAM)
		fmt.Println("get flavor RxTxFactor", data.RxTxFactor)
		fmt.Println("get flavor Swap ", data.Swap)
		fmt.Println("get flavor VCPUs", data.VCPUs)
	}

}

func GetFlavorExtraSpecsList(sc *gophercloud.ServiceClient) {
	flavorID := "t6.xlarge.4"
	data, err := flavors.ListExtraSpecs(sc, flavorID).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Get Flavor Extra Specs info success ")

	for k, b := range data {
		fmt.Printf("os_extra_specs key is %s, value is %s\n", k, b)
		//cond:operation:status : normal;abandon;sellout ;obt;promotion;
		if k == "cond:operation:status" {
			fmt.Printf("flavor %s status is %s\n", flavorID, b)
		}

	}

}
