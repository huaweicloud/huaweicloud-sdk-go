package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/whitelist"
)

var whitelistid string

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

	TestWhilteListCreate(sc)
	TestWhilteListList(sc)
	TestWhiteListUpdate(sc)
	TestWhiteListGet(sc)
	TestWhiteListDelete(sc)

	fmt.Println("main end...")
}

func TestWhilteListList(sc *gophercloud.ServiceClient) {
	enable := false
	allPages, err := whitelist.List(sc, whitelist.ListOpts{EnableWhitelist:&enable}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get whitelist list success!")
	allwhitelist, err := whitelist.ExtractWhiteLists(allPages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(allwhitelist, "", "   ")
	fmt.Println(string(b))
}

func TestWhilteListCreate(sc *gophercloud.ServiceClient) {
	enable := true
	opts := whitelist.CreateOpts{
		TenantId:        sc.ProjectID,
		ListenerId:      "cdc1a958-d9de-4d8f-b728-4a7867e723f9",
		EnableWhitelist: &enable,
		Whitelist:       "192.168.11.1,192.168.0.1/24",
	}
	whiteList, err := whitelist.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create whitelist success!")
	whitelistid = whiteList.WhiteList.ID
	p, err := json.MarshalIndent(whiteList, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(p))
}

func TestWhiteListUpdate(sc *gophercloud.ServiceClient) {
	list := ""
	enable := false
	updatOpts := whitelist.UpdateOpts{
		EnableWhitelist: &enable,
		Whitelist:       &list,
	}

	whiteList, err := whitelist.Update(sc, whitelistid, updatOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("Test update whitelist success!")
	p, _ := json.MarshalIndent(whiteList, "", " ")
	fmt.Println(string(p))
}

func TestWhiteListGet(sc *gophercloud.ServiceClient) {

	whiteList, err := whitelist.Get(sc, whitelistid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test get whitelist detail success!")

	p, _ := json.MarshalIndent(whiteList, "", " ")
	fmt.Println(string(p))

}

func TestWhiteListDelete(sc *gophercloud.ServiceClient) {
	err := whitelist.Delete(sc, whitelistid).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete whitelist success!")
}