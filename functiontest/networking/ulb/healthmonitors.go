package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/monitors"

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

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//c529f8e5-a79f-44be-9bdc-e3b5d8844492

	//TestMonitorList(sc)
	//TestMonitorCreate(sc)

	//TestMonitorUpdate(sc)
	//TestMonitorGet(sc)
	TestMonitorDelete(sc)

	fmt.Println("main end...")
}

func TestMonitorCreate(sc *gophercloud.ServiceClient) {
	//('HTTP', 'HTTPS', 'PING', 'TCP', 'UDP_CONNECT')
	opts := monitors.CreateOpts{
		PoolID:     "abb1b6fd-7c90-46de-80a5-e64894383c12",
		Type:       "UDP_CONNECT",
		Delay:      21,
		Timeout:    12,
		MaxRetries: 3,
		Name:       "mmmmm",
	}

	resp, err := monitors.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("monitors Create success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestMonitorList(sc *gophercloud.ServiceClient) {
	allPages, err := monitors.List(sc, monitors.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test monitor List success!")

	allData, _ := monitors.ExtractMonitors(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestMonitorGet(sc *gophercloud.ServiceClient) {

	id := "04876434-6d22-4e27-bfce-8453c5c82921"

	resp, err := monitors.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("monitor get success!", resp)

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestMonitorUpdate(sc *gophercloud.ServiceClient) {

	id := "04876434-6d22-4e27-bfce-8453c5c82921"

	updatOpts := monitors.UpdateOpts{
		Name:  "KAKAK A monitor",
		Delay: 1,
	}

	resp, err := monitors.Update(sc, id, updatOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("monitor update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestMonitorDelete(sc *gophercloud.ServiceClient) {

	id := "04876434-6d22-4e27-bfce-8453c5c82921"
	err := monitors.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete monitor success!")
}
