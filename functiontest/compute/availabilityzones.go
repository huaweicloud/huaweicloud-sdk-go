package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	az "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
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

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//GetAZList(sc)
	GetAZListDetails(sc)
	fmt.Println("main end...")
}

func GetAZList(sc *gophercloud.ServiceClient) {

	allPages, err := az.List(sc).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	azinfo, err := az.ExtractAvailabilityZones(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get az info success")
	for _, data := range azinfo {
		fmt.Println("az hosts is ", data.Hosts)
		fmt.Println("az ZoneName is ", data.ZoneName)
		fmt.Println("az ZoneState is ", data.ZoneState)
	}
}


func GetAZListDetails(sc *gophercloud.ServiceClient) {

	allPages, err := az.ListDetail(sc).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	azinfo, err := az.ExtractAvailabilityZones(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get az info success")
	for _, data := range azinfo {
		for hostName,service:=range data.Hosts{
			fmt.Printf("az host name is %s, service is %s \n",hostName,service)
		}

		fmt.Println("az hosts list is ", data.Hosts)
		fmt.Println("az ZoneName is ", data.ZoneName)
		fmt.Println("az ZoneState is ", data.ZoneState)
	}
}
