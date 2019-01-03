package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

var portid string

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
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
		fmt.Println("get Network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestPortList(sc)
	TestPortCreate(sc)
	TestPortGet(sc)
	TestPortUpdate(sc)
	TestPortDelete(sc)

	fmt.Println("main end...")
}

func TestPortList(sc *gophercloud.ServiceClient) {
	allpages, err := ports.List(sc, ports.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	ports, err := ports.ExtractPorts(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get port list success!")
	p, _ := json.MarshalIndent(ports, "", " ")
	fmt.Println(string(p))
}

func TestPortGet(sc *gophercloud.ServiceClient) {
	port, err := ports.Get(sc, portid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get port detail success!")
	p, _ := json.MarshalIndent(port, "", " ")
	fmt.Println(string(p))
}

func TestPortCreate(sc *gophercloud.ServiceClient) {
	opts := ports.CreateOpts{
		Name:"testport",
		NetworkID:"879adbf2-d620-4978-982e-73cd3aabcfc0",
	}

	port, err := ports.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create port success!")
	portid=port.ID
	p, _ := json.MarshalIndent(port, "", " ")
	fmt.Println(string(p))
}

func TestPortUpdate(sc *gophercloud.ServiceClient) {
	opts := ports.UpdateOpts{
		Name:"testport2",
	}

	port,err := ports.Update(sc, portid,opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update port success!")
	p, _ := json.MarshalIndent(port, "", " ")
	fmt.Println(string(p))
}

func TestPortDelete(sc *gophercloud.ServiceClient) {
	err := ports.Delete(sc, portid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete port success!")
}