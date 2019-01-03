package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
)

var floatingIpid string

func main() {
	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestFloatingIPList(sc)
	TestFloatingIPCreate(sc)
	TestFloatingIPGet(sc)
	TestFloatingIPUpdate(sc)
	TestFloatingIPDelete(sc)

	fmt.Println("main end...")
}


func TestFloatingIPList(sc *gophercloud.ServiceClient) {
	allpages, err := floatingips.List(sc, floatingips.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	floatingips, err := floatingips.ExtractFloatingIPs(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get floatingip list success!")
	p, _ := json.MarshalIndent(floatingips, "", " ")
	fmt.Println(string(p))
}

func TestFloatingIPGet(sc *gophercloud.ServiceClient) {
	floatingip, err := floatingips.Get(sc, floatingIpid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get floatingip detail success!")
	p, _ := json.MarshalIndent(floatingip, "", " ")
	fmt.Println(string(p))
}

func TestFloatingIPCreate(sc *gophercloud.ServiceClient) {
	opts := floatingips.CreateOpts{
		FloatingNetworkID:"0a2228f2-7f8a-45f1-8e09-9039e1d09975",
		PortID:"2f8254a3-c7ec-4600-bc10-cdfdf9a4384b",
	}

	floatingip, err := floatingips.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create floatingip success!")
	floatingIpid=floatingip.ID
	p, _ := json.MarshalIndent(floatingip, "", " ")
	fmt.Println(string(p))
}

func TestFloatingIPUpdate(sc *gophercloud.ServiceClient) {
	//newportid := "9720c57f-fccf-439e-901a-77d57a01c7ac"
	//fixedip := ""
	opts := floatingips.UpdateOpts{
		//PortID:&newportid,
		//FixedIP:&fixedip,
		PortID:nil,
		FixedIP:nil,
	}

	floatingip,err := floatingips.Update(sc, floatingIpid,opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update floatingip success!")
	p, _ := json.MarshalIndent(floatingip, "", " ")
	fmt.Println(string(p))
}

func TestFloatingIPDelete(sc *gophercloud.ServiceClient) {
	err := floatingips.Delete(sc, floatingIpid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete floatingip success!")
}